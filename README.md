# k8s-deploy

封装 k8s rest api对namespace ingress configmap deployment pvc等部署

# 使用方法

configCli := k8s.K8sFactory(k8s.CONFIGMAP)

configCli.Update(name,namesapce,body)
