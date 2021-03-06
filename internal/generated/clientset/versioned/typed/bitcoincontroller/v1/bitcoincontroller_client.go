// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "io.saagie/bitcoin-controller/internal/apis/bitcoincontroller/v1"
	"io.saagie/bitcoin-controller/internal/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type BitcoincontrollerV1Interface interface {
	RESTClient() rest.Interface
	BitcoinNetworksGetter
}

// BitcoincontrollerV1Client is used to interact with features provided by the bitcoincontroller group.
type BitcoincontrollerV1Client struct {
	restClient rest.Interface
}

func (c *BitcoincontrollerV1Client) BitcoinNetworks(namespace string) BitcoinNetworkInterface {
	return newBitcoinNetworks(c, namespace)
}

// NewForConfig creates a new BitcoincontrollerV1Client for the given config.
func NewForConfig(c *rest.Config) (*BitcoincontrollerV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &BitcoincontrollerV1Client{client}, nil
}

// NewForConfigOrDie creates a new BitcoincontrollerV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *BitcoincontrollerV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new BitcoincontrollerV1Client for the given RESTClient.
func New(c rest.Interface) *BitcoincontrollerV1Client {
	return &BitcoincontrollerV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *BitcoincontrollerV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
