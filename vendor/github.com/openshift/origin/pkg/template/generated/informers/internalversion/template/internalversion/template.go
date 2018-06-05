// This file was automatically generated by informer-gen

package internalversion

import (
	template "github.com/openshift/origin/pkg/template/apis/template"
	internalinterfaces "github.com/openshift/origin/pkg/template/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/template/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/template/generated/listers/template/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// TemplateInformer provides access to a shared informer and lister for
// Templates.
type TemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.TemplateLister
}

type templateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTemplateInformer constructs a new informer for Template type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTemplateInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTemplateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTemplateInformer constructs a new informer for Template type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTemplateInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Template().Templates(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Template().Templates(namespace).Watch(options)
			},
		},
		&template.Template{},
		resyncPeriod,
		indexers,
	)
}

func (f *templateInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTemplateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *templateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&template.Template{}, f.defaultInformer)
}

func (f *templateInformer) Lister() internalversion.TemplateLister {
	return internalversion.NewTemplateLister(f.Informer().GetIndexer())
}