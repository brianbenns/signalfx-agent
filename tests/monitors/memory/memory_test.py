from functools import partial as p

import pytest

from helpers.assertions import has_any_metric_or_dim, has_log_message
from helpers.util import get_monitor_dims_from_selfdescribe, get_monitor_metrics_from_selfdescribe, run_agent, wait_for

pytestmark = [pytest.mark.collectd, pytest.mark.memory, pytest.mark.monitor_without_endpoints]


def test_memory():
    expected_metrics = get_monitor_metrics_from_selfdescribe("collectd/memory")
    expected_dims = get_monitor_dims_from_selfdescribe("collectd/memory")
    with run_agent(
        """
    monitors:
      - type: collectd/memory
    """
    ) as [backend, get_output, _]:
        assert wait_for(
            p(has_any_metric_or_dim, backend, expected_metrics, expected_dims), timeout_seconds=60
        ), "timed out waiting for metrics and/or dimensions!"
        assert not has_log_message(get_output().lower(), "error"), "error found in agent output!"
