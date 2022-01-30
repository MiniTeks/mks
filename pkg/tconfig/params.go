package tconfig

import (
	"net/http"

	tektonclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	versionedResource "github.com/tektoncd/pipeline/pkg/client/resource/clientset/versioned"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

type Client struct {
	Tekton     tektonclient.Interface
	Kube       kubernetes.Interface
	Resource   versionedResource.Interface
	HTTPclient http.Client
	Dynamic    dynamic.Interface
}

type Params interface {
	SetKubeConfigPath(string)

	SetKubeContext(string)

	Client() (*Client, error)

	SetNamespace(string)
}
