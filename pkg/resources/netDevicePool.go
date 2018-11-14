// Copyright 2018 Intel Corp. All Rights Reserved.
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

package resources

import (
	"github.com/golang/glog"
	"github.com/intel/sriov-network-device-plugin/pkg/types"
	"github.com/intel/sriov-network-device-plugin/pkg/utils"
	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"
)

/*
	netDevicePool extends resourcePool and overrides:
	GetDeviceFile(),
	GetEnvs()
	GetMounts()
	Probe()
*/
type netDevicePool struct {
}

func newNetDevicePool(rc *types.ResourceConfig) types.ResourcePool {
	return &resourcePool{
		config:        rc,
		devices:       make(map[string]*pluginapi.Device),
		deviceFiles:   make(map[string]string),
		IBaseResource: &netDevicePool{},
	}
}

// Overrides GetDeviceFile() method
func (rp *netDevicePool) GetDeviceFile(dev string) (devFile string, err error) {
	// There are no specific device file

	return // empty string
}

func (rp *netDevicePool) GetEnvs(resourceName string, deviceIDs []string) map[string]string {
	glog.Infof("generic GetEnvs() called")
	envs := make(map[string]string)
	values := ""
	lastIndex := len(deviceIDs) - 1
	for i, s := range deviceIDs {
		values += s
		if i == lastIndex {
			break
		}
		values += " "
	}
	envs[resourceName] = values
	return envs
}

func (rp *netDevicePool) GetMounts() []*pluginapi.Mount {
	glog.Infof("generic GetMounts() called")
	mounts := make([]*pluginapi.Mount, 0)
	return mounts
}

// netDevicePool returns empty DeviceSpecs
func (rp *netDevicePool) GetDeviceSpecs(deviceFiles map[string]string, deviceIDs []string) []*pluginapi.DeviceSpec {
	glog.Infof("generic GetDeviceSpecs() called")
	devSpecs := make([]*pluginapi.DeviceSpec, 0)
	return devSpecs
}

// Probe returns 'true' if device health changes 'false' otherwise
func (rp *netDevicePool) Probe(rc *types.ResourceConfig, devices map[string]*pluginapi.Device) bool {
	return false
}
