package client

import (
	"github.com/rancher/norman/types"
)

const (
	ClusterType                                      = "cluster"
	ClusterFieldAPIEndpoint                          = "apiEndpoint"
	ClusterFieldAgentImage                           = "agentImage"
	ClusterFieldAllocatable                          = "allocatable"
	ClusterFieldAmazonElasticContainerServiceConfig  = "amazonElasticContainerServiceConfig"
	ClusterFieldAnnotations                          = "annotations"
	ClusterFieldAppliedEnableNetworkPolicy           = "appliedEnableNetworkPolicy"
	ClusterFieldAppliedPodSecurityPolicyTemplateName = "appliedPodSecurityPolicyTemplateId"
	ClusterFieldAppliedSpec                          = "appliedSpec"
	ClusterFieldAzureKubernetesServiceConfig         = "azureKubernetesServiceConfig"
	ClusterFieldCACert                               = "caCert"
	ClusterFieldCapacity                             = "capacity"
	ClusterFieldComponentStatuses                    = "componentStatuses"
	ClusterFieldConditions                           = "conditions"
	ClusterFieldCreated                              = "created"
	ClusterFieldCreatorID                            = "creatorId"
	ClusterFieldDefaultClusterRoleForProjectMembers  = "defaultClusterRoleForProjectMembers"
	ClusterFieldDefaultPodSecurityPolicyTemplateID   = "defaultPodSecurityPolicyTemplateId"
	ClusterFieldDescription                          = "description"
	ClusterFieldDesiredAgentImage                    = "desiredAgentImage"
	ClusterFieldDockerRootDir                        = "dockerRootDir"
	ClusterFieldDriver                               = "driver"
	ClusterFieldEnableNetworkPolicy                  = "enableNetworkPolicy"
	ClusterFieldFailedSpec                           = "failedSpec"
	ClusterFieldGoogleKubernetesEngineConfig         = "googleKubernetesEngineConfig"
	ClusterFieldImportedConfig                       = "importedConfig"
	ClusterFieldInternal                             = "internal"
	ClusterFieldLabels                               = "labels"
	ClusterFieldLimits                               = "limits"
	ClusterFieldName                                 = "name"
	ClusterFieldOwnerReferences                      = "ownerReferences"
	ClusterFieldRancherKubernetesEngineConfig        = "rancherKubernetesEngineConfig"
	ClusterFieldRemoved                              = "removed"
	ClusterFieldRequested                            = "requested"
	ClusterFieldState                                = "state"
	ClusterFieldTransitioning                        = "transitioning"
	ClusterFieldTransitioningMessage                 = "transitioningMessage"
	ClusterFieldUUID                                 = "uuid"
	ClusterFieldVersion                              = "version"
)

type Cluster struct {
	types.Resource
	APIEndpoint                          string                               `json:"apiEndpoint,omitempty" yaml:"apiEndpoint,omitempty"`
	AgentImage                           string                               `json:"agentImage,omitempty" yaml:"agentImage,omitempty"`
	Allocatable                          map[string]string                    `json:"allocatable,omitempty" yaml:"allocatable,omitempty"`
	AmazonElasticContainerServiceConfig  *AmazonElasticContainerServiceConfig `json:"amazonElasticContainerServiceConfig,omitempty" yaml:"amazonElasticContainerServiceConfig,omitempty"`
	Annotations                          map[string]string                    `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AppliedEnableNetworkPolicy           bool                                 `json:"appliedEnableNetworkPolicy,omitempty" yaml:"appliedEnableNetworkPolicy,omitempty"`
	AppliedPodSecurityPolicyTemplateName string                               `json:"appliedPodSecurityPolicyTemplateId,omitempty" yaml:"appliedPodSecurityPolicyTemplateId,omitempty"`
	AppliedSpec                          *ClusterSpec                         `json:"appliedSpec,omitempty" yaml:"appliedSpec,omitempty"`
	AzureKubernetesServiceConfig         *AzureKubernetesServiceConfig        `json:"azureKubernetesServiceConfig,omitempty" yaml:"azureKubernetesServiceConfig,omitempty"`
	CACert                               string                               `json:"caCert,omitempty" yaml:"caCert,omitempty"`
	Capacity                             map[string]string                    `json:"capacity,omitempty" yaml:"capacity,omitempty"`
	ComponentStatuses                    []ClusterComponentStatus             `json:"componentStatuses,omitempty" yaml:"componentStatuses,omitempty"`
	Conditions                           []ClusterCondition                   `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Created                              string                               `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                            string                               `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	DefaultClusterRoleForProjectMembers  string                               `json:"defaultClusterRoleForProjectMembers,omitempty" yaml:"defaultClusterRoleForProjectMembers,omitempty"`
	DefaultPodSecurityPolicyTemplateID   string                               `json:"defaultPodSecurityPolicyTemplateId,omitempty" yaml:"defaultPodSecurityPolicyTemplateId,omitempty"`
	Description                          string                               `json:"description,omitempty" yaml:"description,omitempty"`
	DesiredAgentImage                    string                               `json:"desiredAgentImage,omitempty" yaml:"desiredAgentImage,omitempty"`
	DockerRootDir                        string                               `json:"dockerRootDir,omitempty" yaml:"dockerRootDir,omitempty"`
	Driver                               string                               `json:"driver,omitempty" yaml:"driver,omitempty"`
	EnableNetworkPolicy                  *bool                                `json:"enableNetworkPolicy,omitempty" yaml:"enableNetworkPolicy,omitempty"`
	FailedSpec                           *ClusterSpec                         `json:"failedSpec,omitempty" yaml:"failedSpec,omitempty"`
	GoogleKubernetesEngineConfig         *GoogleKubernetesEngineConfig        `json:"googleKubernetesEngineConfig,omitempty" yaml:"googleKubernetesEngineConfig,omitempty"`
	ImportedConfig                       *ImportedConfig                      `json:"importedConfig,omitempty" yaml:"importedConfig,omitempty"`
	Internal                             bool                                 `json:"internal,omitempty" yaml:"internal,omitempty"`
	Labels                               map[string]string                    `json:"labels,omitempty" yaml:"labels,omitempty"`
	Limits                               map[string]string                    `json:"limits,omitempty" yaml:"limits,omitempty"`
	Name                                 string                               `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences                      []OwnerReference                     `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	RancherKubernetesEngineConfig        *RancherKubernetesEngineConfig       `json:"rancherKubernetesEngineConfig,omitempty" yaml:"rancherKubernetesEngineConfig,omitempty"`
	Removed                              string                               `json:"removed,omitempty" yaml:"removed,omitempty"`
	Requested                            map[string]string                    `json:"requested,omitempty" yaml:"requested,omitempty"`
	State                                string                               `json:"state,omitempty" yaml:"state,omitempty"`
	Transitioning                        string                               `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage                 string                               `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                                 string                               `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Version                              *Info                                `json:"version,omitempty" yaml:"version,omitempty"`
}

type ClusterCollection struct {
	types.Collection
	Data   []Cluster `json:"data,omitempty"`
	client *ClusterClient
}

type ClusterClient struct {
	apiClient *Client
}

type ClusterOperations interface {
	List(opts *types.ListOpts) (*ClusterCollection, error)
	Create(opts *Cluster) (*Cluster, error)
	Update(existing *Cluster, updates interface{}) (*Cluster, error)
	Replace(existing *Cluster) (*Cluster, error)
	ByID(id string) (*Cluster, error)
	Delete(container *Cluster) error

	ActionExportYaml(resource *Cluster) (*ExportOutput, error)

	ActionGenerateKubeconfig(resource *Cluster) (*GenerateKubeConfigOutput, error)

	ActionImportYaml(resource *Cluster, input *ImportClusterYamlInput) (*ImportYamlOutput, error)

	ActionRotateCertificates(resource *Cluster) (*RotateCertificateOutput, error)
}

func newClusterClient(apiClient *Client) *ClusterClient {
	return &ClusterClient{
		apiClient: apiClient,
	}
}

func (c *ClusterClient) Create(container *Cluster) (*Cluster, error) {
	resp := &Cluster{}
	err := c.apiClient.Ops.DoCreate(ClusterType, container, resp)
	return resp, err
}

func (c *ClusterClient) Update(existing *Cluster, updates interface{}) (*Cluster, error) {
	resp := &Cluster{}
	err := c.apiClient.Ops.DoUpdate(ClusterType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterClient) Replace(obj *Cluster) (*Cluster, error) {
	resp := &Cluster{}
	err := c.apiClient.Ops.DoReplace(ClusterType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ClusterClient) List(opts *types.ListOpts) (*ClusterCollection, error) {
	resp := &ClusterCollection{}
	err := c.apiClient.Ops.DoList(ClusterType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ClusterCollection) Next() (*ClusterCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterClient) ByID(id string) (*Cluster, error) {
	resp := &Cluster{}
	err := c.apiClient.Ops.DoByID(ClusterType, id, resp)
	return resp, err
}

func (c *ClusterClient) Delete(container *Cluster) error {
	return c.apiClient.Ops.DoResourceDelete(ClusterType, &container.Resource)
}

func (c *ClusterClient) ActionExportYaml(resource *Cluster) (*ExportOutput, error) {
	resp := &ExportOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "exportYaml", &resource.Resource, nil, resp)
	return resp, err
}

func (c *ClusterClient) ActionGenerateKubeconfig(resource *Cluster) (*GenerateKubeConfigOutput, error) {
	resp := &GenerateKubeConfigOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "generateKubeconfig", &resource.Resource, nil, resp)
	return resp, err
}

func (c *ClusterClient) ActionImportYaml(resource *Cluster, input *ImportClusterYamlInput) (*ImportYamlOutput, error) {
	resp := &ImportYamlOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "importYaml", &resource.Resource, input, resp)
	return resp, err
}

func (c *ClusterClient) ActionRotateCertificates(resource *Cluster) (*RotateCertificateOutput, error) {
	resp := &RotateCertificateOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "rotateCertificates", &resource.Resource, nil, resp)
	return resp, err
}
