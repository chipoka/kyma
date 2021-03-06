// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kyma-project/kyma/components/remote-environment-broker/pkg/apis/remoteenvironment/v1alpha1"
	scheme "github.com/kyma-project/kyma/components/remote-environment-broker/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EnvironmentMappingsGetter has a method to return a EnvironmentMappingInterface.
// A group's client should implement this interface.
type EnvironmentMappingsGetter interface {
	EnvironmentMappings(namespace string) EnvironmentMappingInterface
}

// EnvironmentMappingInterface has methods to work with EnvironmentMapping resources.
type EnvironmentMappingInterface interface {
	Create(*v1alpha1.EnvironmentMapping) (*v1alpha1.EnvironmentMapping, error)
	Update(*v1alpha1.EnvironmentMapping) (*v1alpha1.EnvironmentMapping, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.EnvironmentMapping, error)
	List(opts v1.ListOptions) (*v1alpha1.EnvironmentMappingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EnvironmentMapping, err error)
	EnvironmentMappingExpansion
}

// environmentMappings implements EnvironmentMappingInterface
type environmentMappings struct {
	client rest.Interface
	ns     string
}

// newEnvironmentMappings returns a EnvironmentMappings
func newEnvironmentMappings(c *RemoteenvironmentV1alpha1Client, namespace string) *environmentMappings {
	return &environmentMappings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the environmentMapping, and returns the corresponding environmentMapping object, and an error if there is any.
func (c *environmentMappings) Get(name string, options v1.GetOptions) (result *v1alpha1.EnvironmentMapping, err error) {
	result = &v1alpha1.EnvironmentMapping{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("environmentmappings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EnvironmentMappings that match those selectors.
func (c *environmentMappings) List(opts v1.ListOptions) (result *v1alpha1.EnvironmentMappingList, err error) {
	result = &v1alpha1.EnvironmentMappingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("environmentmappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested environmentMappings.
func (c *environmentMappings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("environmentmappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a environmentMapping and creates it.  Returns the server's representation of the environmentMapping, and an error, if there is any.
func (c *environmentMappings) Create(environmentMapping *v1alpha1.EnvironmentMapping) (result *v1alpha1.EnvironmentMapping, err error) {
	result = &v1alpha1.EnvironmentMapping{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("environmentmappings").
		Body(environmentMapping).
		Do().
		Into(result)
	return
}

// Update takes the representation of a environmentMapping and updates it. Returns the server's representation of the environmentMapping, and an error, if there is any.
func (c *environmentMappings) Update(environmentMapping *v1alpha1.EnvironmentMapping) (result *v1alpha1.EnvironmentMapping, err error) {
	result = &v1alpha1.EnvironmentMapping{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("environmentmappings").
		Name(environmentMapping.Name).
		Body(environmentMapping).
		Do().
		Into(result)
	return
}

// Delete takes name of the environmentMapping and deletes it. Returns an error if one occurs.
func (c *environmentMappings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("environmentmappings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *environmentMappings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("environmentmappings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched environmentMapping.
func (c *environmentMappings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EnvironmentMapping, err error) {
	result = &v1alpha1.EnvironmentMapping{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("environmentmappings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
