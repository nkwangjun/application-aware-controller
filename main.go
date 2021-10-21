package main

import (
	"flag"
	"time"

	clientset "k8s.io/application-aware-controller/pkg/client/clientset/versioned"
	informers "k8s.io/application-aware-controller/pkg/client/informers/externalversions"
	aacontroller "k8s.io/application-aware-controller/pkg/controller"
	"k8s.io/application-aware-controller/pkg/signals"
	"k8s.io/application-aware-controller/pkg/trafficstrategy"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

var (
	masterURL           string
	kubeconfig          string
	trafficStrategyAddr string
	forecastWindow      int64
	nodepoolConfig      string
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	aacClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %s", err.Error())
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	aacInformerFactory := informers.NewSharedInformerFactory(aacClient, time.Second*30)

	controller := aacontroller.NewController(kubeClient, aacClient,
		aacInformerFactory.Appawarecontroller().V1().AppawareHorizontalPodAutoscalers())

	// notice that there is no need to run Start methods in a separate goroutine.
	kubeInformerFactory.Start(stopCh)
	aacInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}

	server, err := trafficstrategy.NewWebServer(trafficStrategyAddr, forecastWindow, nodepoolConfig)
	if err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}

	// start traffic strategy webserver
	go func() {
		server.Serve()
	}()
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&trafficStrategyAddr, "traffic-strategy-bind-address", ":6060", "The address the traffic strategy endpoint binds to.")
	flag.Int64Var(&forecastWindow, "forecast-window", 10, "The forecast window for traffic strategy.")
	flag.StringVar(&nodepoolConfig, "nodepool-config", "", "Path to a nodepoolconfig.")
}
