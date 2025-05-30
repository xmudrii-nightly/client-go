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

	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	applyconfigurationsautoscalingv1 "k8s.io/client-go/applyconfigurations/autoscaling/v1"
	autoscalingv1client "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	"k8s.io/client-go/testing"

	kcpautoscalingv1 "github.com/kcp-dev/client-go/kubernetes/typed/autoscaling/v1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var horizontalPodAutoscalersResource = schema.GroupVersionResource{Group: "autoscaling", Version: "v1", Resource: "horizontalpodautoscalers"}
var horizontalPodAutoscalersKind = schema.GroupVersionKind{Group: "autoscaling", Version: "v1", Kind: "HorizontalPodAutoscaler"}

type horizontalPodAutoscalersClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *horizontalPodAutoscalersClusterClient) Cluster(clusterPath logicalcluster.Path) kcpautoscalingv1.HorizontalPodAutoscalersNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &horizontalPodAutoscalersNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of HorizontalPodAutoscalers that match those selectors across all clusters.
func (c *horizontalPodAutoscalersClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*autoscalingv1.HorizontalPodAutoscalerList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(horizontalPodAutoscalersResource, horizontalPodAutoscalersKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &autoscalingv1.HorizontalPodAutoscalerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &autoscalingv1.HorizontalPodAutoscalerList{ListMeta: obj.(*autoscalingv1.HorizontalPodAutoscalerList).ListMeta}
	for _, item := range obj.(*autoscalingv1.HorizontalPodAutoscalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested HorizontalPodAutoscalers across all clusters.
func (c *horizontalPodAutoscalersClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(horizontalPodAutoscalersResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type horizontalPodAutoscalersNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *horizontalPodAutoscalersNamespacer) Namespace(namespace string) autoscalingv1client.HorizontalPodAutoscalerInterface {
	return &horizontalPodAutoscalersClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type horizontalPodAutoscalersClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *horizontalPodAutoscalersClient) Create(ctx context.Context, horizontalPodAutoscaler *autoscalingv1.HorizontalPodAutoscaler, opts metav1.CreateOptions) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(horizontalPodAutoscalersResource, c.ClusterPath, c.Namespace, horizontalPodAutoscaler), &autoscalingv1.HorizontalPodAutoscaler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.HorizontalPodAutoscaler), err
}

func (c *horizontalPodAutoscalersClient) Update(ctx context.Context, horizontalPodAutoscaler *autoscalingv1.HorizontalPodAutoscaler, opts metav1.UpdateOptions) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(horizontalPodAutoscalersResource, c.ClusterPath, c.Namespace, horizontalPodAutoscaler), &autoscalingv1.HorizontalPodAutoscaler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.HorizontalPodAutoscaler), err
}

func (c *horizontalPodAutoscalersClient) UpdateStatus(ctx context.Context, horizontalPodAutoscaler *autoscalingv1.HorizontalPodAutoscaler, opts metav1.UpdateOptions) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(horizontalPodAutoscalersResource, c.ClusterPath, "status", c.Namespace, horizontalPodAutoscaler), &autoscalingv1.HorizontalPodAutoscaler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.HorizontalPodAutoscaler), err
}

func (c *horizontalPodAutoscalersClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(horizontalPodAutoscalersResource, c.ClusterPath, c.Namespace, name, opts), &autoscalingv1.HorizontalPodAutoscaler{})
	return err
}

func (c *horizontalPodAutoscalersClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(horizontalPodAutoscalersResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &autoscalingv1.HorizontalPodAutoscalerList{})
	return err
}

func (c *horizontalPodAutoscalersClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(horizontalPodAutoscalersResource, c.ClusterPath, c.Namespace, name), &autoscalingv1.HorizontalPodAutoscaler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.HorizontalPodAutoscaler), err
}

// List takes label and field selectors, and returns the list of HorizontalPodAutoscalers that match those selectors.
func (c *horizontalPodAutoscalersClient) List(ctx context.Context, opts metav1.ListOptions) (*autoscalingv1.HorizontalPodAutoscalerList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(horizontalPodAutoscalersResource, horizontalPodAutoscalersKind, c.ClusterPath, c.Namespace, opts), &autoscalingv1.HorizontalPodAutoscalerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &autoscalingv1.HorizontalPodAutoscalerList{ListMeta: obj.(*autoscalingv1.HorizontalPodAutoscalerList).ListMeta}
	for _, item := range obj.(*autoscalingv1.HorizontalPodAutoscalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *horizontalPodAutoscalersClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(horizontalPodAutoscalersResource, c.ClusterPath, c.Namespace, opts))
}

func (c *horizontalPodAutoscalersClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*autoscalingv1.HorizontalPodAutoscaler, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(horizontalPodAutoscalersResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &autoscalingv1.HorizontalPodAutoscaler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.HorizontalPodAutoscaler), err
}

func (c *horizontalPodAutoscalersClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsautoscalingv1.HorizontalPodAutoscalerApplyConfiguration, opts metav1.ApplyOptions) (*autoscalingv1.HorizontalPodAutoscaler, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(horizontalPodAutoscalersResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &autoscalingv1.HorizontalPodAutoscaler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.HorizontalPodAutoscaler), err
}

func (c *horizontalPodAutoscalersClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsautoscalingv1.HorizontalPodAutoscalerApplyConfiguration, opts metav1.ApplyOptions) (*autoscalingv1.HorizontalPodAutoscaler, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(horizontalPodAutoscalersResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data, "status"), &autoscalingv1.HorizontalPodAutoscaler{})
	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.HorizontalPodAutoscaler), err
}
