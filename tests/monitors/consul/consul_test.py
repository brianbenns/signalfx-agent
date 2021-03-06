import os

import pytest

from helpers.kubernetes.utils import get_discovery_rule, run_k8s_monitors_test
from helpers.util import get_monitor_dims_from_selfdescribe, get_monitor_metrics_from_selfdescribe

pytestmark = [pytest.mark.collectd, pytest.mark.consul, pytest.mark.monitor_with_endpoints]


@pytest.mark.k8s
@pytest.mark.kubernetes
def test_consul_in_k8s(agent_image, minikube, k8s_observer, k8s_test_timeout, k8s_namespace):
    yaml = os.path.join(os.path.dirname(os.path.realpath(__file__)), "consul-k8s.yaml")
    monitors = [
        {
            "type": "collectd/consul",
            "discoveryRule": get_discovery_rule(yaml, k8s_observer, namespace=k8s_namespace),
            "aclToken": "testing123",
            "signalFxAccessToken": "testing123",
            "enhancedMetrics": True,
        }
    ]
    run_k8s_monitors_test(
        agent_image,
        minikube,
        monitors,
        namespace=k8s_namespace,
        yamls=[yaml],
        observer=k8s_observer,
        expected_metrics=get_monitor_metrics_from_selfdescribe(monitors[0]["type"]),
        expected_dims=get_monitor_dims_from_selfdescribe(monitors[0]["type"]),
        test_timeout=k8s_test_timeout,
    )
