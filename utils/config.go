package utils

import (
  "os"
  "log"
  "strconv"
  "time"

  "k8s.io/client-go/1.4/pkg/labels"
)

// Config represents a Scalar configuration
type Config struct {
  ScalarSelector labels.Selector
  MinReplicas int
  MaxReplicas int
  TargetCpu int
  PrintStatus bool
  StatusTimer time.Duration
}

const (
  defaultMinReplicas = 2
  defaultMaxReplicas = 8
  defaultTargetCpu = 75
  defaultStatusTimer = 1800
  defaultPrintStatus = true
)

// GetConfig returns a new Scalar configuration
func GetConfig() (config *Config, err error) {
  min := defaultMinReplicas
  max := defaultMaxReplicas
  targetCpu := defaultTargetCpu
  timer := defaultStatusTimer
  printStatus := defaultPrintStatus

  config = &Config{}

  selector := os.Getenv("SCALAR_SELECTOR")
  if selector != "" {
    log.Printf("Using selector: %s\n", selector)
  }

  config.ScalarSelector, err = labels.Parse(selector)
  if err != nil {
    return nil, err
  }


  if minRepStr := os.Getenv("SCALAR_MIN_REPLICAS"); minRepStr != "" {
    min, err = strconv.Atoi(minRepStr)
    if err != nil {
      return nil, err
    }
  }

  config.MinReplicas = min

  if maxRepStr := os.Getenv("SCALAR_MAX_REPLICAS"); maxRepStr != "" {
    max, err = strconv.Atoi(maxRepStr)
    if err != nil {
      return nil, err
    }
  }

  config.MaxReplicas = max

  if targetCpuStr := os.Getenv("SCALAR_TARGET_CPU"); targetCpuStr != "" {
    targetCpu, err = strconv.Atoi(targetCpuStr)
    if err != nil {
      return nil, err
    }
  }

  config.TargetCpu = targetCpu

  if printStatusStr := os.Getenv("SCALAR_PRINT_STATUS"); printStatusStr != "" {
    printStatus, err = strconv.ParseBool(printStatusStr)
    if err != nil {
      return nil, err
    }
  }

  config.PrintStatus = printStatus

  if timerStr := os.Getenv("SCALAR_STATUS_TIMER"); timerStr != "" {
    timer, err = strconv.Atoi(timerStr)
    if err != nil {
      return nil, err
    }
  }

  config.StatusTimer = time.Duration(timer)

  return
}