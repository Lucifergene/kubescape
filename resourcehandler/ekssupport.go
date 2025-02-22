package resourcehandler

import (
	"fmt"
	"os"
	"strings"

	"github.com/armosec/k8s-interface/k8sinterface"
)

type EKSProviderEnvVar struct {
}

func NewEKSProviderEnvVar() *EKSProviderEnvVar {
	return &EKSProviderEnvVar{}
}

func (eksProviderEnvVar *EKSProviderEnvVar) getKubeClusterName() string {
	return eksProviderEnvVar.getKubeCluster()
}

func (eksProviderEnvVar *EKSProviderEnvVar) getKubeCluster() string {
	val, present := os.LookupEnv(KS_KUBE_CLUSTER_ENV_VAR)
	if present {
		return val
	}
	return ""
}

func (eksProviderEnvVar *EKSProviderEnvVar) getRegion(cluster string, provider string) (string, error) {
	return eksProviderEnvVar.getRegionForEKS(cluster)
}

func (eksProviderEnvVar *EKSProviderEnvVar) getProject(cluster string, provider string) (string, error) {
	return "", nil
}

func (eksProviderEnvVar *EKSProviderEnvVar) getRegionForEKS(cluster string) (string, error) {
	region, present := os.LookupEnv(KS_CLOUD_REGION_ENV_VAR)
	if present {
		return region, nil
	}
	splittedClusterContext := strings.Split(cluster, ".")
	if len(splittedClusterContext) < 2 {
		return "", fmt.Errorf("failed to get region")
	}
	region = splittedClusterContext[1]
	return region, nil
}

// ------------------------------------- EKSProviderContext -------------------------

type EKSProviderContext struct {
}

func NewEKSProviderContext() *EKSProviderContext {
	return &EKSProviderContext{}
}

func (eksProviderContext *EKSProviderContext) getKubeClusterName() string {
	context := k8sinterface.GetCurrentContext()
	if context == nil {
		return ""
	}
	cluster := context.Cluster
	if cluster != "" {
		splittedCluster := strings.Split(cluster, ".")
		if len(splittedCluster) > 1 {
			return splittedCluster[0]
		}
	}
	splittedCluster := strings.Split(k8sinterface.GetClusterName(), ".")
	if len(splittedCluster) > 1 {
		return splittedCluster[0]
	}
	return ""
}

func (eksProviderContext *EKSProviderContext) getKubeCluster() string {
	context := k8sinterface.GetCurrentContext()
	if context == nil {
		return ""
	}
	if context.Cluster != "" {
		return context.Cluster
	}
	return k8sinterface.GetClusterName()
}

func (eksProviderContext *EKSProviderContext) getRegion(cluster string, provider string) (string, error) {
	splittedClusterContext := strings.Split(cluster, ".")
	if len(splittedClusterContext) < 2 {
		return "", fmt.Errorf("failed to get region")
	}
	region := splittedClusterContext[1]
	return region, nil
}

func (eksProviderContext *EKSProviderContext) getProject(cluster string, provider string) (string, error) {
	return "", nil
}
