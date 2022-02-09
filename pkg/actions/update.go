package actions

import (
	"context"

	"github.com/MiniTeks/mks-server/pkg/tconfig"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func Update(
	grc schema.GroupVersionResource, clients *tconfig.Client,
	object *unstructured.Unstructured, ns string,
	opt metav1.UpdateOptions) (*unstructured.Unstructured, error) {
	gvr, err := GetGroupVersionResource(grc, clients.Tekton.Discovery())
	if err != nil {
		return nil, err
	}

	obj, err := clients.Dynamic.Resource(*gvr).Namespace(ns).Update(context.Background(), object, opt)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
