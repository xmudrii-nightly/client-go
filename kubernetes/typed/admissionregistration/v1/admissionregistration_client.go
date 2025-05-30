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

package v1

import (
	"net/http"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	admissionregistrationv1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	"k8s.io/client-go/rest"
)

type AdmissionregistrationV1ClusterInterface interface {
	AdmissionregistrationV1ClusterScoper
	ValidatingAdmissionPoliciesClusterGetter
	ValidatingAdmissionPolicyBindingsClusterGetter
	ValidatingWebhookConfigurationsClusterGetter
	MutatingWebhookConfigurationsClusterGetter
}

type AdmissionregistrationV1ClusterScoper interface {
	Cluster(logicalcluster.Path) admissionregistrationv1.AdmissionregistrationV1Interface
}

type AdmissionregistrationV1ClusterClient struct {
	clientCache kcpclient.Cache[*admissionregistrationv1.AdmissionregistrationV1Client]
}

func (c *AdmissionregistrationV1ClusterClient) Cluster(clusterPath logicalcluster.Path) admissionregistrationv1.AdmissionregistrationV1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return c.clientCache.ClusterOrDie(clusterPath)
}

func (c *AdmissionregistrationV1ClusterClient) ValidatingAdmissionPolicies() ValidatingAdmissionPolicyClusterInterface {
	return &validatingAdmissionPoliciesClusterInterface{clientCache: c.clientCache}
}

func (c *AdmissionregistrationV1ClusterClient) ValidatingAdmissionPolicyBindings() ValidatingAdmissionPolicyBindingClusterInterface {
	return &validatingAdmissionPolicyBindingsClusterInterface{clientCache: c.clientCache}
}

func (c *AdmissionregistrationV1ClusterClient) ValidatingWebhookConfigurations() ValidatingWebhookConfigurationClusterInterface {
	return &validatingWebhookConfigurationsClusterInterface{clientCache: c.clientCache}
}

func (c *AdmissionregistrationV1ClusterClient) MutatingWebhookConfigurations() MutatingWebhookConfigurationClusterInterface {
	return &mutatingWebhookConfigurationsClusterInterface{clientCache: c.clientCache}
}

// NewForConfig creates a new AdmissionregistrationV1ClusterClient for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*AdmissionregistrationV1ClusterClient, error) {
	client, err := rest.HTTPClientFor(c)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(c, client)
}

// NewForConfigAndClient creates a new AdmissionregistrationV1ClusterClient for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*AdmissionregistrationV1ClusterClient, error) {
	cache := kcpclient.NewCache(c, h, &kcpclient.Constructor[*admissionregistrationv1.AdmissionregistrationV1Client]{
		NewForConfigAndClient: admissionregistrationv1.NewForConfigAndClient,
	})
	if _, err := cache.Cluster(logicalcluster.Name("root").Path()); err != nil {
		return nil, err
	}
	return &AdmissionregistrationV1ClusterClient{clientCache: cache}, nil
}

// NewForConfigOrDie creates a new AdmissionregistrationV1ClusterClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *AdmissionregistrationV1ClusterClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}
