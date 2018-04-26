// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This config yaml file is loaded into a data structure and pulled in by the
// controller.

package controller

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
)

// SfcConfigYaml is container struct for yaml config file
type SfcConfigYaml struct {
	Version             int                    `json:"sfc_controller_config_version"`
	Description         string                 `json:"description"`
	SysParms            SystemParameters      `json:"system_parameters"`
	IPAMPools           []*IPAMPool            `json:"ipam_pools"`
	NetworkPodToNodeMap []*NetworkPodToNodeMap `json:"network_pod_to_node_map"`
	NetworkNodeOverlays []*NetworkNodeOverlay  `json:"network_node_overlays"`
	NetworkNodes        []*NetworkNode         `json:"network_nodes"`
	NetworkServices     []*NetworkService      `json:"network_services"`
}

// SfcConfigYamlReadFromFile parses the yaml into YamlConfig
func (s *Plugin) SfcConfigYamlReadFromFile(fpath string) (*SfcConfigYaml, error) {

	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	yamlConfig := &SfcConfigYaml{}

	if err := yaml.Unmarshal(b, yamlConfig); err != nil {
		return nil, err
	}

	return yamlConfig, nil
}

// SfcConfigYamlProcessConfig processes each object and adds it to the system
func (s *Plugin) SfcConfigYamlProcessConfig(y *SfcConfigYaml) error {

	if y.Version != 2 {
		return fmt.Errorf("SfcConfigYamlProcessConfig: incorrect yaml version, expecting 2, got: %d",
			y.Version)
	}

	log.Debugf("SfcConfigYamlProcessConfig: system parameters: ", y.SysParms)
	if err := ctlrPlugin.SysParametersMgr.HandleCRUDOperationCU(&y.SysParms, false); err != nil {
		return err
	}

	log.Debugf("SfcConfigYamlProcessConfig: ipam pools: ", y.IPAMPools)
	for _, ipamPool := range y.IPAMPools {
		if err := ctlrPlugin.IpamPoolMgr.HandleCRUDOperationCU(ipamPool, false); err != nil {
			return err
		}
	}

	for _, nn := range y.NetworkNodes {
		log.Debugf("SfcConfigYamlProcessConfig: network node: ", nn)
		if err := ctlrPlugin.NetworkNodeMgr.HandleCRUDOperationCU(nn, false); err != nil {
			return err
		}
	}

	for _, ns := range y.NetworkServices {
		log.Debugf("SfcConfigYamlProcessConfig: network-service: ", ns)
		if err := ctlrPlugin.NetworkServiceMgr.HandleCRUDOperationCU(ns, false); err != nil {
			return err
		}
	}

	for _, p2n := range y.NetworkPodToNodeMap {
		log.Debugf("SfcConfigYamlProcessConfig: network-pod-to-node-map: ", p2n)
		if err := ctlrPlugin.NetworkPodNodeMapMgr.HandleCRUDOperationCU(p2n, false); err != nil {
			return err
		}
	}

	log.Debugf("SfcConfigYamlProcessConfig: network-node-overlays: ", y.NetworkNodeOverlays)
	for _, nno := range y.NetworkNodeOverlays {
		if err := ctlrPlugin.NetworkNodeOverlayMgr.HandleCRUDOperationCU(nno, false); err != nil {
			return err
		}
	}

	return nil
}