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

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	appsv1listers "k8s.io/client-go/listers/apps/v1"
	"k8s.io/client-go/tools/cache"
)

// ReplicaSetClusterLister can list ReplicaSets across all workspaces, or scope down to a ReplicaSetLister for one workspace.
// All objects returned here must be treated as read-only.
type ReplicaSetClusterLister interface {
	// List lists all ReplicaSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*appsv1.ReplicaSet, err error)
	// Cluster returns a lister that can list and get ReplicaSets in one workspace.
	Cluster(clusterName logicalcluster.Name) appsv1listers.ReplicaSetLister
	ReplicaSetClusterListerExpansion
}

type replicaSetClusterLister struct {
	indexer cache.Indexer
}

// NewReplicaSetClusterLister returns a new ReplicaSetClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewReplicaSetClusterLister(indexer cache.Indexer) *replicaSetClusterLister {
	return &replicaSetClusterLister{indexer: indexer}
}

// List lists all ReplicaSets in the indexer across all workspaces.
func (s *replicaSetClusterLister) List(selector labels.Selector) (ret []*appsv1.ReplicaSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*appsv1.ReplicaSet))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get ReplicaSets.
func (s *replicaSetClusterLister) Cluster(clusterName logicalcluster.Name) appsv1listers.ReplicaSetLister {
	return &replicaSetLister{indexer: s.indexer, clusterName: clusterName}
}

// replicaSetLister implements the appsv1listers.ReplicaSetLister interface.
type replicaSetLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all ReplicaSets in the indexer for a workspace.
func (s *replicaSetLister) List(selector labels.Selector) (ret []*appsv1.ReplicaSet, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*appsv1.ReplicaSet))
	})
	return ret, err
}

// ReplicaSets returns an object that can list and get ReplicaSets in one namespace.
func (s *replicaSetLister) ReplicaSets(namespace string) appsv1listers.ReplicaSetNamespaceLister {
	return &replicaSetNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// replicaSetNamespaceLister implements the appsv1listers.ReplicaSetNamespaceLister interface.
type replicaSetNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// List lists all ReplicaSets in the indexer for a given workspace and namespace.
func (s *replicaSetNamespaceLister) List(selector labels.Selector) (ret []*appsv1.ReplicaSet, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*appsv1.ReplicaSet))
	})
	return ret, err
}

// Get retrieves the ReplicaSet from the indexer for a given workspace, namespace and name.
func (s *replicaSetNamespaceLister) Get(name string) (*appsv1.ReplicaSet, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(appsv1.Resource("replicasets"), name)
	}
	return obj.(*appsv1.ReplicaSet), nil
}
