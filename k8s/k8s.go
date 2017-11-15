package k8s

import (
	"io/ioutil"
	"net/http"
)

const (
	NAMESPACE               = "namespace"
	DEPLOYMENT              = "deploment"
	INGRESS                 = "ingress"
	SERVICE                 = "service"
	PERSISTENT_VOLUME_CLAIM = "pvc"
	CONFIGMAP               = "configmap"
)

var k8sUrl string

type K8s interface {
	Single(name, namespace string) (isExist bool, result interface{}, err error)

	Update(name, nameSpace string, body []byte) (err error)

	Upline(nameSpace string, body []byte) (err error)

	Offline(name, nameSpace string) (err error)

	Scale(name, nameSpace string, replicas int) (err error)

	Deploy(fileUrl string) error

	IsExist(name, nameSapce string) (isExist bool, err error)
}

func K8sFactory(sourceType string) (k8s K8s) {

	//TODO  get k8s url from another place
	k8sUrl := "k8s_url"

	switch sourceType {

	case DEPLOYMENT:

		return deploymenyIns(k8sUrl)

	case SERVICE:

		return svcIns(k8sUrl)

	case INGRESS:

		return ingressIns(k8sUrl)

	case PERSISTENT_VOLUME_CLAIM:

		return pvcIns(k8sUrl)

	case CONFIGMAP:

		return configmapIns(k8sUrl)

	case NAMESPACE:

		return namespaceIns(k8sUrl)

	default:
	}

	return
}

func getRequest(url string) (response []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
