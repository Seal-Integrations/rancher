package cluster

import (
	"context"
	errors2 "errors"
	"fmt"

	"github.com/rancher/kontainer-engine/logstream"
	"github.com/rancher/kontainer-engine/types"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
)

const (
	PreCreating = "Pre-Creating"
	Creating    = "Creating"
	PostCheck   = "Post-Checking"
	Running     = "Running"
	Error       = "Error"
	Updating    = "Updating"
)

var (
	// ErrClusterExists This error is checked in rancher, don't change the string
	ErrClusterExists = errors2.New("cluster already exists")
)

// Cluster represents a kubernetes cluster
type Cluster struct {
	// The cluster driver to provision cluster
	Driver types.Driver `json:"-"`
	// The name of the cluster driver
	DriverName string `json:"driverName,omitempty" yaml:"driver_name,omitempty"`
	// The name of the cluster
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// The status of the cluster
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// specific info about kubernetes cluster
	// Kubernetes cluster version
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
	// Service account token to access kubernetes API
	ServiceAccountToken string `json:"serviceAccountToken,omitempty" yaml:"service_account_token,omitempty"`
	// Kubernetes API master endpoint
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	// Username for http basic authentication
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	// Password for http basic authentication
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	// Root CaCertificate for API server(base64 encoded)
	RootCACert string `json:"rootCACert,omitempty" yaml:"root_ca_cert,omitempty"`
	// Client Certificate(base64 encoded)
	ClientCertificate string `json:"clientCertificate,omitempty" yaml:"client_certificate,omitempty"`
	// Client private key(base64 encoded)
	ClientKey string `json:"clientKey,omitempty" yaml:"client_key,omitempty"`
	// Node count in the cluster
	NodeCount int64 `json:"nodeCount,omitempty" yaml:"node_count,omitempty"`

	// Metadata store specific driver options per cloud provider
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	PersistStore PersistentStore `json:"-" yaml:"-"`

	ConfigGetter ConfigGetter `json:"-" yaml:"-"`

	Logger logstream.Logger `json:"-" yaml:"-"`
}

// PersistentStore defines the interface for persist options like check and store
type PersistentStore interface {
	GetStatus(name string) (string, error)
	Get(name string) (Cluster, error)
	Remove(name string) error
	Store(cluster Cluster) error
	PersistStatus(cluster Cluster, status string) error
}

// ConfigGetter defines the interface for getting the driver options.
type ConfigGetter interface {
	GetConfig() (types.DriverOptions, error)
}

// Create creates a cluster
func (c *Cluster) Create(ctx context.Context) error {
	if err := c.createInner(ctx); err != nil && err != ErrClusterExists {
		c.PersistStore.PersistStatus(*c, Error)
		return err
	}
	return c.PersistStore.PersistStatus(*c, Running)
}

func (c *Cluster) create(ctx context.Context, clusterInfo *types.ClusterInfo) error {
	if c.Status == PostCheck {
		return nil
	}

	if err := c.PersistStore.PersistStatus(*c, PreCreating); err != nil {
		return err
	}

	// get cluster config from cli flags or json config
	driverOpts, err := c.ConfigGetter.GetConfig()
	if err != nil {
		return err
	}

	// also set metadata value to retrieve the cluster info
	for k, v := range c.Metadata {
		driverOpts.StringOptions[k] = v
	}

	if err := c.PersistStore.PersistStatus(*c, Creating); err != nil {
		return err
	}

	// create cluster
	info, err := c.Driver.Create(ctx, &driverOpts, clusterInfo)
	if info != nil {
		transformClusterInfo(c, info)
	}
	return err
}

func (c *Cluster) PostCheck(ctx context.Context) error {
	if err := c.PersistStore.PersistStatus(*c, PostCheck); err != nil {
		return err
	}

	// receive cluster info back
	info, err := c.Driver.PostCheck(ctx, toInfo(c))
	if err != nil {
		return err
	}

	transformClusterInfo(c, info)

	// persist cluster info
	return c.Store()
}

func (c *Cluster) GenerateServiceAccount(ctx context.Context) error {
	if err := c.restore(); err != nil {
		return err
	}

	if err := c.PersistStore.PersistStatus(*c, Updating); err != nil {
		return err
	}

	// receive cluster info back
	info, err := c.Driver.PostCheck(ctx, toInfo(c))
	if err != nil {
		return err
	}

	transformClusterInfo(c, info)

	// persist cluster info
	return c.Store()
}

func (c *Cluster) RemoveLegacyServiceAccount(ctx context.Context) error {
	if err := c.restore(); err != nil {
		return err
	}

	return c.Driver.RemoveLegacyServiceAccount(ctx, toInfo(c))
}

func (c *Cluster) createInner(ctx context.Context) error {
	// check if it is already created
	c.restore()

	var info *types.ClusterInfo
	if c.Status == Error {
		logrus.Errorf("Cluster %s previously failed to create", c.Name)
		info = toInfo(c)
	}

	if c.Status == Updating || c.Status == Running || c.Status == PostCheck {
		logrus.Infof("Cluster %s already exists.", c.Name)
		return ErrClusterExists
	}

	if err := c.create(ctx, info); err != nil {
		return err
	}

	return c.PostCheck(ctx)
}

// Update updates a cluster
func (c *Cluster) Update(ctx context.Context) error {
	if err := c.restore(); err != nil {
		return err
	}

	if c.Status == Error {
		logrus.Errorf("Cluster %s previously failed to create", c.Name)
		return c.Create(ctx)
	}

	if c.Status == PreCreating || c.Status == Creating {
		logrus.Errorf("Cluster %s has not been created.", c.Name)
		return fmt.Errorf("cluster %s has not been created", c.Name)
	}

	driverOpts, err := c.ConfigGetter.GetConfig()
	if err != nil {
		return err
	}
	driverOpts.StringOptions["name"] = c.Name
	for k, v := range c.Metadata {
		driverOpts.StringOptions[k] = v
	}

	if err := c.PersistStore.PersistStatus(*c, Updating); err != nil {
		return err
	}

	info := toInfo(c)
	info, err = c.Driver.Update(ctx, info, &driverOpts)
	if err != nil {
		return err
	}

	transformClusterInfo(c, info)

	return c.PostCheck(ctx)
}

func (c *Cluster) GetVersion(ctx context.Context) (*types.KubernetesVersion, error) {
	return c.Driver.GetVersion(ctx, toInfo(c))
}

func (c *Cluster) SetVersion(ctx context.Context, version *types.KubernetesVersion) error {
	return c.Driver.SetVersion(ctx, toInfo(c), version)
}

func (c *Cluster) GetClusterSize(ctx context.Context) (*types.NodeCount, error) {
	return c.Driver.GetClusterSize(ctx, toInfo(c))
}

func (c *Cluster) SetClusterSize(ctx context.Context, count *types.NodeCount) error {
	return c.Driver.SetClusterSize(ctx, toInfo(c), count)
}

func transformClusterInfo(c *Cluster, clusterInfo *types.ClusterInfo) {
	c.ClientCertificate = clusterInfo.ClientCertificate
	c.ClientKey = clusterInfo.ClientKey
	c.RootCACert = clusterInfo.RootCaCertificate
	c.Username = clusterInfo.Username
	c.Password = clusterInfo.Password
	c.Version = clusterInfo.Version
	c.Endpoint = clusterInfo.Endpoint
	c.NodeCount = clusterInfo.NodeCount
	c.Metadata = clusterInfo.Metadata
	c.ServiceAccountToken = clusterInfo.ServiceAccountToken
	c.Status = clusterInfo.Status
}

func toInfo(c *Cluster) *types.ClusterInfo {
	return &types.ClusterInfo{
		ClientCertificate:   c.ClientCertificate,
		ClientKey:           c.ClientKey,
		RootCaCertificate:   c.RootCACert,
		Username:            c.Username,
		Password:            c.Password,
		Version:             c.Version,
		Endpoint:            c.Endpoint,
		NodeCount:           c.NodeCount,
		Metadata:            c.Metadata,
		ServiceAccountToken: c.ServiceAccountToken,
		Status:              c.Status,
	}
}

// Remove removes a cluster
func (c *Cluster) Remove(ctx context.Context) error {
	defer c.PersistStore.Remove(c.Name)
	if err := c.restore(); errors.IsNotFound(err) {
		return nil
	} else if err != nil {
		return err
	}

	return c.Driver.Remove(ctx, toInfo(c))
}

func (c *Cluster) GetCapabilities(ctx context.Context) (*types.Capabilities, error) {
	return c.Driver.GetCapabilities(ctx)
}

func (c *Cluster) getState() (string, error) {
	return c.PersistStore.GetStatus(c.Name)
}

// Store persists cluster information
func (c *Cluster) Store() error {
	return c.PersistStore.Store(*c)
}

func (c *Cluster) restore() error {
	cluster, err := c.PersistStore.Get(c.Name)
	if err != nil {
		return err
	}
	info := toInfo(&cluster)
	transformClusterInfo(c, info)
	return nil
}

// NewCluster create a cluster interface to do operations
func NewCluster(driverName, addr, name string, configGetter ConfigGetter, persistStore PersistentStore) (*Cluster, error) {
	rpcClient, err := types.NewClient(driverName, addr)
	if err != nil {
		return nil, err
	}
	return &Cluster{
		Driver:       rpcClient,
		DriverName:   driverName,
		Name:         name,
		ConfigGetter: configGetter,
		PersistStore: persistStore,
	}, nil
}

func FromCluster(cluster *Cluster, addr string, configGetter ConfigGetter, persistStore PersistentStore) (*Cluster, error) {
	rpcClient, err := types.NewClient(cluster.DriverName, addr)
	if err != nil {
		return nil, err
	}
	cluster.Driver = rpcClient
	cluster.ConfigGetter = configGetter
	cluster.PersistStore = persistStore
	return cluster, nil
}
