// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	bitcoincontrollerv1 "io.saagie/bitcoin-controller/internal/apis/bitcoincontroller/v1"
	versioned "io.saagie/bitcoin-controller/internal/generated/clientset/versioned"
	internalinterfaces "io.saagie/bitcoin-controller/internal/generated/informers/externalversions/internalinterfaces"
	v1 "io.saagie/bitcoin-controller/internal/generated/listers/bitcoincontroller/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BitcoinNetworkInformer provides access to a shared informer and lister for
// BitcoinNetworks.
type BitcoinNetworkInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.BitcoinNetworkLister
}

type bitcoinNetworkInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBitcoinNetworkInformer constructs a new informer for BitcoinNetwork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBitcoinNetworkInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBitcoinNetworkInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBitcoinNetworkInformer constructs a new informer for BitcoinNetwork type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBitcoinNetworkInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BitcoincontrollerV1().BitcoinNetworks(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BitcoincontrollerV1().BitcoinNetworks(namespace).Watch(context.TODO(), options)
			},
		},
		&bitcoincontrollerv1.BitcoinNetwork{},
		resyncPeriod,
		indexers,
	)
}

func (f *bitcoinNetworkInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBitcoinNetworkInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *bitcoinNetworkInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&bitcoincontrollerv1.BitcoinNetwork{}, f.defaultInformer)
}

func (f *bitcoinNetworkInformer) Lister() v1.BitcoinNetworkLister {
	return v1.NewBitcoinNetworkLister(f.Informer().GetIndexer())
}
