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

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	corev1listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

// NodeClusterLister can list Nodes across all workspaces, or scope down to a NodeLister for one workspace.
// All objects returned here must be treated as read-only.
type NodeClusterLister interface {
	// List lists all Nodes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*corev1.Node, err error)
	// Cluster returns a lister that can list and get Nodes in one workspace.
	Cluster(clusterName logicalcluster.Name) corev1listers.NodeLister
	NodeClusterListerExpansion
}

type nodeClusterLister struct {
	indexer cache.Indexer
}

// NewNodeClusterLister returns a new NodeClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewNodeClusterLister(indexer cache.Indexer) *nodeClusterLister {
	return &nodeClusterLister{indexer: indexer}
}

// List lists all Nodes in the indexer across all workspaces.
func (s *nodeClusterLister) List(selector labels.Selector) (ret []*corev1.Node, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*corev1.Node))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get Nodes.
func (s *nodeClusterLister) Cluster(clusterName logicalcluster.Name) corev1listers.NodeLister {
	return &nodeLister{indexer: s.indexer, clusterName: clusterName}
}

// nodeLister implements the corev1listers.NodeLister interface.
type nodeLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all Nodes in the indexer for a workspace.
func (s *nodeLister) List(selector labels.Selector) (ret []*corev1.Node, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*corev1.Node))
	})
	return ret, err
}

// Get retrieves the Node from the indexer for a given workspace and name.
func (s *nodeLister) Get(name string) (*corev1.Node, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(corev1.Resource("nodes"), name)
	}
	return obj.(*corev1.Node), nil
}
