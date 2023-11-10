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
	internalinterfaces "github.com/antromeo/entando-clients/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// EntandoApps returns a EntandoAppInformer.
	EntandoApps() EntandoAppInformer
	// EntandoAppPluginLinks returns a EntandoAppPluginLinkInformer.
	EntandoAppPluginLinks() EntandoAppPluginLinkInformer
	// EntandoDatabaseServices returns a EntandoDatabaseServiceInformer.
	EntandoDatabaseServices() EntandoDatabaseServiceInformer
	// EntandoDeBundles returns a EntandoDeBundleInformer.
	EntandoDeBundles() EntandoDeBundleInformer
	// EntandoKeycloakServers returns a EntandoKeycloakServerInformer.
	EntandoKeycloakServers() EntandoKeycloakServerInformer
	// EntandoPlugins returns a EntandoPluginInformer.
	EntandoPlugins() EntandoPluginInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// EntandoApps returns a EntandoAppInformer.
func (v *version) EntandoApps() EntandoAppInformer {
	return &entandoAppInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EntandoAppPluginLinks returns a EntandoAppPluginLinkInformer.
func (v *version) EntandoAppPluginLinks() EntandoAppPluginLinkInformer {
	return &entandoAppPluginLinkInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EntandoDatabaseServices returns a EntandoDatabaseServiceInformer.
func (v *version) EntandoDatabaseServices() EntandoDatabaseServiceInformer {
	return &entandoDatabaseServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EntandoDeBundles returns a EntandoDeBundleInformer.
func (v *version) EntandoDeBundles() EntandoDeBundleInformer {
	return &entandoDeBundleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EntandoKeycloakServers returns a EntandoKeycloakServerInformer.
func (v *version) EntandoKeycloakServers() EntandoKeycloakServerInformer {
	return &entandoKeycloakServerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EntandoPlugins returns a EntandoPluginInformer.
func (v *version) EntandoPlugins() EntandoPluginInformer {
	return &entandoPluginInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
