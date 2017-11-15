package k8s

import (
	"net/http"
)

var (
	jsonHeader  = http.Header{"Content-Type": []string{"application/json"}}
	patchHeader = http.Header{"Content-Type": []string{"application/strategic-merge-patch+json"}}
)

//configmap
var (
	configMapNew = "/api/v1/namespaces/%s/configmaps/"
	configMapOpt = "/api/v1/namespaces/%s/configmaps/%s"
)

//namespace
var (
	namespaceNew = "/api/v1/namespaces"
	namespaceOpt = "/api/v1/namespaces/%s"
)

//svc
var (
	svcNew = "/api/v1/namespaces/%s/services"
	svcOpt = "/api/v1/namespaces/%s/services/%s"
)

//ing
var (
	ingNew = "/apis/extensions/v1beta1/namespaces/%s/ingresses"
	ingOpt = "/apis/extensions/v1beta1/namespaces/%s/ingresses/%s"
)

//rc
var (
	rcNew = "/api/v1/namespaces/%s/replicationcontrollers"
	rcOpt = "/api/v1/namespaces/%s/replicationcontrollers/%s"
)

//deployment
var (
	deploymentNew = "/apis/apps/v1beta1/namespaces/%s/deployments"
	deplotmentOpt = "/apis/apps/v1beta1/namespaces/%s/deployments/%s"
)

//pvc
var (
	pvcNew = "/api/v1/namespaces/%s/persistentvolumeclaims"
	pvcOpt = "/api/v1/namespaces/%s/persistentvolumeclaims/%s"
)

//scale

//var scaleOpt = "/apis/apps/v1beta1/namespaces/%s/deployments/%s/scale"

//deployment
type Deployment struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Spec struct {
		MinReadySeconds         int  `json:"minReadySeconds "`
		Paused                  bool `json:"paused"`
		ProgressDeadlineSeconds int  `json:"progressDeadlineSeconds"`
		RevisionHistoryLimit    int  `json:"revisionHistoryLimit "`
		RollbackTo              struct {
			Revision int `json:"revision"`
		} `json:"rollbackTo"`
		Strategy struct {
			RollingUpdate struct {
				MaxSurge       int `json"maxSurge`
				MaxUnavailable int `json:"maxUnavailable"`
			} `json:"rollingUpdate "`
			Type string `json:"type"`
		} `json:"strategy "`
		Replicas int `json:"replicas"`
		Template struct {
			Metadata struct {
				Labels struct {
					App string `json:"app"`
				} `json:"labels"`
			} `json:"metadata"`
			Spec struct {
				Volumes []struct {
					Name      string `json:"name"`
					ConfigMap struct {
						Name string `json:"name"`
					} `json:"configMap"`
				} `json:"volumes"`
				RestartPolicy string `json:"restartPolicy"`
				Containers    []struct {
					Image        string `json:"image"`
					Name         string `json:"name"`
					VolumeMounts []struct {
						Name      string `json:"name"`
						MountPath string `json:"mountPath"`
					} `json:"volumeMounts"`
					Ports []struct {
						ContainerPort int `json:"containerPort"`
					} `json:"ports"`
					Resources struct {
						Limits struct {
							Cpu    string `json:"cpu"`
							Memory string `json:"memory"`
						} `json:"limits"`
					} `json:"resources"`
				} `json:"containers"`
			} `json:"spec"`
		} `json:"template"`
	} `json:"spec"`
}

type DeleteDmOptions struct {
	APIVersion         string `json:"apiVersion"`
	GracePeriodSeconds int    `json:"gracePeriodSeconds"`
	PropagationPolicy  string `json:"propagationPolicy"`
	Preconditions      struct {
		UID string `json:"uid"`
	} `json:"preconditions"`
}

type DmResponse struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		CreationTimestamp string `json:"creationTimestamp"`
		Generation        int    `json:"generation"`
		Labels            struct {
			App string `json:"app"`
		} `json:"labels"`
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		ResourceVersion string `json:"resourceVersion"`
		SelfLink        string `json:"selfLink"`
		UID             string `json:"uid"`
		Annotations     string `json:"annotations"`
	} `json:"metadata"`
	Spec struct {
		Replicas int `json:"replicas"`
		Selector struct {
			MatchLabels struct {
				App string `json:"app"`
			} `json:"matchLabels"`
		} `json:"selector"`
		Template struct {
			Metadata struct {
				CreationTimestamp interface{} `json:"creationTimestamp"`
				Labels            struct {
					App string `json:"app"`
				} `json:"labels"`
			} `json:"metadata"`
			Spec struct {
				Containers []struct {
					Image           string `json:"image"`
					ImagePullPolicy string `json:"imagePullPolicy"`
					Name            string `json:"name"`
					Ports           []struct {
						ContainerPort int    `json:"containerPort"`
						Protocol      string `json:"protocol"`
					} `json:"ports"`
					Resources                struct{} `json:"resources"`
					TerminationMessagePath   string   `json:"terminationMessagePath"`
					TerminationMessagePolicy string   `json:"terminationMessagePolicy"`
				} `json:"containers"`
				DNSPolicy                     string   `json:"dnsPolicy"`
				RestartPolicy                 string   `json:"restartPolicy"`
				SchedulerName                 string   `json:"schedulerName"`
				SecurityContext               struct{} `json:"securityContext"`
				TerminationGracePeriodSeconds int      `json:"terminationGracePeriodSeconds"`
			} `json:"spec"`
			Strategy struct {
				RollingUpdate struct {
					MaxSurge       int `json"maxSurge`
					MaxUnavailable int `json:"maxUnavailable"`
				} `json:"rollingUpdate "`
				Type string `json:"type"`
			} `json:"strategy "`
		} `json:"template"`
	} `json:"spec"`
	Status struct {
		Replicas           int `json:"replicas"`
		ObservedGeneration int `json:"observedGeneration"`
		UpdatedReplicas    int `json:"updatedReplicas"`
		AvailableReplicas  int `json:"availableReplicas"`
	} `json:"status"`
}

type DeploymentStatus struct {
	APIVersion string      `json:"apiVersion"`
	Code       int         `json:"code"`
	Kind       string      `json:"kind"`
	Metadata   struct{}    `json:"metadata"`
	Status     interface{} `json:"status"`
}

//ingress
type CreateIngStruct struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Spec struct {
		Rules []struct {
			Host string `json:"host"`
			HTTP struct {
				Paths []struct {
					Backend struct {
						ServiceName string `json:"serviceName"`
						ServicePort int    `json:"servicePort"`
					} `json:"backend"`
					Path string `json:"path"`
				} `json:"paths"`
			} `json:"http"`
		} `json:"rules"`
	} `json:"spec"`
}

type GetIngList struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		Metadata struct {
			CreationTimestamp string `json:"creationTimestamp"`
			Generation        int    `json:"generation"`
			Name              string `json:"name"`
			Namespace         string `json:"namespace"`
			ResourceVersion   string `json:"resourceVersion"`
			SelfLink          string `json:"selfLink"`
			UID               string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			Rules []struct {
				Host string `json:"host"`
				HTTP struct {
					Paths []struct {
						Backend struct {
							ServiceName string `json:"serviceName"`
							ServicePort int    `json:"servicePort"`
						} `json:"backend"`
						Path string `json:"path"`
					} `json:"paths"`
				} `json:"http"`
			} `json:"rules"`
		} `json:"spec"`
		Status struct {
			LoadBalancer struct {
				Ingress []struct {
					IP string `json:"ip"`
				} `json:"ingress"`
			} `json:"loadBalancer"`
		} `json:"status"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
		SelfLink        string `json:"selfLink"`
	} `json:"metadata"`
}

type DeleteIngSuccess struct {
	APIVersion string   `json:"apiVersion"`
	Code       int      `json:"code"`
	Kind       string   `json:"kind"`
	Metadata   struct{} `json:"metadata"`
	Status     string   `json:"status"`
}

type CreateIngSuccess struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		CreationTimestamp string `json:"creationTimestamp"`
		Generation        int    `json:"generation"`
		Name              string `json:"name"`
		Namespace         string `json:"namespace"`
		ResourceVersion   string `json:"resourceVersion"`
		SelfLink          string `json:"selfLink"`
		UID               string `json:"uid"`
	} `json:"metadata"`
	Spec struct {
		Rules []struct {
			Host string `json:"host"`
			HTTP struct {
				Paths []struct {
					Backend struct {
						ServiceName string `json:"serviceName"`
						ServicePort int    `json:"servicePort"`
					} `json:"backend"`
					Path string `json:"path"`
				} `json:"paths"`
			} `json:"http"`
		} `json:"rules"`
	} `json:"spec"`
	Status struct {
		LoadBalancer struct{} `json:"loadBalancer"`
	} `json:"status"`
}

//pvc
type Pvc struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		CreationTimestamp string `json:"creationTimestamp"`
		Generation        int    `json:"generation"`
		Name              string `json:"name"`
		Namespace         string `json:"namespace"`
		ResourceVersion   string `json:"resourceVersion"`
		SelfLink          string `json:"selfLink"`
		UID               string `json:"uid"`
	} `json:"metadata"`
	Spec struct {
		AccessModes []string `json:"accessModes"`
		Resources   struct {
			Requests struct {
				Storage string `json:"storage"`
			} `json:"requests"`
		} `json:"resources"`
	} `json:"spec"`
}

//rc
type CreateRcStruct struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Spec struct {
		Replicas int `json:"replicas"`
		Selector struct {
			App string `json:"app"`
		} `json:"selector"`
		Template struct {
			Metadata struct {
				Labels struct {
					App string `json:"app"`
				} `json:"labels"`
			} `json:"metadata"`
			Spec struct {
				Volumes []struct {
					Name      string `json:"name"`
					ConfigMap struct {
						Name string `json:"name"`
					} `json:"configMap"`
				} `json:"volumes"`
				RestartPolicy string `json:"restartPolicy"`
				Containers    []struct {
					Image        string `json:"image"`
					Name         string `json:"name"`
					VolumeMounts []struct {
						Name      string `json:"name"`
						MountPath string `json:"mountPath"`
					} `json:"volumeMounts"`
					Ports []struct {
						ContainerPort int `json:"containerPort"`
					} `json:"ports"`
					Resources struct {
						Limits struct {
							Cpu    string `json:"cpu"`
							Memory string `json:"memory"`
						} `json:"limits"`
					} `json:"resources"`
				} `json:"containers"`
			} `json:"spec"`
		} `json:"template"`
	} `json:"spec"`
}

type CreateRcSuccess struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		CreationTimestamp string `json:"creationTimestamp"`
		Generation        int    `json:"generation"`
		Labels            struct {
			App string `json:"app"`
		} `json:"labels"`
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		ResourceVersion string `json:"resourceVersion"`
		SelfLink        string `json:"selfLink"`
		UID             string `json:"uid"`
	} `json:"metadata"`
	Spec struct {
		Replicas int `json:"replicas"`
		Selector struct {
			App string `json:"app"`
		} `json:"selector"`
		Template struct {
			Metadata struct {
				CreationTimestamp interface{} `json:"creationTimestamp"`
				Labels            struct {
					App string `json:"app"`
				} `json:"labels"`
			} `json:"metadata"`
			Spec struct {
				Containers []struct {
					Image           string `json:"image"`
					ImagePullPolicy string `json:"imagePullPolicy"`
					Name            string `json:"name"`
					Ports           []struct {
						ContainerPort int    `json:"containerPort"`
						Protocol      string `json:"protocol"`
					} `json:"ports"`
					Resources                struct{} `json:"resources"`
					TerminationMessagePath   string   `json:"terminationMessagePath"`
					TerminationMessagePolicy string   `json:"terminationMessagePolicy"`
				} `json:"containers"`
				DNSPolicy                     string   `json:"dnsPolicy"`
				RestartPolicy                 string   `json:"restartPolicy"`
				SchedulerName                 string   `json:"schedulerName"`
				SecurityContext               struct{} `json:"securityContext"`
				TerminationGracePeriodSeconds int      `json:"terminationGracePeriodSeconds"`
			} `json:"spec"`
		} `json:"template"`
	} `json:"spec"`
	Status struct {
		Replicas int `json:"replicas"`
	} `json:"status"`
}

type SvcCreateStruct struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Spec struct {
		Ports []struct {
			Name       string `json:"name"`
			Port       int    `json:"port"`
			Protocol   string `json:"protocol"`
			TargetPort int    `json:"targetPort"`
		} `json:"ports"`
		Selector struct {
			App string `json:"app"`
		} `json:"selector"`
	} `json:"spec"`
}

type DeleteRc struct {
	APIVersion         string `json:"apiVersion"`
	GracePeriodSeconds int    `json:"gracePeriodSeconds"`
	OrphanDependents   bool   `json:"orphanDependents"`
	Preconditions      struct {
		UID string `json:"uid"`
	} `json:"preconditions"`
}

type DeleteRcSuccess struct {
	APIVersion string   `json:"apiVersion"`
	Code       int      `json:"code"`
	Kind       string   `json:"kind"`
	Metadata   struct{} `json:"metadata"`
	Status     string   `json:"status"`
}

type IngCreateStruct struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Spec struct {
		Rules []struct {
			Host string `json:"host"`
			HTTP struct {
				Paths []struct {
					Backend struct {
						ServiceName string `json:"serviceName"`
						ServicePort int    `json:"servicePort"`
					} `json:"backend"`
					Path string `json:"path"`
				} `json:"paths"`
			} `json:"http"`
		} `json:"rules"`
	} `json:"spec"`
}

//svc
type CreateSvcStruct struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Spec struct {
		Ports []struct {
			Name       string `json:"name"`
			Port       int    `json:"port"`
			Protocol   string `json:"protocol"`
			TargetPort int    `json:"targetPort"`
		} `json:"ports"`
		Selector struct {
			App string `json:"app"`
		} `json:"selector"`
	} `json:"spec"`
}

type GetSvcList struct {
	APIVersion string `json:"apiVersion"`
	Items      []struct {
		Metadata struct {
			CreationTimestamp string `json:"creationTimestamp"`
			Name              string `json:"name"`
			Namespace         string `json:"namespace"`
			ResourceVersion   string `json:"resourceVersion"`
			SelfLink          string `json:"selfLink"`
			UID               string `json:"uid"`
		} `json:"metadata"`
		Spec struct {
			ClusterIP string `json:"clusterIP"`
			Ports     []struct {
				Name       string `json:"name"`
				Port       int    `json:"port"`
				Protocol   string `json:"protocol"`
				TargetPort int    `json:"targetPort"`
			} `json:"ports"`
			Selector struct {
				App string `json:"app"`
			} `json:"selector"`
			SessionAffinity string `json:"sessionAffinity"`
			Type            string `json:"type"`
		} `json:"spec"`
		Status struct {
			LoadBalancer struct{} `json:"loadBalancer"`
		} `json:"status"`
	} `json:"items"`
	Kind     string `json:"kind"`
	Metadata struct {
		ResourceVersion string `json:"resourceVersion"`
		SelfLink        string `json:"selfLink"`
	} `json:"metadata"`
}

type DeleteSvc struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type DeleteSvcSuccess struct {
	APIVersion string   `json:"apiVersion"`
	Code       int      `json:"code"`
	Kind       string   `json:"kind"`
	Metadata   struct{} `json:"metadata"`
	Status     string   `json:"status"`
}

type CreateSvcSuccess struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		CreationTimestamp string `json:"creationTimestamp"`
		Name              string `json:"name"`
		Namespace         string `json:"namespace"`
		ResourceVersion   string `json:"resourceVersion"`
		SelfLink          string `json:"selfLink"`
		UID               string `json:"uid"`
	} `json:"metadata"`
	Spec struct {
		ClusterIP string `json:"clusterIP"`
		Ports     []struct {
			Name       string `json:"name"`
			Port       int    `json:"port"`
			Protocol   string `json:"protocol"`
			TargetPort int    `json:"targetPort"`
		} `json:"ports"`
		Selector struct {
			App string `json:"app"`
		} `json:"selector"`
		SessionAffinity string `json:"sessionAffinity"`
		Type            string `json:"type"`
	} `json:"spec"`
	Status struct {
		LoadBalancer struct{} `json:"loadBalancer"`
	} `json:"status"`
}
