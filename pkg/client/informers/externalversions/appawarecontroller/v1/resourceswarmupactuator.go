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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	appawarecontrollerv1 "k8s.io/application-aware-controller/pkg/apis/appawarecontroller/v1"
	versioned "k8s.io/application-aware-controller/pkg/client/clientset/versioned"
	internalinterfaces "k8s.io/application-aware-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1 "k8s.io/application-aware-controller/pkg/client/listers/appawarecontroller/v1"
	cache "k8s.io/client-go/tools/cache"
)

// ResourcesWarmupActuatorInformer provides access to a shared informer and lister for
// ResourcesWarmupActuators.
type ResourcesWarmupActuatorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ResourcesWarmupActuatorLister
}

type resourcesWarmupActuatorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewResourcesWarmupActuatorInformer constructs a new informer for ResourcesWarmupActuator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourcesWarmupActuatorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourcesWarmupActuatorInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredResourcesWarmupActuatorInformer constructs a new informer for ResourcesWarmupActuator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourcesWarmupActuatorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppawarecontrollerV1().ResourcesWarmupActuators(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppawarecontrollerV1().ResourcesWarmupActuators(namespace).Watch(context.TODO(), options)
			},
		},
		&appawarecontrollerv1.ResourcesWarmupActuator{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourcesWarmupActuatorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourcesWarmupActuatorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourcesWarmupActuatorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appawarecontrollerv1.ResourcesWarmupActuator{}, f.defaultInformer)
}

func (f *resourcesWarmupActuatorInformer) Lister() v1.ResourcesWarmupActuatorLister {
	return v1.NewResourcesWarmupActuatorLister(f.Informer().GetIndexer())
}