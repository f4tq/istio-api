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

package v1alpha2

import (
	v1alpha2 "istio.io/api/pkg/kube/apis/config/v1alpha2"
	scheme "istio.io/api/pkg/kube/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HTTPAPISpecsGetter has a method to return a HTTPAPISpecInterface.
// A group's client should implement this interface.
type HTTPAPISpecsGetter interface {
	HTTPAPISpecs(namespace string) HTTPAPISpecInterface
}

// HTTPAPISpecInterface has methods to work with HTTPAPISpec resources.
type HTTPAPISpecInterface interface {
	Create(*v1alpha2.HTTPAPISpec) (*v1alpha2.HTTPAPISpec, error)
	Update(*v1alpha2.HTTPAPISpec) (*v1alpha2.HTTPAPISpec, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.HTTPAPISpec, error)
	List(opts v1.ListOptions) (*v1alpha2.HTTPAPISpecList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.HTTPAPISpec, err error)
	HTTPAPISpecExpansion
}

// hTTPAPISpecs implements HTTPAPISpecInterface
type hTTPAPISpecs struct {
	client rest.Interface
	ns     string
}

// newHTTPAPISpecs returns a HTTPAPISpecs
func newHTTPAPISpecs(c *ConfigV1alpha2Client, namespace string) *hTTPAPISpecs {
	return &hTTPAPISpecs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hTTPAPISpec, and returns the corresponding hTTPAPISpec object, and an error if there is any.
func (c *hTTPAPISpecs) Get(name string, options v1.GetOptions) (result *v1alpha2.HTTPAPISpec, err error) {
	result = &v1alpha2.HTTPAPISpec{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httpapispecs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HTTPAPISpecs that match those selectors.
func (c *hTTPAPISpecs) List(opts v1.ListOptions) (result *v1alpha2.HTTPAPISpecList, err error) {
	result = &v1alpha2.HTTPAPISpecList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httpapispecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hTTPAPISpecs.
func (c *hTTPAPISpecs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("httpapispecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a hTTPAPISpec and creates it.  Returns the server's representation of the hTTPAPISpec, and an error, if there is any.
func (c *hTTPAPISpecs) Create(hTTPAPISpec *v1alpha2.HTTPAPISpec) (result *v1alpha2.HTTPAPISpec, err error) {
	result = &v1alpha2.HTTPAPISpec{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("httpapispecs").
		Body(hTTPAPISpec).
		Do().
		Into(result)
	return
}

// Update takes the representation of a hTTPAPISpec and updates it. Returns the server's representation of the hTTPAPISpec, and an error, if there is any.
func (c *hTTPAPISpecs) Update(hTTPAPISpec *v1alpha2.HTTPAPISpec) (result *v1alpha2.HTTPAPISpec, err error) {
	result = &v1alpha2.HTTPAPISpec{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("httpapispecs").
		Name(hTTPAPISpec.Name).
		Body(hTTPAPISpec).
		Do().
		Into(result)
	return
}

// Delete takes name of the hTTPAPISpec and deletes it. Returns an error if one occurs.
func (c *hTTPAPISpecs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httpapispecs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hTTPAPISpecs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httpapispecs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched hTTPAPISpec.
func (c *hTTPAPISpecs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.HTTPAPISpec, err error) {
	result = &v1alpha2.HTTPAPISpec{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("httpapispecs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
