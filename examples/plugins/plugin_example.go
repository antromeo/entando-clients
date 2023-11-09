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
	pluginName := "pn-4ddb649c-937377f0-b22ce35d-entandodemo-conference-ms"

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

	plugin, err := clientset.EntandoV1().EntandoPlugins(namespace).Get(context.TODO(), pluginName, metav1.GetOptions{})
	fmt.Printf("%+v\n", plugin)

	pluginList, err := clientset.EntandoV1().EntandoPlugins(namespace).List(context.TODO(), metav1.ListOptions{})
	fmt.Println("List of Entando Plugins")
	for _, plugin := range pluginList.Items {
		fmt.Println(plugin.Name)
	}
}
