package controller

import (
	"cello/k8s-fabric-operator/pkg/controller/peers"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, peers.Add)
}
