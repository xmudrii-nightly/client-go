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
	"github.com/kcp-dev/logicalcluster/v3"

	storagev1alpha1 "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
	"k8s.io/client-go/rest"

	kcpstoragev1alpha1 "github.com/kcp-dev/client-go/kubernetes/typed/storage/v1alpha1"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
)

var _ kcpstoragev1alpha1.StorageV1alpha1ClusterInterface = (*StorageV1alpha1ClusterClient)(nil)

type StorageV1alpha1ClusterClient struct {
	*kcptesting.Fake
}

func (c *StorageV1alpha1ClusterClient) Cluster(clusterPath logicalcluster.Path) storagev1alpha1.StorageV1alpha1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &StorageV1alpha1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *StorageV1alpha1ClusterClient) VolumeAttachments() kcpstoragev1alpha1.VolumeAttachmentClusterInterface {
	return &volumeAttachmentsClusterClient{Fake: c.Fake}
}

func (c *StorageV1alpha1ClusterClient) CSIStorageCapacities() kcpstoragev1alpha1.CSIStorageCapacityClusterInterface {
	return &cSIStorageCapacitiesClusterClient{Fake: c.Fake}
}

func (c *StorageV1alpha1ClusterClient) VolumeAttributesClasses() kcpstoragev1alpha1.VolumeAttributesClassClusterInterface {
	return &volumeAttributesClassesClusterClient{Fake: c.Fake}
}

var _ storagev1alpha1.StorageV1alpha1Interface = (*StorageV1alpha1Client)(nil)

type StorageV1alpha1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *StorageV1alpha1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *StorageV1alpha1Client) VolumeAttachments() storagev1alpha1.VolumeAttachmentInterface {
	return &volumeAttachmentsClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}

func (c *StorageV1alpha1Client) CSIStorageCapacities(namespace string) storagev1alpha1.CSIStorageCapacityInterface {
	return &cSIStorageCapacitiesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *StorageV1alpha1Client) VolumeAttributesClasses() storagev1alpha1.VolumeAttributesClassInterface {
	return &volumeAttributesClassesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}
