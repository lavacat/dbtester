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

package control

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/coreos/dbtester/remotestorage"
)

func step4UploadLogs(cfg Config) error {
	if err := uploadToGoogle(cfg.ControlLogPath, cfg); err != nil {
		return err
	}
	if err := uploadToGoogle(cfg.ControlLatencyThroughputTimeseriesPath, cfg); err != nil {
		return err
	}
	if err := uploadToGoogle(cfg.ControlLatencyDistributionSummaryPath, cfg); err != nil {
		return err
	}
	if err := uploadToGoogle(cfg.ControlLatencyDistributionPercentilePath, cfg); err != nil {
		return err
	}
	if err := uploadToGoogle(cfg.ControlLatencyDistributionAllPath, cfg); err != nil {
		return err
	}
	return nil
}

func uploadToGoogle(path string, cfg Config) error {
	u, err := remotestorage.NewGoogleCloudStorage([]byte(cfg.GoogleCloudStorageKey), cfg.GoogleCloudProjectName)
	if err != nil {
		return err
	}

	srcPath := path
	dstPath := filepath.Base(path)
	if !strings.HasPrefix(dstPath, cfg.TestName) {
		dstPath = fmt.Sprintf("%s-%s", cfg.TestName, dstPath)
	}
	dstPath = filepath.Join(cfg.GoogleCloudStorageSubDirectory, dstPath)

	var uerr error
	for k := 0; k < 15; k++ {
		if uerr = u.UploadFile(cfg.GoogleCloudStorageBucketName, srcPath, dstPath); uerr != nil {
			plog.Printf("#%d: UploadFile error %v", k, uerr)
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}
	return uerr
}
