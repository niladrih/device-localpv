/*
Copyright 2021 The OpenEBS Authors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	scpv1alpha1 "github.com/openebs/device-localpv/pkg/apis/openebs.io/scp/v1alpha1"
	internalclientset "github.com/openebs/device-localpv/pkg/generated/clientset/scp/internalclientset"
	internalinterfaces "github.com/openebs/device-localpv/pkg/generated/informer/scp/externalversions/internalinterfaces"
	v1alpha1 "github.com/openebs/device-localpv/pkg/generated/lister/scp/scp/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// StorageVolumeInformer provides access to a shared informer and lister for
// StorageVolumes.
type StorageVolumeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.StorageVolumeLister
}

type storageVolumeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewStorageVolumeInformer constructs a new informer for StorageVolume type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorageVolumeInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStorageVolumeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredStorageVolumeInformer constructs a new informer for StorageVolume type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStorageVolumeInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ScpV1alpha1().StorageVolumes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ScpV1alpha1().StorageVolumes(namespace).Watch(context.TODO(), options)
			},
		},
		&scpv1alpha1.StorageVolume{},
		resyncPeriod,
		indexers,
	)
}

func (f *storageVolumeInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStorageVolumeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *storageVolumeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&scpv1alpha1.StorageVolume{}, f.defaultInformer)
}

func (f *storageVolumeInformer) Lister() v1alpha1.StorageVolumeLister {
	return v1alpha1.NewStorageVolumeLister(f.Informer().GetIndexer())
}
