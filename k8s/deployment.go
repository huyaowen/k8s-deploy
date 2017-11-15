package k8s

import (
	"context"
	"encoding/json"
	"fmt"
	"k8s-deploy/client"
	"log"
	"time"
)

type deployment struct {
	c client.Client
}

func deploymenyIns(k8sUrl string) K8s {
	return &deployment{
		c: client.ClientInstance(k8sUrl, "", ""),
	}
}

func (this *deployment) Upline(namesapce string, body []byte) (err error) {

	path := fmt.Sprintf(deploymentNew, namesapce)

	status, err := this.c.Post(path, jsonHeader, body, nil)

	if status != 200 {
		return err
	}
	return nil
}

func (this *deployment) Offline(deployment, namesapce string) (err error) {

	isExist, result, err := this.Single(namesapce, deployment)
	if err != nil {
		return err
	}
	if isExist {
		res := result.(DmResponse)
		option := DeleteDmOptions{
			APIVersion:         "v1",
			GracePeriodSeconds: 0,
			PropagationPolicy:  "Foreground",
		}
		option.Preconditions.UID = res.Metadata.UID

		body, _ := json.Marshal(option)

		path := fmt.Sprintf(deplotmentOpt, namesapce, deployment)

		status, err := this.c.Delete(path, jsonHeader, body, nil)

		if status != 200 {

			return err
		}
	}

	return nil
}

func (this *deployment) Update(deployment, namesapce string, body []byte) (err error) {

	path := fmt.Sprintf(deplotmentOpt, namesapce, deployment)

	status, err := this.c.Patch(path, patchHeader, body, nil)

	if status == 200 {
		return nil
	}
	return err
}

func (this *deployment) Single(deployment, namesapce string) (isExist bool, result interface{}, err error) {

	path := fmt.Sprintf(deplotmentOpt, namesapce, deployment)

	status, err := this.c.Get(path, jsonHeader, nil, &result)

	if status == 200 {
		return true, result, nil
	} else if status == 404 {
		return false, result, nil
	} else {
		return false, result, err
	}
	return false, result, nil
}

func (this *deployment) ConfirmDeteleDeploy(ctx context.Context, namespace, deployName string) (isDelete bool) {

	hasDelete := make(chan bool, 1)

	go func() {
		for {
			time.Sleep(2 * time.Second)

			if deadline, ok := ctx.Deadline(); ok { 
				if time.Now().After(deadline) {
					break
				}
			}
			if isExist, _, _ := this.Single(namespace, deployName); !isExist {
				hasDelete <- true
				break
			}
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("delete rc timeout：", ctx.Err())
		isDelete = false
		return

	case <-hasDelete:

		isDelete = true
		close(hasDelete)
		return
	}

	return isDelete
}

func (this *deployment) Scale(name, nameSpace string, replicas int) (err error) {

	return
}

func (this *deployment) Deploy(fileUrl string) error {
	var (
		dm      Deployment
		isExist bool
		err     error
		body    []byte
	)

	if body, err = getRequest(fileUrl); err != nil {

		return err
	}

	json.Unmarshal(body, &dm)

	name := dm.Metadata.Name
	namespace := dm.Metadata.Namespace

	if isExist, _, err = this.Single(name, namespace); err != nil {

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

func (this *deployment) IsExist(name, NameSpace string) (isExist bool, err error) {

	return
}
