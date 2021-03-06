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

package v1alpha1

import (
	v1alpha1 "github.com/awslabs/aws-eks-cluster-controller/pkg/apis/cluster/v1alpha1"
	scheme "github.com/awslabs/aws-eks-cluster-controller/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ControlPlanesGetter has a method to return a ControlPlaneInterface.
// A group's client should implement this interface.
type ControlPlanesGetter interface {
	ControlPlanes(namespace string) ControlPlaneInterface
}

// ControlPlaneInterface has methods to work with ControlPlane resources.
type ControlPlaneInterface interface {
	Create(*v1alpha1.ControlPlane) (*v1alpha1.ControlPlane, error)
	Update(*v1alpha1.ControlPlane) (*v1alpha1.ControlPlane, error)
	UpdateStatus(*v1alpha1.ControlPlane) (*v1alpha1.ControlPlane, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ControlPlane, error)
	List(opts v1.ListOptions) (*v1alpha1.ControlPlaneList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ControlPlane, err error)
	ControlPlaneExpansion
}

// controlPlanes implements ControlPlaneInterface
type controlPlanes struct {
	client rest.Interface
	ns     string
}

// newControlPlanes returns a ControlPlanes
func newControlPlanes(c *ClusterV1alpha1Client, namespace string) *controlPlanes {
	return &controlPlanes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the controlPlane, and returns the corresponding controlPlane object, and an error if there is any.
func (c *controlPlanes) Get(name string, options v1.GetOptions) (result *v1alpha1.ControlPlane, err error) {
	result = &v1alpha1.ControlPlane{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("controlplanes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ControlPlanes that match those selectors.
func (c *controlPlanes) List(opts v1.ListOptions) (result *v1alpha1.ControlPlaneList, err error) {
	result = &v1alpha1.ControlPlaneList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("controlplanes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested controlPlanes.
func (c *controlPlanes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("controlplanes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a controlPlane and creates it.  Returns the server's representation of the controlPlane, and an error, if there is any.
func (c *controlPlanes) Create(controlPlane *v1alpha1.ControlPlane) (result *v1alpha1.ControlPlane, err error) {
	result = &v1alpha1.ControlPlane{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("controlplanes").
		Body(controlPlane).
		Do().
		Into(result)
	return
}

// Update takes the representation of a controlPlane and updates it. Returns the server's representation of the controlPlane, and an error, if there is any.
func (c *controlPlanes) Update(controlPlane *v1alpha1.ControlPlane) (result *v1alpha1.ControlPlane, err error) {
	result = &v1alpha1.ControlPlane{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("controlplanes").
		Name(controlPlane.Name).
		Body(controlPlane).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *controlPlanes) UpdateStatus(controlPlane *v1alpha1.ControlPlane) (result *v1alpha1.ControlPlane, err error) {
	result = &v1alpha1.ControlPlane{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("controlplanes").
		Name(controlPlane.Name).
		SubResource("status").
		Body(controlPlane).
		Do().
		Into(result)
	return
}

// Delete takes name of the controlPlane and deletes it. Returns an error if one occurs.
func (c *controlPlanes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("controlplanes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *controlPlanes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("controlplanes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched controlPlane.
func (c *controlPlanes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ControlPlane, err error) {
	result = &v1alpha1.ControlPlane{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("controlplanes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
