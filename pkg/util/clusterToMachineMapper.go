package util

import (
	"context"

	clusterv1alpha1 "github.com/samsung-cnct/ims-kaas/pkg/apis/cluster/v1alpha1"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// A mapper that returns all the Registries
type ClusterToMachineMapper struct {
	client.Client
}

func (m ClusterToMachineMapper) Map(obj handler.MapObject) []reconcile.Request {
	var res []reconcile.Request

	cluster, ok := obj.Object.(*clusterv1alpha1.CnctCluster)
	if !ok {
		return res // This wasn't a Cluster
	}

	machines := &clusterv1alpha1.CnctMachineList{}
	if err := m.List(context.Background(), &client.ListOptions{}, machines); err != nil {
		klog.Errorf("could not get list of machines: %q", err)
		return res
	}

	// Add all the Machines that are members of this Cluster
	for _, machine := range machines.Items {
		clusterName := GetClusterNameFromMachineOwnerRef(&machine)
		if clusterName == cluster.GetName() {
			res = append(res, reconcile.Request{
				NamespacedName: types.NamespacedName{
					Name:      machine.GetName(),
					Namespace: machine.GetNamespace(),
				},
			})
		}
	}
	return res
}
