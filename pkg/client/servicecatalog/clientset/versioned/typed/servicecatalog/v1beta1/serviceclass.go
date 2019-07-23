// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	scheme "github.com/google/kf/pkg/client/servicecatalog/clientset/versioned/scheme"
	v1beta1 "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServiceClassesGetter has a method to return a ServiceClassInterface.
// A group's client should implement this interface.
type ServiceClassesGetter interface {
	ServiceClasses(namespace string) ServiceClassInterface
}

// ServiceClassInterface has methods to work with ServiceClass resources.
type ServiceClassInterface interface {
	Create(*v1beta1.ServiceClass) (*v1beta1.ServiceClass, error)
	Update(*v1beta1.ServiceClass) (*v1beta1.ServiceClass, error)
	UpdateStatus(*v1beta1.ServiceClass) (*v1beta1.ServiceClass, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ServiceClass, error)
	List(opts v1.ListOptions) (*v1beta1.ServiceClassList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServiceClass, err error)
	ServiceClassExpansion
}

// serviceClasses implements ServiceClassInterface
type serviceClasses struct {
	client rest.Interface
	ns     string
}

// newServiceClasses returns a ServiceClasses
func newServiceClasses(c *ServicecatalogV1beta1Client, namespace string) *serviceClasses {
	return &serviceClasses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceClass, and returns the corresponding serviceClass object, and an error if there is any.
func (c *serviceClasses) Get(name string, options v1.GetOptions) (result *v1beta1.ServiceClass, err error) {
	result = &v1beta1.ServiceClass{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceClasses that match those selectors.
func (c *serviceClasses) List(opts v1.ListOptions) (result *v1beta1.ServiceClassList, err error) {
	result = &v1beta1.ServiceClassList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceClasses.
func (c *serviceClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("serviceclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a serviceClass and creates it.  Returns the server's representation of the serviceClass, and an error, if there is any.
func (c *serviceClasses) Create(serviceClass *v1beta1.ServiceClass) (result *v1beta1.ServiceClass, err error) {
	result = &v1beta1.ServiceClass{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("serviceclasses").
		Body(serviceClass).
		Do().
		Into(result)
	return
}

// Update takes the representation of a serviceClass and updates it. Returns the server's representation of the serviceClass, and an error, if there is any.
func (c *serviceClasses) Update(serviceClass *v1beta1.ServiceClass) (result *v1beta1.ServiceClass, err error) {
	result = &v1beta1.ServiceClass{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceclasses").
		Name(serviceClass.Name).
		Body(serviceClass).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *serviceClasses) UpdateStatus(serviceClass *v1beta1.ServiceClass) (result *v1beta1.ServiceClass, err error) {
	result = &v1beta1.ServiceClass{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceclasses").
		Name(serviceClass.Name).
		SubResource("status").
		Body(serviceClass).
		Do().
		Into(result)
	return
}

// Delete takes name of the serviceClass and deletes it. Returns an error if one occurs.
func (c *serviceClasses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceclasses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceclasses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched serviceClass.
func (c *serviceClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ServiceClass, err error) {
	result = &v1beta1.ServiceClass{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("serviceclasses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
