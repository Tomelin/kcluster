package v1aplha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runetime.Object
type Kluster struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec KusterSpec
}

type KusterSpec struct {
	Name      string
	Region    string
	Version   string
	NodePools []NodePool
}

type NodePool struct {
	Size  string
	Name  string
	Count int
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runetime.Object
type KlusterList struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Items []Kluster
}
