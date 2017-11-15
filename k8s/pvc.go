package k8s

import (
	"FACP/arch/utils"
	"FACP/third/client"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type pvc struct {
	c client.Client
}

func pvcIns(k8sUrl string) K8s {
	return &pvc{
		c: client.ClientInstance(k8sUrl, "", ""),
	}
}

func (this *pvc) Upline(namesapce string, body []byte) (err error) {

	path := fmt.Sprintf(pvcNew, namesapce)

	status, err := this.c.Post(path, jsonHeader, body, nil)

	if status != 200 {
		return err
	}
	return nil
}

func (this *pvc) Update(pvcName, namesapce string, body []byte) (err error) {

	var hasDelete bool

	isExist, _, err := this.Single(namesapce, pvcName)

	if err != nil {
		return err
	}

	if isExist {

		if err = this.Offline(namesapce, pvcName); err != nil {

		}

		ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(15*time.Second))
		if hasDelete = this.ConfirmDetelePvc(ctx, namesapce, pvcName); hasDelete {
			log.Println("delete pvc:" + pvcName + ",namespace:" + namesapce + "succuess!")
		}
	}

	if err = this.Upline(namesapce, body); err != nil {
		return err
	}

	return nil
}

func (this *pvc) Offline(namesapce, pvcName string) (err error) {

	isExist, result, err := this.Single(namesapce, pvcName)
	if err != nil {
		return err
	}
	if isExist {

		p := result.(Pvc)

		option := DeleteDmOptions{
			APIVersion:         "v1",
			GracePeriodSeconds: 0,
			PropagationPolicy:  "Foreground",
		}
		option.Preconditions.UID = p.Metadata.UID

		body, _ := json.Marshal(option)

		path := fmt.Sprintf(pvcOpt, namesapce, pvcName)

		status, err := this.c.Delete(path, jsonHeader, body, nil)

		if status != 200 {

			return err
		}
	}

	return nil
}

func (this *pvc) Single(namesapce, pvcName string) (isExist bool, result interface{}, err error) {

	path := fmt.Sprintf(pvcOpt, namesapce, pvcName)

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

func (this *pvc) ConfirmDetelePvc(ctx context.Context, namespace, pvcName string) (isDelete bool) {

	hasDelete := make(chan bool, 1)

	go func() {
		for {
			time.Sleep(2 * time.Second)

			if deadline, ok := ctx.Deadline(); ok { //设置了deadl
				if time.Now().After(deadline) {
					break
				}
			}
			if isExist, _, _ := this.Single(namespace, pvcName); !isExist {
				hasDelete <- true
				break
			}
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("删除rc超时：", ctx.Err())
		isDelete = false
		return

	case <-hasDelete: // 已删除

		isDelete = true
		close(hasDelete)
		return
	}

	return isDelete
}

func (this *pvc) Scale(name, nameSpace string, replicas int) (err error) {

	return
}

func (this *pvc) Deploy(filePath string) error {

	var body []byte
	var err error

	os.Chdir(filePath)
	defer os.Chdir(filePath)

	//此处url为文件路径
	body = utils.ReadBodyFromFile("pvc.json", filePath)

	var (
		dm      Pvc
		isExist bool
	)
	if err = json.Unmarshal(body, &dm); err != nil {

		return err
	}

	name := dm.Metadata.Name
	namespace := dm.Metadata.Namespace

	if isExist, _, err = this.Single(namespace, name); err != nil {

		return err
	}

	if isExist {

		if err = this.Update(namespace, name, body); err != nil {

			return err
		}

	} else {

		if err = this.Upline(namespace, body); err != nil {

			return err
		}

	}

	return nil

}

func (this *pvc) IsExist(name, NameSpace string) (isExist bool, err error) {

	return
}
