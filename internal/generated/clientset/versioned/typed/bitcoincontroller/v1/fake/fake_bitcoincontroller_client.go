// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "io.saagie/bitcoin-controller/internal/generated/clientset/versioned/typed/bitcoincontroller/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeBitcoincontrollerV1 struct {
	*testing.Fake
}

func (c *FakeBitcoincontrollerV1) BitcoinNetworks(namespace string) v1.BitcoinNetworkInterface {
	return &FakeBitcoinNetworks{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeBitcoincontrollerV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
