package main

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"

	// client
	client, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	// get data
	svc := v1.Service{}

	err = client.Get().Namespace("kube-system").Resource("services").Name("kube-dns").Do(context.TODO()).Into(&svc)
	if err != nil {
		println("service get error", err)
	} else {
		println("service name: ", svc.Name)
	}

}
