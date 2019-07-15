package controller

import (
	"cello/k8s-fabric-operator/pkg/controller/cas"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, cas.Add)
}
