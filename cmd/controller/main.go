package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

var (
	masterURL  string
	kubeconfig string
)

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

func createSignalHandler() (stopCh <-chan struct{}) {
	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		sig := <-c
		fmt.Printf("Signal handler: received signal %s\n", sig)
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

func main() {
	klog.InitFlags(nil)
	klog.Info("Starting up")

	//Establish a signal handler - this will return a channel which
	//will be closed if a signal is received
	stopCh := createSignalHandler()

	//Get command line arguments - kubeconfig = kube config file to use
	flag.StringVar(&kubeconfig, "kubeconfig", "", "The kubectl configuration file to use")
	flag.Parse()
	if kubeconfig != "" {
		klog.Infof("Trying standalone configuration with kubeconfig config file %s\n", kubeconfig)
	}

	//Create client cmd
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		klog.Error("Could not get config")
		os.Exit(1)
	}

	//Create clientset from config
	_, err = kubernetes.NewForConfig(config)
	if err != nil {
		klog.Error("Could not create clientset")
		os.Exit(1)
	}

	klog.Info("Entering main loop")
	<-stopCh
}
