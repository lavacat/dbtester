// Copyright 2017 CoreOS, Inc.
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

package analyze

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// RawData defines how to aggregate data from each machine.
type RawData struct {
	Legend                     string   `yaml:"legend"`
	DatabaseTag                string   `yaml:"database_tag"`
	OutputPath                 string   `yaml:"output_path"`
	SourceSystemMetricsPaths   []string `yaml:"source_system_metrics_paths"`
	SourceBenchmarkMetricsPath string   `yaml:"source_benchmark_metrics_path"`
}

// Config defines analyze configuration.
type Config struct {
	Title             string       `yaml:"title"`
	AllAggregatedPath string       `yaml:"all_aggregated_path"`
	RawData           []RawData    `yaml:"raw_data"`
	PlotList          []PlotConfig `yaml:"plot_list"`
	READMEConfig      READMEConfig `yaml:"readme"`
}

// readConfig reads analyze configuration.
func readConfig(fpath string) (Config, error) {
	bts, err := ioutil.ReadFile(fpath)
	if err != nil {
		return Config{}, err
	}
	rs := Config{}
	if err := yaml.Unmarshal(bts, &rs); err != nil {
		return Config{}, err
	}
	return rs, nil
}