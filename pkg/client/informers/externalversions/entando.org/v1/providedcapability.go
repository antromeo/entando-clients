/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"
	time "time"

	entandoorgv1 "github.com/antromeo/entando-clients/pkg/apis/entando.org/v1"
	versioned "github.com/antromeo/entando-clients/pkg/client/clientset/versioned"
	internalinterfaces "github.com/antromeo/entando-clients/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/antromeo/entando-clients/pkg/client/listers/entando.org/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ProvidedCapabilityInformer provides access to a shared informer and lister for
// ProvidedCapabilities.
type ProvidedCapabilityInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ProvidedCapabilityLister
}

type providedCapabilityInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewProvidedCapabilityInformer constructs a new informer for ProvidedCapability type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewProvidedCapabilityInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredProvidedCapabilityInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredProvidedCapabilityInformer constructs a new informer for ProvidedCapability type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredProvidedCapabilityInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EntandoV1().ProvidedCapabilities(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EntandoV1().ProvidedCapabilities(namespace).Watch(context.TODO(), options)
			},
		},
		&entandoorgv1.ProvidedCapability{},
		resyncPeriod,
		indexers,
	)
}

func (f *providedCapabilityInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredProvidedCapabilityInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *providedCapabilityInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&entandoorgv1.ProvidedCapability{}, f.defaultInformer)
}

func (f *providedCapabilityInformer) Lister() v1.ProvidedCapabilityLister {
	return v1.NewProvidedCapabilityLister(f.Informer().GetIndexer())
}
