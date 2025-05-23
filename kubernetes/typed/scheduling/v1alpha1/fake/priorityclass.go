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

	schedulingv1alpha1 "k8s.io/api/scheduling/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsschedulingv1alpha1 "k8s.io/client-go/applyconfigurations/scheduling/v1alpha1"
	schedulingv1alpha1client "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	"k8s.io/client-go/testing"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var priorityClassesResource = schema.GroupVersionResource{Group: "scheduling.k8s.io", Version: "v1alpha1", Resource: "priorityclasses"}
var priorityClassesKind = schema.GroupVersionKind{Group: "scheduling.k8s.io", Version: "v1alpha1", Kind: "PriorityClass"}

type priorityClassesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *priorityClassesClusterClient) Cluster(clusterPath logicalcluster.Path) schedulingv1alpha1client.PriorityClassInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &priorityClassesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of PriorityClasses that match those selectors across all clusters.
func (c *priorityClassesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*schedulingv1alpha1.PriorityClassList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(priorityClassesResource, priorityClassesKind, logicalcluster.Wildcard, opts), &schedulingv1alpha1.PriorityClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &schedulingv1alpha1.PriorityClassList{ListMeta: obj.(*schedulingv1alpha1.PriorityClassList).ListMeta}
	for _, item := range obj.(*schedulingv1alpha1.PriorityClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested PriorityClasses across all clusters.
func (c *priorityClassesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(priorityClassesResource, logicalcluster.Wildcard, opts))
}

type priorityClassesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *priorityClassesClient) Create(ctx context.Context, priorityClass *schedulingv1alpha1.PriorityClass, opts metav1.CreateOptions) (*schedulingv1alpha1.PriorityClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(priorityClassesResource, c.ClusterPath, priorityClass), &schedulingv1alpha1.PriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*schedulingv1alpha1.PriorityClass), err
}

func (c *priorityClassesClient) Update(ctx context.Context, priorityClass *schedulingv1alpha1.PriorityClass, opts metav1.UpdateOptions) (*schedulingv1alpha1.PriorityClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(priorityClassesResource, c.ClusterPath, priorityClass), &schedulingv1alpha1.PriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*schedulingv1alpha1.PriorityClass), err
}

func (c *priorityClassesClient) UpdateStatus(ctx context.Context, priorityClass *schedulingv1alpha1.PriorityClass, opts metav1.UpdateOptions) (*schedulingv1alpha1.PriorityClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(priorityClassesResource, c.ClusterPath, "status", priorityClass), &schedulingv1alpha1.PriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*schedulingv1alpha1.PriorityClass), err
}

func (c *priorityClassesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(priorityClassesResource, c.ClusterPath, name, opts), &schedulingv1alpha1.PriorityClass{})
	return err
}

func (c *priorityClassesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(priorityClassesResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &schedulingv1alpha1.PriorityClassList{})
	return err
}

func (c *priorityClassesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*schedulingv1alpha1.PriorityClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(priorityClassesResource, c.ClusterPath, name), &schedulingv1alpha1.PriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*schedulingv1alpha1.PriorityClass), err
}

// List takes label and field selectors, and returns the list of PriorityClasses that match those selectors.
func (c *priorityClassesClient) List(ctx context.Context, opts metav1.ListOptions) (*schedulingv1alpha1.PriorityClassList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(priorityClassesResource, priorityClassesKind, c.ClusterPath, opts), &schedulingv1alpha1.PriorityClassList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &schedulingv1alpha1.PriorityClassList{ListMeta: obj.(*schedulingv1alpha1.PriorityClassList).ListMeta}
	for _, item := range obj.(*schedulingv1alpha1.PriorityClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *priorityClassesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(priorityClassesResource, c.ClusterPath, opts))
}

func (c *priorityClassesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*schedulingv1alpha1.PriorityClass, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(priorityClassesResource, c.ClusterPath, name, pt, data, subresources...), &schedulingv1alpha1.PriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*schedulingv1alpha1.PriorityClass), err
}

func (c *priorityClassesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsschedulingv1alpha1.PriorityClassApplyConfiguration, opts metav1.ApplyOptions) (*schedulingv1alpha1.PriorityClass, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(priorityClassesResource, c.ClusterPath, *name, types.ApplyPatchType, data), &schedulingv1alpha1.PriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*schedulingv1alpha1.PriorityClass), err
}

func (c *priorityClassesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsschedulingv1alpha1.PriorityClassApplyConfiguration, opts metav1.ApplyOptions) (*schedulingv1alpha1.PriorityClass, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(priorityClassesResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &schedulingv1alpha1.PriorityClass{})
	if obj == nil {
		return nil, err
	}
	return obj.(*schedulingv1alpha1.PriorityClass), err
}
