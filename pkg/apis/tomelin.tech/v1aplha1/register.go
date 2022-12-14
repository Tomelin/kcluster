package v1aplha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{Group: "tomelin.tech", Version: "v1alpha1"}
var SchemeBuilder runtime.SchemeBuilder
var AddToScheme = SchemeBuilder.AddToScheme

func init() {

	SchemeBuilder.Register(addKnowTypes)
}

func addKnowTypes(scheme *runtime.Scheme) error {

	scheme.AddKnownTypes(SchemeGroupVersion, &Kluster{}, &KlusterList{})

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
