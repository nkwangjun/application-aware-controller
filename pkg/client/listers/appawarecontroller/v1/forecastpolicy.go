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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	v1 "k8s.io/application-aware-controller/pkg/apis/appawarecontroller/v1"
	"k8s.io/client-go/tools/cache"
)

// ForecastPolicyLister helps list ForecastPolicies.
// All objects returned here must be treated as read-only.
type ForecastPolicyLister interface {
	// List lists all ForecastPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ForecastPolicy, err error)
	// ForecastPolicies returns an object that can list and get ForecastPolicies.
	ForecastPolicies(namespace string) ForecastPolicyNamespaceLister
	ForecastPolicyListerExpansion
}

// forecastPolicyLister implements the ForecastPolicyLister interface.
type forecastPolicyLister struct {
	indexer cache.Indexer
}

// NewForecastPolicyLister returns a new ForecastPolicyLister.
func NewForecastPolicyLister(indexer cache.Indexer) ForecastPolicyLister {
	return &forecastPolicyLister{indexer: indexer}
}

// List lists all ForecastPolicies in the indexer.
func (s *forecastPolicyLister) List(selector labels.Selector) (ret []*v1.ForecastPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ForecastPolicy))
	})
	return ret, err
}

// ForecastPolicies returns an object that can list and get ForecastPolicies.
func (s *forecastPolicyLister) ForecastPolicies(namespace string) ForecastPolicyNamespaceLister {
	return forecastPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ForecastPolicyNamespaceLister helps list and get ForecastPolicies.
// All objects returned here must be treated as read-only.
type ForecastPolicyNamespaceLister interface {
	// List lists all ForecastPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ForecastPolicy, err error)
	// Get retrieves the ForecastPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ForecastPolicy, error)
	ForecastPolicyNamespaceListerExpansion
}

// forecastPolicyNamespaceLister implements the ForecastPolicyNamespaceLister
// interface.
type forecastPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ForecastPolicies in the indexer for a given namespace.
func (s forecastPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1.ForecastPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ForecastPolicy))
	})
	return ret, err
}

// Get retrieves the ForecastPolicy from the indexer for a given namespace and name.
func (s forecastPolicyNamespaceLister) Get(name string) (*v1.ForecastPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("forecastpolicy"), name)
	}
	return obj.(*v1.ForecastPolicy), nil
}