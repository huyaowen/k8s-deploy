package k8s

import (
	"FACP/third/client"
	"errors"
	"fmt"
)

type namespace struct {
	c client.Client
}

func namespaceIns(k8sUrl string) K8s {
	return &namespace{
		c: client.ClientInstance(k8sUrl, "", ""),
	}
}

func (this *namespace) Upline(nameSpace string, body []byte) (err error) {

	status, err := this.c.Post(namespaceNew, jsonHeader, body, nil)
	if status == 201 && err == nil {
		return nil
	}
	if err != nil {

	}
	return errors.New("create namespace fail")
}

func (this *namespace) Update(name, nameSpace string, body []byte) (err error) {

	status, err := this.c.Patch(fmt.Sprintf(namespaceOpt, nameSpace), jsonHeader, body, nil)
	if status == 200 && err == nil {
		return nil
	}
	return errors.New("update namespace fail")
}

func (this *namespace) Offline(name, nameSpace string) (err error) {

	var (
		status  int
		isExist bool
	)

	isExist, err = this.IsExist("", nameSpace)
	if err != nil {
		return err
	}
	if isExist {
		status, err = this.c.Delete(fmt.Sprintf(namespaceOpt, nameSpace), jsonHeader, nil, nil)
		if status == 200 && err == nil {
			return nil
		}
	}

	return err
}

func (this *namespace) Single(name, nameSpace string) (isExist bool, result interface{}, err error) {
	var namespace interface{}
	status, err := this.c.Get(fmt.Sprintf(namespaceOpt, nameSpace), nil, nil, &namespace)
	if status == 200 {
		return true, namespace, nil
	} else if status == 404 {
		return false, nil, nil
	} else {
		return false, nil, err
	}
	return false, nil, nil
}

func (this *namespace) IsExist(name, NameSpace string) (isExsit bool, err error) {
	isExist, _, err := this.Single("", NameSpace)
	return isExist, err

}

func (this *namespace) Scale(name, nameSpace string, replicas int) (err error) {

	return
}

func (this *namespace) Deploy(fileUrl string) error {
	return nil
}
