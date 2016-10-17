package scalar

import (
  "log"
  "fmt"

  "github.com/noahdietz/scalar/utils"

  k8s "k8s.io/client-go/1.4/kubernetes"
  "k8s.io/client-go/1.4/pkg/apis/extensions/v1beta1"
  "k8s.io/client-go/1.4/pkg/apis/autoscaling/v1"
  "k8s.io/client-go/1.4/pkg/api"
  apiV1 "k8s.io/client-go/1.4/pkg/api/v1"
)

// DeleteHPA deletes the HorizontalPodAutoscaler specified by the given name and namespace
func DeleteHPA(client *k8s.Clientset, name string, namespace string) (err error) {
  log.Printf("Deleting horizontal pod autoscaler %s in %s\n", name, namespace)
  return client.Autoscaling().HorizontalPodAutoscalers(namespace).Delete(name, &api.DeleteOptions{})
}

// CreateDeploymentHPA is used to create a horizontal pod autoscaler for the given deployment
func CreateDeploymentHPA(client *k8s.Clientset, config *utils.Config, deployment *v1beta1.Deployment) (err error) {
  log.Printf("Creating horizontal pod autoscaler for deployment %s in %s\n", deployment.ObjectMeta.Name, deployment.ObjectMeta.Namespace)

  hpa := buildHPA(deployment.ObjectMeta.Name,
    "Deployment",
    "extensions/v1beta1",
    config.MinReplicas,
    config.MaxReplicas,
    config.TargetCpu)

  _, err = client.Autoscaling().HorizontalPodAutoscalers(deployment.ObjectMeta.Namespace).Create(hpa)

  return
}

// CreateReplicationControllerHPA is used to create a horizontal pod autoscaler for the given replication controller
func CreateReplicationControllerHPA(client *k8s.Clientset, config *utils.Config, rc *api.ReplicationController) (err error) {
  log.Printf("Creating horizontal pod autoscaler for replication controller %s in %s\n", rc.ObjectMeta.Name, rc.ObjectMeta.Namespace)

  hpa := buildHPA(rc.ObjectMeta.Name,
    "ReplicationController",
    "v1",
    config.MinReplicas,
    config.MaxReplicas,
    config.TargetCpu)

    _, err = client.Autoscaling().HorizontalPodAutoscalers(rc.ObjectMeta.Namespace).Create(hpa)

  return
}

// PrintStatus prints the status of each HPA
func PrintStatus(client *k8s.Clientset, cache *Cache) (err error) {
  var obsvGen int64
  var currCPU int32

  for namespace := range cache.Scalars {
    list, err := client.Autoscaling().HorizontalPodAutoscalers(namespace).List(api.ListOptions{})
    if err != nil {
      return err
    }

    for _, hpa := range list.Items {
      obsvGenPtr := hpa.Status.ObservedGeneration
      if obsvGenPtr != nil {
        obsvGen = *obsvGenPtr
      }

      currCPUPtr := hpa.Status.CurrentCPUUtilizationPercentage
      if currCPUPtr != nil {
        currCPU = *currCPUPtr
      }

      fmt.Printf("Status for HPA \"%s\" in %s: ", hpa.ObjectMeta.Name, namespace)
      fmt.Printf("| ObservedGeneration: %d | LastScaleTime: %v | CurrentReplicas: %d | DesiredReplicas: %d | CurrentCPUUtilizationPercentage: %d |\n",
        obsvGen, hpa.Status.LastScaleTime, hpa.Status.CurrentReplicas, hpa.Status.DesiredReplicas, currCPU)
    }
  }

  return nil
}

func buildHPA(name string, kind string, apiV string, min int, max int, targetCpu int) (hpa *v1.HorizontalPodAutoscaler) {
  min32 := int32(min)
  max32 := int32(max)
  cpu32 := int32(targetCpu)

  return &v1.HorizontalPodAutoscaler{
    ObjectMeta: apiV1.ObjectMeta{
      Name: name,
    },
    Spec: v1.HorizontalPodAutoscalerSpec{
      ScaleTargetRef: v1.CrossVersionObjectReference{
        Kind: kind,
        Name: name,
        APIVersion: apiV,
      },
      MinReplicas: &min32,
      MaxReplicas: max32,
      TargetCPUUtilizationPercentage: &cpu32,
    },
  }
}