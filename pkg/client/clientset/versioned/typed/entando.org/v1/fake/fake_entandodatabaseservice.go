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

// FakeEntandoDatabaseServices implements EntandoDatabaseServiceInterface
type FakeEntandoDatabaseServices struct {
	Fake *FakeEntandoV1
	ns   string
}

var entandodatabaseservicesResource = v1.SchemeGroupVersion.WithResource("entandodatabaseservices")

var entandodatabaseservicesKind = v1.SchemeGroupVersion.WithKind("EntandoDatabaseService")

// Get takes name of the entandoDatabaseService, and returns the corresponding entandoDatabaseService object, and an error if there is any.
func (c *FakeEntandoDatabaseServices) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.EntandoDatabaseService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(entandodatabaseservicesResource, c.ns, name), &v1.EntandoDatabaseService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EntandoDatabaseService), err
}

// List takes label and field selectors, and returns the list of EntandoDatabaseServices that match those selectors.
func (c *FakeEntandoDatabaseServices) List(ctx context.Context, opts metav1.ListOptions) (result *v1.EntandoDatabaseServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(entandodatabaseservicesResource, entandodatabaseservicesKind, c.ns, opts), &v1.EntandoDatabaseServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.EntandoDatabaseServiceList{ListMeta: obj.(*v1.EntandoDatabaseServiceList).ListMeta}
	for _, item := range obj.(*v1.EntandoDatabaseServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested entandoDatabaseServices.
func (c *FakeEntandoDatabaseServices) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(entandodatabaseservicesResource, c.ns, opts))

}

// Create takes the representation of a entandoDatabaseService and creates it.  Returns the server's representation of the entandoDatabaseService, and an error, if there is any.
func (c *FakeEntandoDatabaseServices) Create(ctx context.Context, entandoDatabaseService *v1.EntandoDatabaseService, opts metav1.CreateOptions) (result *v1.EntandoDatabaseService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(entandodatabaseservicesResource, c.ns, entandoDatabaseService), &v1.EntandoDatabaseService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EntandoDatabaseService), err
}

// Update takes the representation of a entandoDatabaseService and updates it. Returns the server's representation of the entandoDatabaseService, and an error, if there is any.
func (c *FakeEntandoDatabaseServices) Update(ctx context.Context, entandoDatabaseService *v1.EntandoDatabaseService, opts metav1.UpdateOptions) (result *v1.EntandoDatabaseService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(entandodatabaseservicesResource, c.ns, entandoDatabaseService), &v1.EntandoDatabaseService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EntandoDatabaseService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEntandoDatabaseServices) UpdateStatus(ctx context.Context, entandoDatabaseService *v1.EntandoDatabaseService, opts metav1.UpdateOptions) (*v1.EntandoDatabaseService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(entandodatabaseservicesResource, "status", c.ns, entandoDatabaseService), &v1.EntandoDatabaseService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EntandoDatabaseService), err
}

// Delete takes name of the entandoDatabaseService and deletes it. Returns an error if one occurs.
func (c *FakeEntandoDatabaseServices) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(entandodatabaseservicesResource, c.ns, name, opts), &v1.EntandoDatabaseService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEntandoDatabaseServices) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(entandodatabaseservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.EntandoDatabaseServiceList{})
	return err
}

// Patch applies the patch and returns the patched entandoDatabaseService.
func (c *FakeEntandoDatabaseServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.EntandoDatabaseService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(entandodatabaseservicesResource, c.ns, name, pt, data, subresources...), &v1.EntandoDatabaseService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EntandoDatabaseService), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied entandoDatabaseService.
func (c *FakeEntandoDatabaseServices) Apply(ctx context.Context, entandoDatabaseService *entandoorgv1.EntandoDatabaseServiceApplyConfiguration, opts metav1.ApplyOptions) (result *v1.EntandoDatabaseService, err error) {
	if entandoDatabaseService == nil {
		return nil, fmt.Errorf("entandoDatabaseService provided to Apply must not be nil")
	}
	data, err := json.Marshal(entandoDatabaseService)
	if err != nil {
		return nil, err
	}
	name := entandoDatabaseService.Name
	if name == nil {
		return nil, fmt.Errorf("entandoDatabaseService.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(entandodatabaseservicesResource, c.ns, *name, types.ApplyPatchType, data), &v1.EntandoDatabaseService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EntandoDatabaseService), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeEntandoDatabaseServices) ApplyStatus(ctx context.Context, entandoDatabaseService *entandoorgv1.EntandoDatabaseServiceApplyConfiguration, opts metav1.ApplyOptions) (result *v1.EntandoDatabaseService, err error) {
	if entandoDatabaseService == nil {
		return nil, fmt.Errorf("entandoDatabaseService provided to Apply must not be nil")
	}
	data, err := json.Marshal(entandoDatabaseService)
	if err != nil {
		return nil, err
	}
	name := entandoDatabaseService.Name
	if name == nil {
		return nil, fmt.Errorf("entandoDatabaseService.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(entandodatabaseservicesResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1.EntandoDatabaseService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EntandoDatabaseService), err
}
