/*
Copyright The Red Hat Authors.
https://radanalytics.io/

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

// This file was automatically generated by informer-gen

package v1

import (
	radanalytics_io_v1 "github.com/radanalyticsio/oshinko-cli/pkg/apis/radanalytics.io/v1"
	versioned "github.com/radanalyticsio/oshinko-cli/pkg/client/clientset/versioned"
	internalinterfaces "github.com/radanalyticsio/oshinko-cli/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/radanalyticsio/oshinko-cli/pkg/client/listers/radanalytics.io/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// SparkClusterInformer provides access to a shared informer and lister for
// SparkClusters.
type SparkClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SparkClusterLister
}

type sparkClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSparkClusterInformer constructs a new informer for SparkCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSparkClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSparkClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSparkClusterInformer constructs a new informer for SparkCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSparkClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RadanalyticsV1().SparkClusters(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RadanalyticsV1().SparkClusters(namespace).Watch(options)
			},
		},
		&radanalytics_io_v1.SparkCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *sparkClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSparkClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sparkClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&radanalytics_io_v1.SparkCluster{}, f.defaultInformer)
}

func (f *sparkClusterInformer) Lister() v1.SparkClusterLister {
	return v1.NewSparkClusterLister(f.Informer().GetIndexer())
}
