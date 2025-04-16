package main

import (
	"context"
	"fmt"
	sspclientset "github.com/nevidanniu/sample-apispec/client-go/ssp"
	ldapv1alpha1 "github.com/nevidanniu/sample-apispec/ldap/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func Main() {
	configPath := ""
	kubeconfig, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		fmt.Errorf("Unable to load kubeconfig from %s: %v", configPath, err)
		os.Exit(1)
	}

	client, err := sspclientset.NewForConfig(kubeconfig)
	if err != nil {
		fmt.Errorf("Unable to load clienset from kubeconfig: %v", err)
		os.Exit(1)
	}
	client.CoreV1alpha1().Tenants().List(context.Background(), v1.ListOptions{})
	client.LdapV1alpha1().Users().Create(context.Background(), &ldapv1alpha1.User{
		ObjectMeta: v1.ObjectMeta{Name: "FIO"},
	}, v1.CreateOptions{})
}
