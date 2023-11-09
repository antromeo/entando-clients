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

// EntandoPluginLister helps list EntandoPlugins.
// All objects returned here must be treated as read-only.
type EntandoPluginLister interface {
	// List lists all EntandoPlugins in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EntandoPlugin, err error)
	// EntandoPlugins returns an object that can list and get EntandoPlugins.
	EntandoPlugins(namespace string) EntandoPluginNamespaceLister
	EntandoPluginListerExpansion
}

// entandoPluginLister implements the EntandoPluginLister interface.
type entandoPluginLister struct {
	indexer cache.Indexer
}

// NewEntandoPluginLister returns a new EntandoPluginLister.
func NewEntandoPluginLister(indexer cache.Indexer) EntandoPluginLister {
	return &entandoPluginLister{indexer: indexer}
}

// List lists all EntandoPlugins in the indexer.
func (s *entandoPluginLister) List(selector labels.Selector) (ret []*v1.EntandoPlugin, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EntandoPlugin))
	})
	return ret, err
}

// EntandoPlugins returns an object that can list and get EntandoPlugins.
func (s *entandoPluginLister) EntandoPlugins(namespace string) EntandoPluginNamespaceLister {
	return entandoPluginNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EntandoPluginNamespaceLister helps list and get EntandoPlugins.
// All objects returned here must be treated as read-only.
type EntandoPluginNamespaceLister interface {
	// List lists all EntandoPlugins in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EntandoPlugin, err error)
	// Get retrieves the EntandoPlugin from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.EntandoPlugin, error)
	EntandoPluginNamespaceListerExpansion
}

// entandoPluginNamespaceLister implements the EntandoPluginNamespaceLister
// interface.
type entandoPluginNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EntandoPlugins in the indexer for a given namespace.
func (s entandoPluginNamespaceLister) List(selector labels.Selector) (ret []*v1.EntandoPlugin, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EntandoPlugin))
	})
	return ret, err
}

// Get retrieves the EntandoPlugin from the indexer for a given namespace and name.
func (s entandoPluginNamespaceLister) Get(name string) (*v1.EntandoPlugin, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("entandoplugin"), name)
	}
	return obj.(*v1.EntandoPlugin), nil
}