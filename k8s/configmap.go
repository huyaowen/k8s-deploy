package k8s

import (
	"errors"
	"fmt"
	"k8s-deploy/client"
)

type configmap struct {
	c client.Client
}

func configmapIns(k8sUrl string) K8s {
	return &configmap{
		c: client.ClientInstance(k8sUrl, "", ""),
	}
}

func (this *configmap) Upline(nameSpace string, body []byte) error {

	_, err := this.c.Post(fmt.Sprintf(configMapNew, nameSpace), jsonHeader, body, nil)
	if err != nil {
		return err
	}
	return nil
}

func (this *configmap) Update(configName, nameSpace string, body []byte) error {

	_, err := this.c.Put(fmt.Sprintf(configMapOpt, nameSpace, configName), jsonHeader, body, nil)
	if err != nil {
		return err
	}
	return nil
}

/*
*	删除需传入url namespace 和configName
 */
func (this *configmap) Offline(configName, nameSpace string) (err error) {
	isExist, err := this.IsExist(configName, nameSpace)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New("该配置不存在")
	}
	_, err = this.c.Delete(fmt.Sprintf(configMapOpt, nameSpace, configName), jsonHeader, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

func (this *configmap) Single(configName, namespace string) (isExist bool, result interface{}, err error) {
	var config interface{}
	status, err := this.c.Get(fmt.Sprintf(configMapOpt, namespace, configName), nil, nil, &config)
	if status == 200 {
		return true, config, nil
	} else if status == 404 {
		return false, nil, nil
	} else {
		return false, nil, err
	}
	return
}

func (this *configmap) IsExist(ConfigName, NameSpace string) (isExist bool, err error) {
	isExist, _, err = this.Single(ConfigName, NameSpace)
	return isExist, err
}

func (this *configmap) Scale(name, nameSpace string, replicas int) (err error) {

	return
}

func (this *configmap) Deploy(fileUrl string) error {
	return nil
}
