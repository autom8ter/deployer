package deployer

import (
	"github.com/autom8ter/gcloud/clients"
	"github.com/autom8ter/objectify"
	"k8s.io/api/apps/v1"
	v13 "k8s.io/api/core/v1" //https://godoc.org/k8s.io/api/core/v1
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apps "k8s.io/client-go/kubernetes/typed/apps/v1"
)

const DefaultNamespace = v13.NamespaceDefault

var util = objectify.Default()

//Client holds an authenticated kubernetes client set
type Deployer struct {
	api apps.DeploymentInterface
}

//NewDeployer returns a Deployer with a kubernetes client set based on whether in/out of cluster
func NewDeployer(incluster bool, namespace string) (*Deployer, error) {
	cli, err := clients.NewKubernetesClientSet(incluster)
	if err != nil {
		return nil, err
	}
	cli.AppsV1().Deployments(namespace)
	return &Deployer{
		api: cli.AppsV1().Deployments(namespace),
	}, nil
}

func (d *Deployer) Validate() error {
	return util.Validate(d)
}

func (d *Deployer) API() apps.DeploymentInterface {
	return d.api
}

func (d *Deployer) CreateDeployment(name string, replicas int, labels map[string]string, containers []*Container) (*v1.Deployment, error) {
	tainers := []v13.Container{}
	for _, c := range containers {
		tainers = append(tainers, c.toConatainer())
	}
	return d.API().Create(&v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1.DeploymentSpec{
			Replicas: int32Ptr(int32(replicas)),
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: v13.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: v13.PodSpec{
					Volumes:        nil,
					InitContainers: nil,
					Containers:     tainers,
				},
			},
		},
	})
}

type ContainerOption func(c *Container)

type Container struct {
	Name          string
	Image         string
	Command       []string
	Args          []string
	Directory     string
	ContainerPort int
	HostPort      int
	Env           map[string]string
}

func NewContainer(opts ...ContainerOption) *Container {
	c := &Container{}
	for _, o := range opts {
		o(c)
	}
	return c
}

func (c *Container) toConatainer() v13.Container {
	vars := []v13.EnvVar{}
	for k, v := range c.Env {
		vars = append(vars, v13.EnvVar{
			Name:  k,
			Value: v,
		})
	}
	return v13.Container{
		Name:       c.Name,
		Image:      c.Image,
		Command:    c.Command,
		Args:       c.Args,
		WorkingDir: c.Directory,
		Ports: []v13.ContainerPort{
			{
				Name:          c.Name + "_port",
				HostPort:      int32(c.HostPort),
				ContainerPort: int32(c.ContainerPort),
				Protocol:      "http",
			},
		},
		Env: vars,
	}
}
func int32Ptr(i int32) *int32 { return &i }
