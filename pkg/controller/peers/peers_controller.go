package peers

import (
	"context"

	appv1alpha1 "cello/k8s-fabric-operator/pkg/apis/app/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_peers")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Peers Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcilePeers{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("peers-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Peers
	err = c.Watch(&source.Kind{Type: &appv1alpha1.Peers{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner Peers
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &appv1alpha1.Peers{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcilePeers implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcilePeers{}

// ReconcilePeers reconciles a Peers object
type ReconcilePeers struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Peers object and makes changes based on the state read
// and what is in the Peers.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcilePeers) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Peers")

	// Fetch the Peers instance
	instance := &appv1alpha1.Peers{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	lbls := labels.Set{
		"app":     instance.Metadata.Labels.App,
		"role":     instance.Metadata.Labels.Role,
		"peerId":     instance.Metadata.Labels.PeerId,
		"org":     instance.Metadata.Labels.Org,
	}
	existingPeers := &corev1.Peers{}
	err = r.client.List(context.TODO(),
		&client.ListOptions{
			Namespace:     request.Namespace,
			LabelSelector: labels.SelectorFromSet(lbls),
		},
		existingPeers)
	if err != nil {
		reqLogger.Error(err, "failed to list existing pods in the Peers")
		return reconcile.Result{}, err
	}

	existingPeersName := []string{}

	for _, peer := range existingPeers.Items {
		if peer.GetObjectMeta().GetDeletionTimestamp() != nil {
			continue
		}
		if peer.Status.Phase == corev1.PodPending || peer.Status.Phase == corev1.PodRunning {
			existingPeersNames = append(existingPeersNames, peer.GetObjectMeta().GetName())
		}
	}

	reqLogger.Info("Checking podset", "expected replicas", instance.Spec.Replicas, "Pod.Names", existingPeersNames)

	if int32(len(existingPeersNames)) > instance.Spec.Replicas {
		// delete a pod. Just one at a time (this reconciler will be called again afterwards)
		reqLogger.Info("Deleting a pod in the podset", "expected replicas", instance.Spec.Replicas, "Pod.Names", existingPeersNames)
		peer := existingPeers.Items[0]
		err = r.client.Delete(context.TODO(), &peer)
		if err != nil {
			reqLogger.Error(err, "failed to delete a pod")
			return reconcile.Result{}, err
		}
	}

	// Scale Up Pods
	if int32(len(existingPeersNames)) < instance.Spec.Replicas {
		// create a new pod. Just one at a time (this reconciler will be called again afterwards)
		reqLogger.Info("Adding a pod in the podset", "expected replicas", instance.Spec.Replicas, "Pod.Names", existingPeersNames)
		pod := newPodForCR(instance)
		if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
			reqLogger.Error(err, "unable to set owner reference on new pod")
			return reconcile.Result{}, err
		}
		err = r.client.Create(context.TODO(), pod)
		if err != nil {
			reqLogger.Error(err, "failed to create a pod")
			return reconcile.Result{}, err
		}
	}
	return reconcile.Result{Requeue: true}, nil
}

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newPodForCR(cr *appv1alpha1.Peers) *corev1.Pod {
	labels := map[string]string{
		"app":     cr.Metadata.Labels.App,
		"role":     cr.Metadata.Labels.Role,
		"peerId":     cr.Metadata.Labels.PeerId,
		"org":     cr.Metadata.Labels.Org,

	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name + "-pod",
			Namespace: cr.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
