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

package v1alpha2

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	coordinationv1alpha2 "k8s.io/api/coordination/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	coordinationv1alpha2listers "k8s.io/client-go/listers/coordination/v1alpha2"
	"k8s.io/client-go/tools/cache"
)

// LeaseCandidateClusterLister can list LeaseCandidates across all workspaces, or scope down to a LeaseCandidateLister for one workspace.
// All objects returned here must be treated as read-only.
type LeaseCandidateClusterLister interface {
	// List lists all LeaseCandidates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*coordinationv1alpha2.LeaseCandidate, err error)
	// Cluster returns a lister that can list and get LeaseCandidates in one workspace.
	Cluster(clusterName logicalcluster.Name) coordinationv1alpha2listers.LeaseCandidateLister
	LeaseCandidateClusterListerExpansion
}

type leaseCandidateClusterLister struct {
	indexer cache.Indexer
}

// NewLeaseCandidateClusterLister returns a new LeaseCandidateClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
// - has the kcpcache.ClusterAndNamespaceIndex as an index
func NewLeaseCandidateClusterLister(indexer cache.Indexer) *leaseCandidateClusterLister {
	return &leaseCandidateClusterLister{indexer: indexer}
}

// List lists all LeaseCandidates in the indexer across all workspaces.
func (s *leaseCandidateClusterLister) List(selector labels.Selector) (ret []*coordinationv1alpha2.LeaseCandidate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*coordinationv1alpha2.LeaseCandidate))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get LeaseCandidates.
func (s *leaseCandidateClusterLister) Cluster(clusterName logicalcluster.Name) coordinationv1alpha2listers.LeaseCandidateLister {
	return &leaseCandidateLister{indexer: s.indexer, clusterName: clusterName}
}

// leaseCandidateLister implements the coordinationv1alpha2listers.LeaseCandidateLister interface.
type leaseCandidateLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all LeaseCandidates in the indexer for a workspace.
func (s *leaseCandidateLister) List(selector labels.Selector) (ret []*coordinationv1alpha2.LeaseCandidate, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*coordinationv1alpha2.LeaseCandidate))
	})
	return ret, err
}

// LeaseCandidates returns an object that can list and get LeaseCandidates in one namespace.
func (s *leaseCandidateLister) LeaseCandidates(namespace string) coordinationv1alpha2listers.LeaseCandidateNamespaceLister {
	return &leaseCandidateNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// leaseCandidateNamespaceLister implements the coordinationv1alpha2listers.LeaseCandidateNamespaceLister interface.
type leaseCandidateNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// List lists all LeaseCandidates in the indexer for a given workspace and namespace.
func (s *leaseCandidateNamespaceLister) List(selector labels.Selector) (ret []*coordinationv1alpha2.LeaseCandidate, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*coordinationv1alpha2.LeaseCandidate))
	})
	return ret, err
}

// Get retrieves the LeaseCandidate from the indexer for a given workspace, namespace and name.
func (s *leaseCandidateNamespaceLister) Get(name string) (*coordinationv1alpha2.LeaseCandidate, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(coordinationv1alpha2.Resource("leasecandidates"), name)
	}
	return obj.(*coordinationv1alpha2.LeaseCandidate), nil
}
