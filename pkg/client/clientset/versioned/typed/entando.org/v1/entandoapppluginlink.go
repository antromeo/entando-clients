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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/antromeo/entando-clients/pkg/apis/entando.org/v1"
	entandoorgv1 "github.com/antromeo/entando-clients/pkg/client/applyconfiguration/entando.org/v1"
	scheme "github.com/antromeo/entando-clients/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EntandoAppPluginLinksGetter has a method to return a EntandoAppPluginLinkInterface.
// A group's client should implement this interface.
type EntandoAppPluginLinksGetter interface {
	EntandoAppPluginLinks(namespace string) EntandoAppPluginLinkInterface
}

// EntandoAppPluginLinkInterface has methods to work with EntandoAppPluginLink resources.
type EntandoAppPluginLinkInterface interface {
	Create(ctx context.Context, entandoAppPluginLink *v1.EntandoAppPluginLink, opts metav1.CreateOptions) (*v1.EntandoAppPluginLink, error)
	Update(ctx context.Context, entandoAppPluginLink *v1.EntandoAppPluginLink, opts metav1.UpdateOptions) (*v1.EntandoAppPluginLink, error)
	UpdateStatus(ctx context.Context, entandoAppPluginLink *v1.EntandoAppPluginLink, opts metav1.UpdateOptions) (*v1.EntandoAppPluginLink, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.EntandoAppPluginLink, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.EntandoAppPluginLinkList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.EntandoAppPluginLink, err error)
	Apply(ctx context.Context, entandoAppPluginLink *entandoorgv1.EntandoAppPluginLinkApplyConfiguration, opts metav1.ApplyOptions) (result *v1.EntandoAppPluginLink, err error)
	ApplyStatus(ctx context.Context, entandoAppPluginLink *entandoorgv1.EntandoAppPluginLinkApplyConfiguration, opts metav1.ApplyOptions) (result *v1.EntandoAppPluginLink, err error)
	EntandoAppPluginLinkExpansion
}

// entandoAppPluginLinks implements EntandoAppPluginLinkInterface
type entandoAppPluginLinks struct {
	client rest.Interface
	ns     string
}

// newEntandoAppPluginLinks returns a EntandoAppPluginLinks
func newEntandoAppPluginLinks(c *EntandoV1Client, namespace string) *entandoAppPluginLinks {
	return &entandoAppPluginLinks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the entandoAppPluginLink, and returns the corresponding entandoAppPluginLink object, and an error if there is any.
func (c *entandoAppPluginLinks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.EntandoAppPluginLink, err error) {
	result = &v1.EntandoAppPluginLink{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("entandoapppluginlinks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EntandoAppPluginLinks that match those selectors.
func (c *entandoAppPluginLinks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.EntandoAppPluginLinkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.EntandoAppPluginLinkList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("entandoapppluginlinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested entandoAppPluginLinks.
func (c *entandoAppPluginLinks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("entandoapppluginlinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a entandoAppPluginLink and creates it.  Returns the server's representation of the entandoAppPluginLink, and an error, if there is any.
func (c *entandoAppPluginLinks) Create(ctx context.Context, entandoAppPluginLink *v1.EntandoAppPluginLink, opts metav1.CreateOptions) (result *v1.EntandoAppPluginLink, err error) {
	result = &v1.EntandoAppPluginLink{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("entandoapppluginlinks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(entandoAppPluginLink).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a entandoAppPluginLink and updates it. Returns the server's representation of the entandoAppPluginLink, and an error, if there is any.
func (c *entandoAppPluginLinks) Update(ctx context.Context, entandoAppPluginLink *v1.EntandoAppPluginLink, opts metav1.UpdateOptions) (result *v1.EntandoAppPluginLink, err error) {
	result = &v1.EntandoAppPluginLink{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("entandoapppluginlinks").
		Name(entandoAppPluginLink.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(entandoAppPluginLink).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *entandoAppPluginLinks) UpdateStatus(ctx context.Context, entandoAppPluginLink *v1.EntandoAppPluginLink, opts metav1.UpdateOptions) (result *v1.EntandoAppPluginLink, err error) {
	result = &v1.EntandoAppPluginLink{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("entandoapppluginlinks").
		Name(entandoAppPluginLink.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(entandoAppPluginLink).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the entandoAppPluginLink and deletes it. Returns an error if one occurs.
func (c *entandoAppPluginLinks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("entandoapppluginlinks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *entandoAppPluginLinks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("entandoapppluginlinks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched entandoAppPluginLink.
func (c *entandoAppPluginLinks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.EntandoAppPluginLink, err error) {
	result = &v1.EntandoAppPluginLink{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("entandoapppluginlinks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied entandoAppPluginLink.
func (c *entandoAppPluginLinks) Apply(ctx context.Context, entandoAppPluginLink *entandoorgv1.EntandoAppPluginLinkApplyConfiguration, opts metav1.ApplyOptions) (result *v1.EntandoAppPluginLink, err error) {
	if entandoAppPluginLink == nil {
		return nil, fmt.Errorf("entandoAppPluginLink provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(entandoAppPluginLink)
	if err != nil {
		return nil, err
	}
	name := entandoAppPluginLink.Name
	if name == nil {
		return nil, fmt.Errorf("entandoAppPluginLink.Name must be provided to Apply")
	}
	result = &v1.EntandoAppPluginLink{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("entandoapppluginlinks").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *entandoAppPluginLinks) ApplyStatus(ctx context.Context, entandoAppPluginLink *entandoorgv1.EntandoAppPluginLinkApplyConfiguration, opts metav1.ApplyOptions) (result *v1.EntandoAppPluginLink, err error) {
	if entandoAppPluginLink == nil {
		return nil, fmt.Errorf("entandoAppPluginLink provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(entandoAppPluginLink)
	if err != nil {
		return nil, err
	}

	name := entandoAppPluginLink.Name
	if name == nil {
		return nil, fmt.Errorf("entandoAppPluginLink.Name must be provided to Apply")
	}

	result = &v1.EntandoAppPluginLink{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("entandoapppluginlinks").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
