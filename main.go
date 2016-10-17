package main

import (
  "log"
  "time"

  "github.com/noahdietz/scalar/utils"
  "github.com/noahdietz/scalar/scalar"

  "k8s.io/client-go/1.4/pkg/api"
  "k8s.io/client-go/1.4/pkg/apis/extensions/v1beta1"
  k8s "k8s.io/client-go/1.4/kubernetes"
  "k8s.io/client-go/1.4/pkg/watch"
)

func main() {
  client, err := utils.GetClient()
  if err != nil {
    log.Fatal(err)
  }

  config, err := utils.GetConfig()
  if err != nil {
    log.Fatal(err)
  }

  cache, err := initCache(client)
  if err != nil {
    log.Fatal(err)
  }

  deploymentWatcher, rcWatcher, err := initScalar(client, config)
  if err != nil {
    log.Fatal(err)
  }

  log.Println("Scalar is configured and ready to scale!")

  var doRestart bool

  for {
    select {
    case depEvent, ok := <-deploymentWatcher.ResultChan():
      if !ok {
        doRestart = true
      } else {
        deployment := depEvent.Object.(*v1beta1.Deployment)

        switch depEvent.Type {
        case watch.Added:
          if !cache.Contains(deployment.ObjectMeta.Name, deployment.ObjectMeta.Namespace) {
            err := scalar.CreateDeploymentHPA(client, config, deployment)
            if err != nil {
              log.Printf("Error creating HPA for %s: %v\n", deployment.ObjectMeta.Name, err)
            }

            cache.Add(deployment.ObjectMeta.Name, deployment.ObjectMeta.Namespace)
          }
        case watch.Deleted:
          err := scalar.DeleteHPA(client, deployment.ObjectMeta.Name, deployment.ObjectMeta.Namespace)
          if err != nil {
            log.Printf("Error deleting HPA for %s in %s: %v\n", deployment.ObjectMeta.Name, deployment.ObjectMeta.Namespace, err)
          }

          cache.Remove(deployment.ObjectMeta.Name, deployment.ObjectMeta.Namespace)
        }
      }

    case rcEvent, ok := <-rcWatcher.ResultChan():
      if !ok {
        doRestart = true
      } else {
        rc := rcEvent.Object.(*api.ReplicationController)

        switch rcEvent.Type {
        case watch.Added:
          if !cache.Contains(rc.ObjectMeta.Name, rc.ObjectMeta.Namespace) {
            err := scalar.CreateReplicationControllerHPA(client, config, rc)
            if err != nil {
              log.Printf("Error creating HPA for replication controller %s in %s: %v\n", rc.ObjectMeta.Name, rc.ObjectMeta.Namespace, err)
            }

            cache.Add(rc.ObjectMeta.Name, rc.ObjectMeta.Namespace)
          }
        case watch.Deleted:
          err := scalar.DeleteHPA(client, rc.ObjectMeta.Name, rc.ObjectMeta.Namespace)
          if err != nil {
            log.Printf("Error deleting HPA for %s in %s: %v\n", rc.ObjectMeta.Name, rc.ObjectMeta.Namespace, err)
          }

          cache.Remove(rc.ObjectMeta.Name, rc.ObjectMeta.Namespace)
        }
      }

    case <- time.After(config.StatusTimer * time.Second):
      if config.PrintStatus {
        err := scalar.PrintStatus(client, cache)
        if err != nil {
          log.Printf("Error printing status: %v\n", err)
        }
      }
    }

    if doRestart {
      log.Println("Restarting watchers.")
      deploymentWatcher, rcWatcher, err = initScalar(client, config)
      if err != nil {
        log.Fatal(err)
      }

      doRestart = false
    }
  }
}

func initScalar(client *k8s.Clientset, config *utils.Config) (deploymentWatcher watch.Interface, rcWatcher watch.Interface, err error) {
  deploymentWatcher, err = client.Extensions().Deployments(api.NamespaceAll).Watch(api.ListOptions{
    LabelSelector: config.ScalarSelector,
  })

  if err != nil {
    return nil, nil, err
  }

  rcWatcher, err = client.Core().ReplicationControllers(api.NamespaceAll).Watch(api.ListOptions{
    LabelSelector: config.ScalarSelector,
  })

  if err != nil {
    return nil, nil, err
  }

  return
}

func initCache(client *k8s.Clientset) (cache *scalar.Cache, err error) {
  cache = scalar.InitCache()

  hpaList, err := client.Autoscaling().HorizontalPodAutoscalers(api.NamespaceAll).List(api.ListOptions{})
  if err != nil {
    return nil, err
  }

  for _, hpa := range hpaList.Items {
    cache.Add(hpa.ObjectMeta.Name, hpa.ObjectMeta.Namespace)
  }

  return cache, nil
}