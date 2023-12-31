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

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/antromeo/entando-clients/pkg/apis/entando.org/v1"
	entandoorgv1 "github.com/antromeo/entando-clients/pkg/client/applyconfiguration/entando.org/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEntandoDeBundles implements EntandoDeBundleInterface
type FakeEntandoDeBundles struct {
	Fake *FakeEntandoV1
	ns   string
}

var entandodebundlesResource = v1.SchemeGroupVersion.WithResource("entandodebundles")

var entandodebundlesKind = v1.SchemeGroupVersion.WithKind("EntandoDeBundle")

// Get takes name of the entandoDeBundle, and returns the corresponding entandoDeBundle object, and an error if there is any.
func (c *FakeEntandoDeBundles) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.EntandoDeBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(entandodebundlesResource, c.ns, name), &v1.EntandoDeBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EntandoDeBundle), err
}

// List takes label and field selectors, and returns the list of EntandoDeBundles that match those selectors.
func (c *FakeEntandoDeBundles) List(ctx context.Context, opts metav1.ListOptions) (result *v1.EntandoDeBundleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(entandodebundlesResource, entandodebundlesKind, c.ns, opts), &v1.EntandoDeBundleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.EntandoDeBundleList{ListMeta: obj.(*v1.EntandoDeBundleList).ListMeta}
	for _, item := range obj.(*v1.EntandoDeBundleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested entandoDeBundles.
func (c *FakeEntandoDeBundles) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(entandodebundlesResource, c.ns, opts))

}

// Create takes the representation of a entandoDeBundle and creates it.  Returns the server's representation of the entandoDeBundle, and an error, if there is any.
func (c *FakeEntandoDeBundles) Create(ctx context.Context, entandoDeBundle *v1.EntandoDeBundle, opts metav1.CreateOptions) (result *v1.EntandoDeBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(entandodebundlesResource, c.ns, entandoDeBundle), &v1.EntandoDeBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EntandoDeBundle), err
}

// Update takes the representation of a entandoDeBundle and updates it. Returns the server's representation of the entandoDeBundle, and an error, if there is any.
func (c *FakeEntandoDeBundles) Update(ctx context.Context, entandoDeBundle *v1.EntandoDeBundle, opts metav1.UpdateOptions) (result *v1.EntandoDeBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(entandodebundlesResource, c.ns, entandoDeBundle), &v1.EntandoDeBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EntandoDeBundle), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEntandoDeBundles) UpdateStatus(ctx context.Context, entandoDeBundle *v1.EntandoDeBundle, opts metav1.UpdateOptions) (*v1.EntandoDeBundle, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(entandodebundlesResource, "status", c.ns, entandoDeBundle), &v1.EntandoDeBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EntandoDeBundle), err
}

// Delete takes name of the entandoDeBundle and deletes it. Returns an error if one occurs.
func (c *FakeEntandoDeBundles) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(entandodebundlesResource, c.ns, name, opts), &v1.EntandoDeBundle{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEntandoDeBundles) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(entandodebundlesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.EntandoDeBundleList{})
	return err
}

// Patch applies the patch and returns the patched entandoDeBundle.
func (c *FakeEntandoDeBundles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.EntandoDeBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(entandodebundlesResource, c.ns, name, pt, data, subresources...), &v1.EntandoDeBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EntandoDeBundle), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied entandoDeBundle.
func (c *FakeEntandoDeBundles) Apply(ctx context.Context, entandoDeBundle *entandoorgv1.EntandoDeBundleApplyConfiguration, opts metav1.ApplyOptions) (result *v1.EntandoDeBundle, err error) {
	if entandoDeBundle == nil {
		return nil, fmt.Errorf("entandoDeBundle provided to Apply must not be nil")
	}
	data, err := json.Marshal(entandoDeBundle)
	if err != nil {
		return nil, err
	}
	name := entandoDeBundle.Name
	if name == nil {
		return nil, fmt.Errorf("entandoDeBundle.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(entandodebundlesResource, c.ns, *name, types.ApplyPatchType, data), &v1.EntandoDeBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EntandoDeBundle), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeEntandoDeBundles) ApplyStatus(ctx context.Context, entandoDeBundle *entandoorgv1.EntandoDeBundleApplyConfiguration, opts metav1.ApplyOptions) (result *v1.EntandoDeBundle, err error) {
	if entandoDeBundle == nil {
		return nil, fmt.Errorf("entandoDeBundle provided to Apply must not be nil")
	}
	data, err := json.Marshal(entandoDeBundle)
	if err != nil {
		return nil, err
	}
	name := entandoDeBundle.Name
	if name == nil {
		return nil, fmt.Errorf("entandoDeBundle.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(entandodebundlesResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1.EntandoDeBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EntandoDeBundle), err
}
