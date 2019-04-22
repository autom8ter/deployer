package deployer

import (
	"github.com/autom8ter/gcloud/clients"
	"k8s.io/client-go/kubernetes"
)

//Client holds an authenticated kubernetes client set
type Deployer struct {
	set *kubernetes.Clientset
}

//NewDeployer returns a Deployer with a kubernetes client set based on whether in/out of cluster
func NewDeployer(incluster bool) (*Deployer, error) {
	cli, err := clients.NewKubernetesClientSet(incluster)
	if err != nil {
		return nil, err
	}
	return &Deployer{
		set: cli,
	}, nil
}

//Set returns a kubernetes client set
func (c *Deployer) Set() *kubernetes.Clientset {
	return c.set
}
