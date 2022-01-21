package main

import (
	"io/ioutil"
	"os"

	"github.com/golang/glog"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/kube"
)

type KubeInformation struct {
	AimNamespace string
	AimContext   string
	endpoint     string
	token        string
	kubeUserName string
	kubePassword string
}

func InitKubeInformation(namespace, context string) *KubeInformation {
	return &KubeInformation{
		AimNamespace: namespace,
		AimContext:   context,
	}
}

func InitKubeAllInformation(namespace, context, endpoint, token, kubeUserName, kubePassword string) *KubeInformation {
	return &KubeInformation{
		AimNamespace: namespace,
		AimContext:   context,
		endpoint:     endpoint,
		token:        token,
		kubeUserName: kubeUserName,
		kubePassword: kubePassword,
	}
}

func actionConfigInit(kubeInfo *KubeInformation) (*action.Configuration, error) {
	actionConfig := new(action.Configuration)
	if kubeInfo.AimContext == "" {
		kubeInfo.AimContext = settings.KubeContext
	}
	clientConfig := kube.GetConfig(settings.KubeConfig, kubeInfo.AimContext, kubeInfo.AimNamespace)
	if settings.KubeToken != "" {
		clientConfig.BearerToken = &settings.KubeToken
	}
	if settings.KubeAPIServer != "" {
		clientConfig.APIServer = &settings.KubeAPIServer
	}

	//若k8s endpoint token 都为空, 则默认default集群，否则cce集群
	if kubeInfo.endpoint == "" && kubeInfo.token == "" {
		endpoint := "https://kubernetes.default.svc"
		clientConfig.APIServer = &endpoint
		bytes, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
		if err != nil {
			glog.Fatal(err)
		}
		token := string(bytes)
		clientConfig.BearerToken = &token
	}

	if kubeInfo.endpoint != "" {
		clientConfig.APIServer = &kubeInfo.endpoint
	}
	if kubeInfo.token != "" {
		clientConfig.BearerToken = &kubeInfo.token
	}
	//k8s user password 认证
	if kubeInfo.kubeUserName != "" && kubeInfo.kubePassword != "" {
		glog.Info("use user password auth for APIServer :{}", clientConfig.APIServer)
		clientConfig.Username = &kubeInfo.kubeUserName
		clientConfig.Password = &kubeInfo.kubePassword
	}

	insecure := true
	clientConfig.Insecure = &insecure
	err := actionConfig.Init(clientConfig, kubeInfo.AimNamespace, os.Getenv("HELM_DRIVER"), glog.Infof)
	if err != nil {
		glog.Errorf("%+v", err)
		return nil, err
	}

	return actionConfig, nil
}
