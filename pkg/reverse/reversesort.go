package reverse

import (
	"k8s.io/apimachinery/pkg/runtime"
<<<<<<< HEAD
	"k8s.io/kubernetes/pkg/api/v1/pod"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
	"k8s.io/klog"
)


=======
	"k8s.io/klog"
	"k8s.io/kubernetes/pkg/api/v1/pod"
	framework "k8s.io/kubernetes/pkg/scheduler/framework/v1alpha1"
)

>>>>>>> 882d9b4... initial commit
const Name = "ReverseSort"

// ReverseSort is a plugin that reverts the default sort
type Sort struct{}
var _ framework.QueueSortPlugin = &Sort{}

// Name returns name of the plugin.
func (_ *Sort) Name() string {
	return Name
}

// Less is the function used by the activeQ heap algorithm to sort pods.
// It sorts pods based on their priority.
func (*Sort) Less(pInfo1, pInfo2 *framework.PodInfo) bool {
	p1 := pod.GetPodPriority(pInfo1.Pod)
	p2 := pod.GetPodPriority(pInfo2.Pod)
	klog.V(1).Infof("comparing priorities %s < %s", p1, p2)
	return p1 < p2
}

// New initializes a new plugin and returns it.
func New(_ *runtime.Unknown, _ framework.FrameworkHandle) (framework.Plugin, error) {
	return &Sort{}, nil
}