package apiserver

import (
	"context"

	"github.com/samsung-cnct/cma-ssh/pkg/apis/cluster/common"
	"github.com/samsung-cnct/cma-ssh/pkg/generated/api"
	"github.com/samsung-cnct/cma-ssh/pkg/util"

	corev1 "k8s.io/api/core/v1"
	clientlib "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func TranslateClusterStatus(crStatus common.ClusterStatusPhase) api.ClusterStatus {

	var clusterStatus = api.ClusterStatus_STATUS_UNSPECIFIED
	switch crStatus {
	case common.UnspecifiedClusterPhase:
		clusterStatus = api.ClusterStatus_STATUS_UNSPECIFIED
	case common.ErrorClusterPhase:
		clusterStatus = api.ClusterStatus_ERROR
	case common.RunningClusterPhase:
		clusterStatus = api.ClusterStatus_RUNNING
	case common.StoppingClusterPhase:
		clusterStatus = api.ClusterStatus_STOPPING
	case common.ReconcilingClusterPhase:
		clusterStatus = api.ClusterStatus_RECONCILING

	}

	return clusterStatus
}

func GetKubeConfig(clusterName string, manager manager.Manager) ([]byte, error) {
	// get client
	client := manager.GetClient()

	// get kubeconfig from cluster secret
	clusterSecret := &corev1.Secret{}
	err := client.Get(context.Background(), clientlib.ObjectKey{
		Namespace: clusterName,
		Name:      "cluster-private-key",
	}, clusterSecret)
	if err != nil {
		return nil, err
	}

	return clusterSecret.Data[corev1.ServiceAccountKubeconfigKey], nil
}

func GetMachineName(clusterName string, hostIp string, manager manager.Manager) (string, error) {
	// get client
	client := manager.GetClient()

	// get list of cluster machines
	machineList, err := util.GetClusterMachineList(client, clusterName)
	if err != nil {
		return "", err
	}

	for _, machine := range machineList {
		if machine.Status.SshConfig.Host == hostIp {
			return machine.GetName(), nil
		}
	}

	return "", nil
}
