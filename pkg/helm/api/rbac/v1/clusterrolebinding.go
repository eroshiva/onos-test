// Code generated by onit-generate. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/onos-test/pkg/helm/api/resource"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

var ClusterRoleBindingKind = resource.Kind{
	Group:   "rbac.authorization.k8s.io",
	Version: "v1",
	Kind:    "ClusterRoleBinding",
	Scoped:  false,
}

var ClusterRoleBindingResource = resource.Type{
	Kind: ClusterRoleBindingKind,
	Name: "clusterrolebindings",
}

func NewClusterRoleBinding(clusterRoleBinding *rbacv1.ClusterRoleBinding, client resource.Client) *ClusterRoleBinding {
	return &ClusterRoleBinding{
		Resource: resource.NewResource(clusterRoleBinding.ObjectMeta, ClusterRoleBindingKind, client),
		Object:   clusterRoleBinding,
	}
}

type ClusterRoleBinding struct {
	*resource.Resource
	Object *rbacv1.ClusterRoleBinding
}

func (r *ClusterRoleBinding) Delete() error {
	return r.Clientset().
		RbacV1().
		RESTClient().
		Delete().
		NamespaceIfScoped(r.Namespace, ClusterRoleBindingKind.Scoped).
		Resource(ClusterRoleBindingResource.Name).
		Name(r.Name).
		VersionedParams(&metav1.DeleteOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Error()
}
