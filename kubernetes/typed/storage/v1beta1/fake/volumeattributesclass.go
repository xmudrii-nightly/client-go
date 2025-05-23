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

	storagev1beta1 "k8s.io/api/storage/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsstoragev1beta1 "k8s.io/client-go/applyconfigurations/storage/v1beta1"
	storagev1beta1client "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var volumeAttributesClassesResource = schema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1beta1", Resource: "volumeattributesclasses"}
var volumeAttributesClassesKind = schema.GroupVersionKind{Group: "storage.k8s.io", Version: "v1beta1", Kind: "VolumeAttributesClass"}

type volumeAttributesClassesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *volumeAttributesClassesClusterClient) Cluster(clusterPath logicalcluster.Path) storagev1beta1client.VolumeAttributesClassInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &volumeAttributesClassesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of VolumeAttributesClasses that match those selectors across all clusters.
func (c *volumeAttributesClassesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*storagev1beta1.VolumeAttributesClassList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(volumeAttributesClassesResource, volumeAttributesClassesKind, logicalcluster.Wildcard, opts), &storagev1beta1.VolumeAttributesClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1beta1.VolumeAttributesClassList{ListMeta: obj.(*storagev1beta1.VolumeAttributesClassList).ListMeta}
	for _, item := range obj.(*storagev1beta1.VolumeAttributesClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested VolumeAttributesClasses across all clusters.
func (c *volumeAttributesClassesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(volumeAttributesClassesResource, logicalcluster.Wildcard, opts))
}

type volumeAttributesClassesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *volumeAttributesClassesClient) Create(ctx context.Context, volumeAttributesClass *storagev1beta1.VolumeAttributesClass, opts metav1.CreateOptions) (*storagev1beta1.VolumeAttributesClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(volumeAttributesClassesResource, c.ClusterPath, volumeAttributesClass), &storagev1beta1.VolumeAttributesClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.VolumeAttributesClass), err
}

func (c *volumeAttributesClassesClient) Update(ctx context.Context, volumeAttributesClass *storagev1beta1.VolumeAttributesClass, opts metav1.UpdateOptions) (*storagev1beta1.VolumeAttributesClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(volumeAttributesClassesResource, c.ClusterPath, volumeAttributesClass), &storagev1beta1.VolumeAttributesClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.VolumeAttributesClass), err
}

func (c *volumeAttributesClassesClient) UpdateStatus(ctx context.Context, volumeAttributesClass *storagev1beta1.VolumeAttributesClass, opts metav1.UpdateOptions) (*storagev1beta1.VolumeAttributesClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(volumeAttributesClassesResource, c.ClusterPath, "status", volumeAttributesClass), &storagev1beta1.VolumeAttributesClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.VolumeAttributesClass), err
}

func (c *volumeAttributesClassesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(volumeAttributesClassesResource, c.ClusterPath, name, opts), &storagev1beta1.VolumeAttributesClass{})
	return err
}

func (c *volumeAttributesClassesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(volumeAttributesClassesResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &storagev1beta1.VolumeAttributesClassList{})
	return err
}

func (c *volumeAttributesClassesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*storagev1beta1.VolumeAttributesClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(volumeAttributesClassesResource, c.ClusterPath, name), &storagev1beta1.VolumeAttributesClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.VolumeAttributesClass), err
}

// List takes label and field selectors, and returns the list of VolumeAttributesClasses that match those selectors.
func (c *volumeAttributesClassesClient) List(ctx context.Context, opts metav1.ListOptions) (*storagev1beta1.VolumeAttributesClassList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(volumeAttributesClassesResource, volumeAttributesClassesKind, c.ClusterPath, opts), &storagev1beta1.VolumeAttributesClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &storagev1beta1.VolumeAttributesClassList{ListMeta: obj.(*storagev1beta1.VolumeAttributesClassList).ListMeta}
	for _, item := range obj.(*storagev1beta1.VolumeAttributesClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *volumeAttributesClassesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(volumeAttributesClassesResource, c.ClusterPath, opts))
}

func (c *volumeAttributesClassesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*storagev1beta1.VolumeAttributesClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(volumeAttributesClassesResource, c.ClusterPath, name, pt, data, subresources...), &storagev1beta1.VolumeAttributesClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.VolumeAttributesClass), err
}

func (c *volumeAttributesClassesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsstoragev1beta1.VolumeAttributesClassApplyConfiguration, opts metav1.ApplyOptions) (*storagev1beta1.VolumeAttributesClass, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(volumeAttributesClassesResource, c.ClusterPath, *name, types.ApplyPatchType, data), &storagev1beta1.VolumeAttributesClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.VolumeAttributesClass), err
}

func (c *volumeAttributesClassesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsstoragev1beta1.VolumeAttributesClassApplyConfiguration, opts metav1.ApplyOptions) (*storagev1beta1.VolumeAttributesClass, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(volumeAttributesClassesResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &storagev1beta1.VolumeAttributesClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*storagev1beta1.VolumeAttributesClass), err
}
