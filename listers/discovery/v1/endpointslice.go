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
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	discoveryv1 "k8s.io/api/discovery/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	discoveryv1listers "k8s.io/client-go/listers/discovery/v1"
	"k8s.io/client-go/tools/cache"
)

// EndpointSliceClusterLister can list EndpointSlices across all workspaces, or scope down to a EndpointSliceLister for one workspace.
// All objects returned here must be treated as read-only.
type EndpointSliceClusterLister interface {
	// List lists all EndpointSlices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*discoveryv1.EndpointSlice, err error)
	// Cluster returns a lister that can list and get EndpointSlices in one workspace.
	Cluster(clusterName logicalcluster.Name) discoveryv1listers.EndpointSliceLister
	EndpointSliceClusterListerExpansion
}

type endpointSliceClusterLister struct {
	indexer cache.Indexer
}

// NewEndpointSliceClusterLister returns a new EndpointSliceClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewEndpointSliceClusterLister(indexer cache.Indexer) *endpointSliceClusterLister {
	return &endpointSliceClusterLister{indexer: indexer}
}

// List lists all EndpointSlices in the indexer across all workspaces.
func (s *endpointSliceClusterLister) List(selector labels.Selector) (ret []*discoveryv1.EndpointSlice, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*discoveryv1.EndpointSlice))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get EndpointSlices.
func (s *endpointSliceClusterLister) Cluster(clusterName logicalcluster.Name) discoveryv1listers.EndpointSliceLister {
	return &endpointSliceLister{indexer: s.indexer, clusterName: clusterName}
}

// endpointSliceLister implements the discoveryv1listers.EndpointSliceLister interface.
type endpointSliceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all EndpointSlices in the indexer for a workspace.
func (s *endpointSliceLister) List(selector labels.Selector) (ret []*discoveryv1.EndpointSlice, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*discoveryv1.EndpointSlice))
	})
	return ret, err
}

// EndpointSlices returns an object that can list and get EndpointSlices in one namespace.
func (s *endpointSliceLister) EndpointSlices(namespace string) discoveryv1listers.EndpointSliceNamespaceLister {
	return &endpointSliceNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// endpointSliceNamespaceLister implements the discoveryv1listers.EndpointSliceNamespaceLister interface.
type endpointSliceNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// List lists all EndpointSlices in the indexer for a given workspace and namespace.
func (s *endpointSliceNamespaceLister) List(selector labels.Selector) (ret []*discoveryv1.EndpointSlice, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*discoveryv1.EndpointSlice))
	})
	return ret, err
}

// Get retrieves the EndpointSlice from the indexer for a given workspace, namespace and name.
func (s *endpointSliceNamespaceLister) Get(name string) (*discoveryv1.EndpointSlice, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(discoveryv1.Resource("endpointslices"), name)
	}
	return obj.(*discoveryv1.EndpointSlice), nil
}
