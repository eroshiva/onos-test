package v1beta1

import (
	clustercorev1 "github.com/onosproject/onos-test/pkg/onit/cluster/core/v1"
	clustermetav1 "github.com/onosproject/onos-test/pkg/onit/cluster/meta/v1"
	appsv1beta1 "k8s.io/api/apps/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

var DeploymentKind = clustermetav1.Kind{
	Group:   "apps",
	Version: "v1beta1",
	Kind:    "Deployment",
}

var DeploymentResource = clustermetav1.Resource{
	Kind: DeploymentKind,
	Name: "Deployment",
	ObjectFactory: func() runtime.Object {
		return &appsv1beta1.Deployment{}
	},
	ObjectsFactory: func() runtime.Object {
		return &appsv1beta1.DeploymentList{}
	},
}

type DeploymentsClient interface {
	Get(name string) (*Deployment, error)
	List() ([]*Deployment, error)
}

// newDeploymentsClient creates a new DeploymentsClient
func newDeploymentsClient(objects clustermetav1.ObjectsClient) DeploymentsClient {
	return &deploymentsClient{
		ObjectsClient: objects,
	}
}

// deploymentsClient implements the DeploymentsClient interface
type deploymentsClient struct {
	clustermetav1.ObjectsClient
}

func (c *deploymentsClient) Get(name string) (*Deployment, error) {
	object, err := c.ObjectsClient.Get(name, DeploymentResource)
	if err != nil {
		return nil, err
	}
	return newDeployment(object), nil
}

func (c *deploymentsClient) List() ([]*Deployment, error) {
	objects, err := c.ObjectsClient.List(DeploymentResource)
	if err != nil {
		return nil, err
	}
	deployments := make([]*Deployment, len(objects))
	for i, object := range objects {
		deployments[i] = newDeployment(object)
	}
	return deployments, nil
}

// newDeployment creates a new Deployment resource
func newDeployment(object *clustermetav1.Object) *Deployment {
	return &Deployment{
		PodSet: clustercorev1.NewPodSet(object),
		Spec:   object.Object.(*appsv1beta1.Deployment).Spec,
	}
}

// Deployment provides functions for querying a deployment
type Deployment struct {
	*clustercorev1.PodSet
	Spec appsv1beta1.DeploymentSpec
}