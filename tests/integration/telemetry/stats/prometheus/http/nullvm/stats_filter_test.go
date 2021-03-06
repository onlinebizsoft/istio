// Copyright 2020 Istio Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nullvm

import (
	"testing"

	"istio.io/istio/pkg/test/framework/label"
	"istio.io/istio/pkg/test/framework/resource/environment"

	"istio.io/istio/pkg/test/framework"
	"istio.io/istio/pkg/test/framework/components/istio"
	common "istio.io/istio/tests/integration/telemetry/stats/prometheus/http"
)

// TestStatsFilter verifies the stats filter could emit expected client and server side metrics.
// This test focuses on stats filter and metadata exchange filter could work coherently with
// proxy bootstrap config. To avoid flake, it does not verify correctness of metrics, which
// should be covered by integration test in proxy repo.
func TestStatsFilter(t *testing.T) {
	common.TestStatsFilter(t)
}

func TestMain(m *testing.M) {
	framework.NewSuite("stats_filter_test", m).
		RequireEnvironment(environment.Kube).
		RequireSingleCluster().
		Label(label.CustomSetup).
		SetupOnEnv(environment.Kube, istio.Setup(common.GetIstioInstance(), setupConfig)).
		Setup(common.TestSetup).
		// Application scraping will not work with permissive mode
		Setup(common.SetupStrictMTLS).
		Run()
}

func setupConfig(cfg *istio.Config) {
	if cfg == nil {
		return
	}
	// disable mixer telemetry and enable telemetry v2
	cfg.Values["telemetry.enabled"] = "true"
	cfg.Values["telemetry.v1.enabled"] = "false"
	cfg.Values["telemetry.v2.enabled"] = "true"
	cfg.Values["telemetry.v2.prometheus.enabled"] = "true"
	cfg.Values["prometheus.enabled"] = "true"
}
