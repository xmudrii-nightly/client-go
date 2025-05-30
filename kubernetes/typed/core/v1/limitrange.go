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
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

// LimitRangesClusterGetter has a method to return a LimitRangeClusterInterface.
// A group's cluster client should implement this interface.
type LimitRangesClusterGetter interface {
	LimitRanges() LimitRangeClusterInterface
}

// LimitRangeClusterInterface can operate on LimitRanges across all clusters,
// or scope down to one cluster and return a LimitRangesNamespacer.
type LimitRangeClusterInterface interface {
	Cluster(logicalcluster.Path) LimitRangesNamespacer
	List(ctx context.Context, opts metav1.ListOptions) (*corev1.LimitRangeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type limitRangesClusterInterface struct {
	clientCache kcpclient.Cache[*corev1client.CoreV1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *limitRangesClusterInterface) Cluster(clusterPath logicalcluster.Path) LimitRangesNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &limitRangesNamespacer{clientCache: c.clientCache, clusterPath: clusterPath}
}

// List returns the entire collection of all LimitRanges across all clusters.
func (c *limitRangesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*corev1.LimitRangeList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).LimitRanges(metav1.NamespaceAll).List(ctx, opts)
}

// Watch begins to watch all LimitRanges across all clusters.
func (c *limitRangesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).LimitRanges(metav1.NamespaceAll).Watch(ctx, opts)
}

// LimitRangesNamespacer can scope to objects within a namespace, returning a corev1client.LimitRangeInterface.
type LimitRangesNamespacer interface {
	Namespace(string) corev1client.LimitRangeInterface
}

type limitRangesNamespacer struct {
	clientCache kcpclient.Cache[*corev1client.CoreV1Client]
	clusterPath logicalcluster.Path
}

func (n *limitRangesNamespacer) Namespace(namespace string) corev1client.LimitRangeInterface {
	return n.clientCache.ClusterOrDie(n.clusterPath).LimitRanges(namespace)
}
