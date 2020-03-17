// Code generated by onit-generate. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/onos-test/pkg/helm/api/resource"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"time"
)

type RoleBindingsReader interface {
	Get(name string) (*RoleBinding, error)
	List() ([]*RoleBinding, error)
}

func NewRoleBindingsReader(client resource.Client, filter resource.Filter) RoleBindingsReader {
	return &roleBindingsReader{
		Client: client,
		filter: filter,
	}
}

type roleBindingsReader struct {
	resource.Client
	filter resource.Filter
}

func (c *roleBindingsReader) Get(name string) (*RoleBinding, error) {
	roleBinding := &rbacv1.RoleBinding{}
	err := c.Clientset().
		RbacV1().
		RESTClient().
		Get().
		NamespaceIfScoped(c.Namespace(), RoleBindingKind.Scoped).
		Resource(RoleBindingResource.Name).
		Name(name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Into(roleBinding)
	if err != nil {
		return nil, err
	} else {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   RoleBindingKind.Group,
			Version: RoleBindingKind.Version,
			Kind:    RoleBindingKind.Kind,
		}, roleBinding.ObjectMeta)
		if err != nil {
			return nil, err
		} else if !ok {
			return nil, errors.NewNotFound(schema.GroupResource{
				Group:    RoleBindingKind.Group,
				Resource: RoleBindingResource.Name,
			}, name)
		}
	}
	return NewRoleBinding(roleBinding, c.Client), nil
}

func (c *roleBindingsReader) List() ([]*RoleBinding, error) {
	list := &rbacv1.RoleBindingList{}
	err := c.Clientset().
		RbacV1().
		RESTClient().
		Get().
		Namespace(c.Namespace()).
		Resource(RoleBindingResource.Name).
		VersionedParams(&metav1.ListOptions{}, metav1.ParameterCodec).
		Timeout(time.Minute).
		Do().
		Into(list)
	if err != nil {
		return nil, err
	}

	results := make([]*RoleBinding, 0, len(list.Items))
	for _, roleBinding := range list.Items {
		ok, err := c.filter(metav1.GroupVersionKind{
			Group:   RoleBindingKind.Group,
			Version: RoleBindingKind.Version,
			Kind:    RoleBindingKind.Kind,
		}, roleBinding.ObjectMeta)
		if err != nil {
			return nil, err
		} else if ok {
			copy := roleBinding
			results = append(results, NewRoleBinding(&copy, c.Client))
		}
	}
	return results, nil
}