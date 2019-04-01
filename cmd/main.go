package main

import (
	"flag"

	"github.com/golang/glog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var (
	masterURL = flag.String("master", "",
		"The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	kubeconfig = flag.String("kubeconfig", "",
		"Path to a kubeconfig. Only required if out-of-cluster.")
	svc1 = flag.String("svc1", "", "The first k8s service.")
	svc2 = flag.String("svc2", "", "The second k8s service.")
)

const defNS = "default"

func main() {
	flag.Parse()
	glog.Info("Initializing the service shuffler program")
	if *svc1 == "" {
		glog.Fatal("svc1 not provided")
	}
	if *svc2 == "" {
		glog.Fatal("svc2 not provided")
	}

	clusterConfig, err := clientcmd.BuildConfigFromFlags(*masterURL, *kubeconfig)
	if err != nil {
		glog.Fatalf("Error getting cluster configuration: %v", err)
	}
	kubeClient, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		glog.Fatalf("Error building new kubernetes client: %v", err)
	}

	svc1, err := kubeClient.CoreV1().Services(defNS).Get(*svc1, metav1.GetOptions{})
	if err != nil {
		glog.Fatalf("Error obtaining svc1: %v", err)
	}
	glog.Infof("Obtained svc1: %#v", svc1)
}
