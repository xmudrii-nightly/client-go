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
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"

	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsadmissionregistrationv1beta1 "k8s.io/client-go/applyconfigurations/admissionregistration/v1beta1"
	admissionregistrationv1beta1client "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var validatingAdmissionPolicyBindingsResource = schema.GroupVersionResource{Group: "admissionregistration.k8s.io", Version: "v1beta1", Resource: "validatingadmissionpolicybindings"}
var validatingAdmissionPolicyBindingsKind = schema.GroupVersionKind{Group: "admissionregistration.k8s.io", Version: "v1beta1", Kind: "ValidatingAdmissionPolicyBinding"}

type validatingAdmissionPolicyBindingsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *validatingAdmissionPolicyBindingsClusterClient) Cluster(clusterPath logicalcluster.Path) admissionregistrationv1beta1client.ValidatingAdmissionPolicyBindingInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &validatingAdmissionPolicyBindingsClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ValidatingAdmissionPolicyBindings that match those selectors across all clusters.
func (c *validatingAdmissionPolicyBindingsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*admissionregistrationv1beta1.ValidatingAdmissionPolicyBindingList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(validatingAdmissionPolicyBindingsResource, validatingAdmissionPolicyBindingsKind, logicalcluster.Wildcard, opts), &admissionregistrationv1beta1.ValidatingAdmissionPolicyBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &admissionregistrationv1beta1.ValidatingAdmissionPolicyBindingList{ListMeta: obj.(*admissionregistrationv1beta1.ValidatingAdmissionPolicyBindingList).ListMeta}
	for _, item := range obj.(*admissionregistrationv1beta1.ValidatingAdmissionPolicyBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ValidatingAdmissionPolicyBindings across all clusters.
func (c *validatingAdmissionPolicyBindingsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(validatingAdmissionPolicyBindingsResource, logicalcluster.Wildcard, opts))
}

type validatingAdmissionPolicyBindingsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *validatingAdmissionPolicyBindingsClient) Create(ctx context.Context, validatingAdmissionPolicyBinding *admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding, opts metav1.CreateOptions) (*admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(validatingAdmissionPolicyBindingsResource, c.ClusterPath, validatingAdmissionPolicyBinding), &admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding), err
}

func (c *validatingAdmissionPolicyBindingsClient) Update(ctx context.Context, validatingAdmissionPolicyBinding *admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding, opts metav1.UpdateOptions) (*admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(validatingAdmissionPolicyBindingsResource, c.ClusterPath, validatingAdmissionPolicyBinding), &admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding), err
}

func (c *validatingAdmissionPolicyBindingsClient) UpdateStatus(ctx context.Context, validatingAdmissionPolicyBinding *admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding, opts metav1.UpdateOptions) (*admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(validatingAdmissionPolicyBindingsResource, c.ClusterPath, "status", validatingAdmissionPolicyBinding), &admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding), err
}

func (c *validatingAdmissionPolicyBindingsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(validatingAdmissionPolicyBindingsResource, c.ClusterPath, name, opts), &admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding{})
	return err
}

func (c *validatingAdmissionPolicyBindingsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(validatingAdmissionPolicyBindingsResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &admissionregistrationv1beta1.ValidatingAdmissionPolicyBindingList{})
	return err
}

func (c *validatingAdmissionPolicyBindingsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(validatingAdmissionPolicyBindingsResource, c.ClusterPath, name), &admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding), err
}

// List takes label and field selectors, and returns the list of ValidatingAdmissionPolicyBindings that match those selectors.
func (c *validatingAdmissionPolicyBindingsClient) List(ctx context.Context, opts metav1.ListOptions) (*admissionregistrationv1beta1.ValidatingAdmissionPolicyBindingList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(validatingAdmissionPolicyBindingsResource, validatingAdmissionPolicyBindingsKind, c.ClusterPath, opts), &admissionregistrationv1beta1.ValidatingAdmissionPolicyBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &admissionregistrationv1beta1.ValidatingAdmissionPolicyBindingList{ListMeta: obj.(*admissionregistrationv1beta1.ValidatingAdmissionPolicyBindingList).ListMeta}
	for _, item := range obj.(*admissionregistrationv1beta1.ValidatingAdmissionPolicyBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *validatingAdmissionPolicyBindingsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(validatingAdmissionPolicyBindingsResource, c.ClusterPath, opts))
}

func (c *validatingAdmissionPolicyBindingsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(validatingAdmissionPolicyBindingsResource, c.ClusterPath, name, pt, data, subresources...), &admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding), err
}

func (c *validatingAdmissionPolicyBindingsClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsadmissionregistrationv1beta1.ValidatingAdmissionPolicyBindingApplyConfiguration, opts metav1.ApplyOptions) (*admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(validatingAdmissionPolicyBindingsResource, c.ClusterPath, *name, types.ApplyPatchType, data), &admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding), err
}

func (c *validatingAdmissionPolicyBindingsClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsadmissionregistrationv1beta1.ValidatingAdmissionPolicyBindingApplyConfiguration, opts metav1.ApplyOptions) (*admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(validatingAdmissionPolicyBindingsResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*admissionregistrationv1beta1.ValidatingAdmissionPolicyBinding), err
}
