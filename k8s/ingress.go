package k8s

import (
	"FACP/third/client"
	"encoding/json"
	"errors"
	"fmt"
)

type ingress struct {
	c client.Client
}

func ingressIns(k8sUrl string) K8s {
	return &ingress{
		c: client.ClientInstance(k8sUrl, "", ""),
	}
}

func (this *ingress) Offline(name, nameSpace string) (err error) {

	var (
		res     interface{}
		status  int
		isExist bool
	)

	path := fmt.Sprintf(ingOpt, nameSpace, name)

	if isExist, _, _ = this.Single(name, nameSpace); isExist {

		status, err = this.c.Delete(path, jsonHeader, nil, &res)
		if status == 200 {
			return nil
		}
	}

	return err
}

func (this *ingress) Upline(nameSpace string, ingbyte []byte) (err error) {

	path := fmt.Sprintf(ingNew, nameSpace)

	var result interface{}

	status, err := this.c.Post(path, jsonHeader, ingbyte, &result)

	if status == 200 {
		return nil
	}
	return err

}

func (this *ingress) Single(name, nameSpace string) (isExist bool, result interface{}, err error) {

	path := fmt.Sprintf(ingOpt, nameSpace, name)

	status, err := this.c.Get(path, jsonHeader, nil, &result)

	if status == 200 {
		return true, result, nil
	} else if status == 404 {
		return false, result, nil
	}
	return false, result, err
}

func (this *ingress) Scale(name, nameSpace string, replicas int) (err error) {

	return
}

func (this *ingress) Update(name, nameSpace string, body []byte) (err error) {

	return
}

func (this *ingress) Deploy(fileUrl string) error {
	
	var model CreateIngStruct

	body, err := getRequest(fileUrl)
	if err != nil {
		return errors.New("get ing file fail:" + err.Error())
	}

	json.Unmarshal(body, &model)

	
	name, namespace := model.Metadata.Name, model.Metadata.Namespace

	isExist, _, err := this.Single(name, namespace)

	if err != nil {

		return err
	}

	if isExist {

		if err = this.Update(name, namespace, body); err != nil {

			return err
		}

	} else {

		if err = this.Upline(namespace, body); err != nil {

			return err
		}
	}

	return nil

}

func (this *ingress) IsExist(name, NameSpace string) (isExist bool, err error) {

	return
}
