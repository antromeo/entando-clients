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

// EntandoDeBundleLister helps list EntandoDeBundles.
// All objects returned here must be treated as read-only.
type EntandoDeBundleLister interface {
	// List lists all EntandoDeBundles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EntandoDeBundle, err error)
	// EntandoDeBundles returns an object that can list and get EntandoDeBundles.
	EntandoDeBundles(namespace string) EntandoDeBundleNamespaceLister
	EntandoDeBundleListerExpansion
}

// entandoDeBundleLister implements the EntandoDeBundleLister interface.
type entandoDeBundleLister struct {
	indexer cache.Indexer
}

// NewEntandoDeBundleLister returns a new EntandoDeBundleLister.
func NewEntandoDeBundleLister(indexer cache.Indexer) EntandoDeBundleLister {
	return &entandoDeBundleLister{indexer: indexer}
}

// List lists all EntandoDeBundles in the indexer.
func (s *entandoDeBundleLister) List(selector labels.Selector) (ret []*v1.EntandoDeBundle, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EntandoDeBundle))
	})
	return ret, err
}

// EntandoDeBundles returns an object that can list and get EntandoDeBundles.
func (s *entandoDeBundleLister) EntandoDeBundles(namespace string) EntandoDeBundleNamespaceLister {
	return entandoDeBundleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EntandoDeBundleNamespaceLister helps list and get EntandoDeBundles.
// All objects returned here must be treated as read-only.
type EntandoDeBundleNamespaceLister interface {
	// List lists all EntandoDeBundles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EntandoDeBundle, err error)
	// Get retrieves the EntandoDeBundle from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.EntandoDeBundle, error)
	EntandoDeBundleNamespaceListerExpansion
}

// entandoDeBundleNamespaceLister implements the EntandoDeBundleNamespaceLister
// interface.
type entandoDeBundleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EntandoDeBundles in the indexer for a given namespace.
func (s entandoDeBundleNamespaceLister) List(selector labels.Selector) (ret []*v1.EntandoDeBundle, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EntandoDeBundle))
	})
	return ret, err
}

// Get retrieves the EntandoDeBundle from the indexer for a given namespace and name.
func (s entandoDeBundleNamespaceLister) Get(name string) (*v1.EntandoDeBundle, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("entandodebundle"), name)
	}
	return obj.(*v1.EntandoDeBundle), nil
}
