package k8s

import (
	"FACP/third/client"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type rc struct {
	c client.Client
}

func rcIns(k8sUrl string) K8s {
	return &rc{
		c: client.ClientInstance(k8sUrl, "", ""),
	}
}

//删除已存在的rc  "/api/v1/namespaces/" + namespace + "/replicationcontrollers/" + name
func (this *rc) Offline(name, namespace string) (err error) {

	var res DeleteRcSuccess

	path := fmt.Sprintf(rcOpt, namespace, name)

	isExist, result, _ := this.Single(name, namespace)

	if isExist {

		r := result.(CreateRcSuccess)

		Rc := DeleteRc{
			APIVersion:         "v1",
			GracePeriodSeconds: 0,
			OrphanDependents:   false,
		}
		Rc.Preconditions.UID = r.Metadata.UID

		RCdel, _ := json.Marshal(Rc)

		_, err = this.c.Delete(path, jsonHeader, RCdel, &res)
		if res.Code == 200 {
			return nil
		}
	}

	return err
}

//创建rc  "/api/v1/namespaces/" + namespace + "/replicationcontrollers"
func (this *rc) Upline(nameSpace string, body []byte) (err error) {

	var result CreateRcSuccess

	path := fmt.Sprintf(rcNew, nameSpace)

	status, err := this.c.Post(path, jsonHeader, body, &result)

	if status == 200 {
		return nil
	}
	return err

}

// GET /api/v1/namespaces/{namespace}/replicationcontrollers/{name}.
func (this *rc) Single(rcName, nameSpace string) (isExist bool, result interface{}, err error) {

	path := fmt.Sprintf(rcOpt, nameSpace, rcName)

	status, err := this.c.Get(path, jsonHeader, nil, &result)

	if status == 200 {
		return true, result, nil
	}
	return false, result, err

}

func (this *rc) ConfirmDetele(ctx context.Context, rcName, nameSpace string) (isDelete bool) {

	hasDelete := make(chan bool, 1)

	go func() {
		for {
			time.Sleep(2 * time.Second)

			if deadline, ok := ctx.Deadline(); ok { //设置了deadl
				if time.Now().After(deadline) {
					break
				}
			}
			if isExist, _, _ := this.Single(rcName, nameSpace); !isExist {
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

	case <-hasDelete: // 已删除

		isDelete = true
		close(hasDelete)
		return
	}

	return isDelete
}

func (this *rc) Update(name, nameSpace string, body []byte) (err error) {

	return
}

func (this *rc) Scale(name, nameSpace string, replicas int) (err error) {

	return
}

func (this *rc) Deploy(fileUrl string) error {
	return nil
}

func (this *rc) IsExist(name, NameSpace string) (isExist bool, err error) {

	return
}
