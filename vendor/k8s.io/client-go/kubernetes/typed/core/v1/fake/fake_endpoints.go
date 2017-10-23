/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
<<<<<<< HEAD
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
=======
	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
>>>>>>> upstream/master
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
<<<<<<< HEAD
	v1 "k8s.io/client-go/pkg/api/v1"
=======
>>>>>>> upstream/master
	testing "k8s.io/client-go/testing"
)

// FakeEndpoints implements EndpointsInterface
type FakeEndpoints struct {
	Fake *FakeCoreV1
	ns   string
}

var endpointsResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "endpoints"}

var endpointsKind = schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Endpoints"}

<<<<<<< HEAD
func (c *FakeEndpoints) Create(endpoints *v1.Endpoints) (result *v1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(endpointsResource, c.ns, endpoints), &v1.Endpoints{})
=======
// Get takes name of the endpoints, and returns the corresponding endpoints object, and an error if there is any.
func (c *FakeEndpoints) Get(name string, options v1.GetOptions) (result *core_v1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(endpointsResource, c.ns, name), &core_v1.Endpoints{})
>>>>>>> upstream/master

	if obj == nil {
		return nil, err
	}
<<<<<<< HEAD
	return obj.(*v1.Endpoints), err
}

func (c *FakeEndpoints) Update(endpoints *v1.Endpoints) (result *v1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(endpointsResource, c.ns, endpoints), &v1.Endpoints{})
=======
	return obj.(*core_v1.Endpoints), err
}

// List takes label and field selectors, and returns the list of Endpoints that match those selectors.
func (c *FakeEndpoints) List(opts v1.ListOptions) (result *core_v1.EndpointsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(endpointsResource, endpointsKind, c.ns, opts), &core_v1.EndpointsList{})
>>>>>>> upstream/master

	if obj == nil {
		return nil, err
	}
<<<<<<< HEAD
	return obj.(*v1.Endpoints), err
}

func (c *FakeEndpoints) Delete(name string, options *meta_v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(endpointsResource, c.ns, name), &v1.Endpoints{})

	return err
}

func (c *FakeEndpoints) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(endpointsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1.EndpointsList{})
	return err
}

func (c *FakeEndpoints) Get(name string, options meta_v1.GetOptions) (result *v1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(endpointsResource, c.ns, name), &v1.Endpoints{})
=======

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &core_v1.EndpointsList{}
	for _, item := range obj.(*core_v1.EndpointsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested endpoints.
func (c *FakeEndpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(endpointsResource, c.ns, opts))

}

// Create takes the representation of a endpoints and creates it.  Returns the server's representation of the endpoints, and an error, if there is any.
func (c *FakeEndpoints) Create(endpoints *core_v1.Endpoints) (result *core_v1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(endpointsResource, c.ns, endpoints), &core_v1.Endpoints{})
>>>>>>> upstream/master

	if obj == nil {
		return nil, err
	}
<<<<<<< HEAD
	return obj.(*v1.Endpoints), err
}

func (c *FakeEndpoints) List(opts meta_v1.ListOptions) (result *v1.EndpointsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(endpointsResource, endpointsKind, c.ns, opts), &v1.EndpointsList{})
=======
	return obj.(*core_v1.Endpoints), err
}

// Update takes the representation of a endpoints and updates it. Returns the server's representation of the endpoints, and an error, if there is any.
func (c *FakeEndpoints) Update(endpoints *core_v1.Endpoints) (result *core_v1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(endpointsResource, c.ns, endpoints), &core_v1.Endpoints{})
>>>>>>> upstream/master

	if obj == nil {
		return nil, err
	}
<<<<<<< HEAD

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.EndpointsList{}
	for _, item := range obj.(*v1.EndpointsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested endpoints.
func (c *FakeEndpoints) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(endpointsResource, c.ns, opts))

}

// Patch applies the patch and returns the patched endpoints.
func (c *FakeEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(endpointsResource, c.ns, name, data, subresources...), &v1.Endpoints{})
=======
	return obj.(*core_v1.Endpoints), err
}

// Delete takes name of the endpoints and deletes it. Returns an error if one occurs.
func (c *FakeEndpoints) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(endpointsResource, c.ns, name), &core_v1.Endpoints{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEndpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(endpointsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &core_v1.EndpointsList{})
	return err
}

// Patch applies the patch and returns the patched endpoints.
func (c *FakeEndpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *core_v1.Endpoints, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(endpointsResource, c.ns, name, data, subresources...), &core_v1.Endpoints{})
>>>>>>> upstream/master

	if obj == nil {
		return nil, err
	}
<<<<<<< HEAD
	return obj.(*v1.Endpoints), err
=======
	return obj.(*core_v1.Endpoints), err
>>>>>>> upstream/master
}
