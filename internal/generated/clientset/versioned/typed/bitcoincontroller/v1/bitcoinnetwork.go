// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "io.saagie/bitcoin-controller/internal/apis/bitcoincontroller/v1"
	scheme "io.saagie/bitcoin-controller/internal/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BitcoinNetworksGetter has a method to return a BitcoinNetworkInterface.
// A group's client should implement this interface.
type BitcoinNetworksGetter interface {
	BitcoinNetworks(namespace string) BitcoinNetworkInterface
}

// BitcoinNetworkInterface has methods to work with BitcoinNetwork resources.
type BitcoinNetworkInterface interface {
	Create(ctx context.Context, bitcoinNetwork *v1.BitcoinNetwork, opts metav1.CreateOptions) (*v1.BitcoinNetwork, error)
	Update(ctx context.Context, bitcoinNetwork *v1.BitcoinNetwork, opts metav1.UpdateOptions) (*v1.BitcoinNetwork, error)
	UpdateStatus(ctx context.Context, bitcoinNetwork *v1.BitcoinNetwork, opts metav1.UpdateOptions) (*v1.BitcoinNetwork, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.BitcoinNetwork, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.BitcoinNetworkList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.BitcoinNetwork, err error)
	BitcoinNetworkExpansion
}

// bitcoinNetworks implements BitcoinNetworkInterface
type bitcoinNetworks struct {
	client rest.Interface
	ns     string
}

// newBitcoinNetworks returns a BitcoinNetworks
func newBitcoinNetworks(c *BitcoincontrollerV1Client, namespace string) *bitcoinNetworks {
	return &bitcoinNetworks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bitcoinNetwork, and returns the corresponding bitcoinNetwork object, and an error if there is any.
func (c *bitcoinNetworks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.BitcoinNetwork, err error) {
	result = &v1.BitcoinNetwork{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bitcoinnetworks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BitcoinNetworks that match those selectors.
func (c *bitcoinNetworks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.BitcoinNetworkList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.BitcoinNetworkList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bitcoinnetworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bitcoinNetworks.
func (c *bitcoinNetworks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bitcoinnetworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bitcoinNetwork and creates it.  Returns the server's representation of the bitcoinNetwork, and an error, if there is any.
func (c *bitcoinNetworks) Create(ctx context.Context, bitcoinNetwork *v1.BitcoinNetwork, opts metav1.CreateOptions) (result *v1.BitcoinNetwork, err error) {
	result = &v1.BitcoinNetwork{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bitcoinnetworks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bitcoinNetwork).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bitcoinNetwork and updates it. Returns the server's representation of the bitcoinNetwork, and an error, if there is any.
func (c *bitcoinNetworks) Update(ctx context.Context, bitcoinNetwork *v1.BitcoinNetwork, opts metav1.UpdateOptions) (result *v1.BitcoinNetwork, err error) {
	result = &v1.BitcoinNetwork{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bitcoinnetworks").
		Name(bitcoinNetwork.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bitcoinNetwork).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *bitcoinNetworks) UpdateStatus(ctx context.Context, bitcoinNetwork *v1.BitcoinNetwork, opts metav1.UpdateOptions) (result *v1.BitcoinNetwork, err error) {
	result = &v1.BitcoinNetwork{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bitcoinnetworks").
		Name(bitcoinNetwork.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bitcoinNetwork).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bitcoinNetwork and deletes it. Returns an error if one occurs.
func (c *bitcoinNetworks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bitcoinnetworks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bitcoinNetworks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bitcoinnetworks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bitcoinNetwork.
func (c *bitcoinNetworks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.BitcoinNetwork, err error) {
	result = &v1.BitcoinNetwork{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bitcoinnetworks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
