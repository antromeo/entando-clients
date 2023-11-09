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
	v1 "github.com/antromeo/entando-clients/pkg/apis/entando.org/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EntandoAppLister helps list EntandoApps.
// All objects returned here must be treated as read-only.
type EntandoAppLister interface {
	// List lists all EntandoApps in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EntandoApp, err error)
	// EntandoApps returns an object that can list and get EntandoApps.
	EntandoApps(namespace string) EntandoAppNamespaceLister
	EntandoAppListerExpansion
}

// entandoAppLister implements the EntandoAppLister interface.
type entandoAppLister struct {
	indexer cache.Indexer
}

// NewEntandoAppLister returns a new EntandoAppLister.
func NewEntandoAppLister(indexer cache.Indexer) EntandoAppLister {
	return &entandoAppLister{indexer: indexer}
}

// List lists all EntandoApps in the indexer.
func (s *entandoAppLister) List(selector labels.Selector) (ret []*v1.EntandoApp, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EntandoApp))
	})
	return ret, err
}

// EntandoApps returns an object that can list and get EntandoApps.
func (s *entandoAppLister) EntandoApps(namespace string) EntandoAppNamespaceLister {
	return entandoAppNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EntandoAppNamespaceLister helps list and get EntandoApps.
// All objects returned here must be treated as read-only.
type EntandoAppNamespaceLister interface {
	// List lists all EntandoApps in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EntandoApp, err error)
	// Get retrieves the EntandoApp from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.EntandoApp, error)
	EntandoAppNamespaceListerExpansion
}

// entandoAppNamespaceLister implements the EntandoAppNamespaceLister
// interface.
type entandoAppNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EntandoApps in the indexer for a given namespace.
func (s entandoAppNamespaceLister) List(selector labels.Selector) (ret []*v1.EntandoApp, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EntandoApp))
	})
	return ret, err
}

// Get retrieves the EntandoApp from the indexer for a given namespace and name.
func (s entandoAppNamespaceLister) Get(name string) (*v1.EntandoApp, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("entandoapp"), name)
	}
	return obj.(*v1.EntandoApp), nil
}