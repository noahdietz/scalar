package utils

import (
  "k8s.io/client-go/1.4/kubernetes"
  "k8s.io/client-go/1.4/rest"
)

// GetClient returns a Kubernetes client from in cluster configuration
func GetClient() (*kubernetes.Clientset, error) {
  config, err := rest.InClusterConfig()
  if err != nil {
    return nil, err
  }

  return kubernetes.NewForConfig(config)
}