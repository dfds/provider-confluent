package controller

import (
	"context"
	"path/filepath"
	"testing"

	v1alpha1sa "github.com/dfds/provider-confluent/apis/serviceaccount/v1alpha1"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func kubeclient() *rest.Config {
	var kubeconfig string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	return config
}

func crdConfig(crdSchemeGroupVersion *schema.GroupVersion) *rest.Config {
	config := kubeclient()
	// config.ContentConfig.GroupVersion = &v1alpha1sa.SchemeGroupVersion
	config.ContentConfig.GroupVersion = crdSchemeGroupVersion
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.NewCodecFactory(scheme.Scheme)
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	return config
}

func TestServiceAccountController(t *testing.T) {
	assert := assert.New(t)

	saRestClient, err := rest.UnversionedRESTClientFor(crdConfig(&v1alpha1sa.SchemeGroupVersion))
	if err != nil {
		panic(err)
	}

	var result v1alpha1sa.ServiceAccountList
	err = saRestClient.Get().Resource("serviceaccounts").Do(context.TODO()).Into(&result)
	if err != nil {
		t.Errorf("cannot get serviceaccount resources %s:", err.Error())
	} else {
		assert.Len(result.Items, 1, "expect only one service account")
	}
}
