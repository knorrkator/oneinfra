/*
Copyright 2020 Rafael Fernández López <ereslibre@ereslibre.es>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package constants

import (
	"path/filepath"
)

const (
	// DefaultKeyBitSize is the default key bit size
	DefaultKeyBitSize = 1024
	// OneInfraNamespace is the namespace for storing OneInfra resources
	OneInfraNamespace = "oneinfra-system"
	// OneInfraNodeJoinTokenExtraGroups represents the bootstrap token
	// extra groups used to identify oneinfra bootstrap tokens
	OneInfraNodeJoinTokenExtraGroups = "system:bootstrappers:oneinfra"
	// OneInfraConfigDir represents the configuration directory for oneinfra
	OneInfraConfigDir = "/etc/oneinfra"
	// OneInfraControlPlaneIngressVPNPeerName represents the control plane
	// ingress peer VPN name
	OneInfraControlPlaneIngressVPNPeerName = "control-plane-ingress"
	// KubeletDir is the kubelet configuration dir
	KubeletDir = "/var/lib/kubelet"
)

var (
	// KubeletKubeConfigPath represents the kubelet kubeconfig path
	KubeletKubeConfigPath = filepath.Join(OneInfraConfigDir, "kubelet.conf")
	// KubeletConfigPath represents the kubelet configuration path
	KubeletConfigPath = filepath.Join(KubeletDir, "config.yaml")
)
