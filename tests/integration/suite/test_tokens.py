import pytest
import rancher
from .conftest import BASE_URL




def test_websocket(admin_mc):
    client = rancher.Client(url=BASE_URL, token=admin_mc.client.token,
                            verify=False)
    # make a request that looks like a websocket
    client._session.headers["Connection"] = "upgrade"
    client._session.headers["Upgrade"] = "websocket"
    client._session.headers["Origin"] = "badStuff"
    # do something with client now that we have a "websocket"

    with pytest.raises(rancher.ApiError) as e:
        client.list_cluster()

    assert e.value.error.Code.Status == 403
