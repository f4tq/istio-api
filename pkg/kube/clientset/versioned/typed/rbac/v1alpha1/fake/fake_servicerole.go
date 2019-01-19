// Copyright 2019 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "istio.io/api/pkg/kube/apis/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceRoles implements ServiceRoleInterface
type FakeServiceRoles struct {
	Fake *FakeRbacV1alpha1
	ns   string
}

var servicerolesResource = schema.GroupVersionResource{Group: "rbac", Version: "v1alpha1", Resource: "serviceroles"}

var servicerolesKind = schema.GroupVersionKind{Group: "rbac", Version: "v1alpha1", Kind: "ServiceRole"}

// Get takes name of the serviceRole, and returns the corresponding serviceRole object, and an error if there is any.
func (c *FakeServiceRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.ServiceRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicerolesResource, c.ns, name), &v1alpha1.ServiceRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceRole), err
}

// List takes label and field selectors, and returns the list of ServiceRoles that match those selectors.
func (c *FakeServiceRoles) List(opts v1.ListOptions) (result *v1alpha1.ServiceRoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicerolesResource, servicerolesKind, c.ns, opts), &v1alpha1.ServiceRoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceRoleList{ListMeta: obj.(*v1alpha1.ServiceRoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceRoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceRoles.
func (c *FakeServiceRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicerolesResource, c.ns, opts))

}

// Create takes the representation of a serviceRole and creates it.  Returns the server's representation of the serviceRole, and an error, if there is any.
func (c *FakeServiceRoles) Create(serviceRole *v1alpha1.ServiceRole) (result *v1alpha1.ServiceRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicerolesResource, c.ns, serviceRole), &v1alpha1.ServiceRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceRole), err
}

// Update takes the representation of a serviceRole and updates it. Returns the server's representation of the serviceRole, and an error, if there is any.
func (c *FakeServiceRoles) Update(serviceRole *v1alpha1.ServiceRole) (result *v1alpha1.ServiceRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicerolesResource, c.ns, serviceRole), &v1alpha1.ServiceRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceRole), err
}

// Delete takes name of the serviceRole and deletes it. Returns an error if one occurs.
func (c *FakeServiceRoles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicerolesResource, c.ns, name), &v1alpha1.ServiceRole{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicerolesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceRoleList{})
	return err
}

// Patch applies the patch and returns the patched serviceRole.
func (c *FakeServiceRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceRole, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicerolesResource, c.ns, name, data, subresources...), &v1alpha1.ServiceRole{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceRole), err
}
