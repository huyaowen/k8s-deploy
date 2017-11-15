package k8s

import (
	"FACP/third/client"
	"encoding/json"
	"errors"
	"fmt"
)

type svc struct {
	c client.Client
}

func svcIns(k8sUrl string) K8s {
	return &svc{
		c: client.ClientInstance(k8sUrl, "", ""),
	}
}

//删除已存在的svc  "/api/v1/namespaces/" + namespace + "/services/" + name
func (this *svc) Offline(name, namespace string) (err error) {

	var res DeleteSvcSuccess

	path := fmt.Sprintf(svcOpt, namespace, name)

	if isExist, _, _ := this.Single(name, namespace); isExist {

		Svc := DeleteSvc{
			Name:      name,
			Namespace: namespace,
		}

		SVCdel, _ := json.Marshal(Svc)

		_, err = this.c.Delete(path, jsonHeader, SVCdel, &res)
		if res.Code == 200 {
			return nil
		}
	}

	return err
}

//创建svc "/api/v1/namespaces/" + namespace + "/services"
func (this *svc) Upline(nameSpace string, body []byte) (err error) {

	var result CreateSvcSuccess

	path := fmt.Sprintf(svcNew, nameSpace)

	status, err := this.c.Post(path, jsonHeader, body, &result)
	if status == 200 {
		return nil
	}
	return err
}

// GET  "/api/v1/namespaces/" + namespace + "/services/" + name
func (this *svc) Single(name, nameSpace string) (isExist bool, result interface{}, err error) {

	path := fmt.Sprintf(svcOpt, nameSpace, name)

	status, err := this.c.Get(path, jsonHeader, nil, &result)

	if status == 200 {
		return true, result, nil
	} else if status == 404 {
		return false, result, nil
	}
	return false, result, err
}

func (this *svc) Update(name, nameSpace string, body []byte) (err error) {

	return
}

func (this *svc) Scale(name, nameSpace string, replicas int) (err error) {

	return
}

func (this *svc) Deploy(fileUrl string) error {

	var svc SvcCreateStruct

	body, err := getRequest(fileUrl)
	if err != nil {
		return errors.New("svc配置文件获取失败:" + err.Error())
	}

	json.Unmarshal(body, &svc)
	//获取svc配置文件的name,namespace

	name, namespace := svc.Metadata.Name, svc.Metadata.Namespace

	isExsit, _, err := this.Single(name, namespace)
	if err != nil {

		return err
	}

	if isExsit {
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

func (this *svc) IsExist(name, NameSpace string) (isExist bool, err error) {

	return
}
