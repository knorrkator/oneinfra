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

package cluster

import (
	"github.com/oneinfra/oneinfra/internal/pkg/certificates"
)

// EtcdServer represents the etcd component
type EtcdServer struct {
	CA            *certificates.Certificate
	TLSCert       string
	TLSPrivateKey string
	ExtraSANs     []string
}

func newEtcdServer(etcdServerExtraSANs []string) (*EtcdServer, error) {
	certificateAuthority, err := certificates.NewCertificateAuthority("etcd-authority")
	if err != nil {
		return nil, err
	}
	etcdServer := EtcdServer{
		CA:        certificateAuthority,
		ExtraSANs: etcdServerExtraSANs,
	}
	tlsCert, tlsKey, err := certificateAuthority.CreateCertificate(
		"etcd-server",
		[]string{"etcd-server"},
		etcdServerExtraSANs,
	)
	if err != nil {
		return nil, err
	}
	etcdServer.TLSCert = tlsCert
	etcdServer.TLSPrivateKey = tlsKey
	return &etcdServer, nil
}
