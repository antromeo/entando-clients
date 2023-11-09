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
	pluginLinkName := "placeholder-pn-4ddb649c-937377f0-b22ce35d-entandodemo-conf-link"

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

	pluginLink, err := clientset.EntandoV1().EntandoAppPluginLinks(namespace).Get(context.TODO(), pluginLinkName, metav1.GetOptions{})
	fmt.Printf("%+v\n", pluginLink)

	pluginLinkList, err := clientset.EntandoV1().EntandoAppPluginLinks(namespace).List(context.TODO(), metav1.ListOptions{})
	fmt.Println("List of EntandoAppPluginLinks")
	for _, pluginLink := range pluginLinkList.Items {
		fmt.Println(pluginLink.Name)
	}
}
