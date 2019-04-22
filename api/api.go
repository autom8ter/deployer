package api

import (
	"k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type Creater func(*v1.Deployment) (*v1.Deployment, error)
type Deleter func(name string, options *metav1.DeleteOptions) error
type Updater func(*v1.Deployment) (*v1.Deployment, error)
type Getter func(name string, options metav1.GetOptions) (*v1.Deployment, error)
type Lister func(opts metav1.ListOptions) (*v1.DeploymentList, error)
type Watcher func(opts metav1.ListOptions) (watch.Interface, error)
