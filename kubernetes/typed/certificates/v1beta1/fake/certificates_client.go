/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package fake

import (
	"github.com/kcp-dev/logicalcluster/v3"

	certificatesv1beta1 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	"k8s.io/client-go/rest"

	kcpcertificatesv1beta1 "github.com/kcp-dev/client-go/kubernetes/typed/certificates/v1beta1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpcertificatesv1beta1.CertificatesV1beta1ClusterInterface = (*CertificatesV1beta1ClusterClient)(nil)

type CertificatesV1beta1ClusterClient struct {
	*kcptesting.Fake
}

func (c *CertificatesV1beta1ClusterClient) Cluster(clusterPath logicalcluster.Path) certificatesv1beta1.CertificatesV1beta1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &CertificatesV1beta1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *CertificatesV1beta1ClusterClient) CertificateSigningRequests() kcpcertificatesv1beta1.CertificateSigningRequestClusterInterface {
	return &certificateSigningRequestsClusterClient{Fake: c.Fake}
}

var _ certificatesv1beta1.CertificatesV1beta1Interface = (*CertificatesV1beta1Client)(nil)

type CertificatesV1beta1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *CertificatesV1beta1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *CertificatesV1beta1Client) CertificateSigningRequests() certificatesv1beta1.CertificateSigningRequestInterface {
	return &certificateSigningRequestsClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}
