package main

import (
	"context"
	"flag"
	"fmt"
	entandoclientset "github.com/antromeo/entando-clients/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {

	namespace := "placeholder"
	bundleName := "placeholder"

	home := homedir.HomeDir()
	kubeconfig := flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("erorr %s building config from flags\n", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error %s, getting inclusterconfig", err.Error())
		}
	}

	clientset, err := entandoclientset.NewForConfig(config)

	debundle, err := clientset.EntandoV1().EntandoDeBundles(namespace).Get(context.TODO(), bundleName, metav1.GetOptions{})
	fmt.Printf("%+v\n", debundle)

	debundleList, err := clientset.EntandoV1().EntandoDeBundles(namespace).List(context.TODO(), metav1.ListOptions{})
	fmt.Println("List of Entando Bundles")
	for _, debundle := range debundleList.Items {
		fmt.Println(debundle.Name)
	}
}
