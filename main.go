package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	tektonServer string
	namespace    string
	name         string
	paramImg     string
)

func init() {
	//读取yaml文件
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName("application")
	//添加读取的配置文件路径
	v.AddConfigPath("./config/")
	//windows环境下为%GOPATH，linux环境下为$GOPATH
	v.AddConfigPath("$GOPATH/src/")
	//设置配置文件类型
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}

	tektonServer = v.GetString("TEKTON_SERVER")
	namespace = v.GetString("PIPELINERUN.NAMESPACE")
	name = v.GetString("PIPELINERUN.NAME")
	paramImg = v.GetString("PIPELINERUN.PARAMIMG")

	pipelineRunParamImg := os.Getenv("PIPELINERUN_PARAMIMG")
	if pipelineRunParamImg != "" {
		paramImg = pipelineRunParamImg
	}

	pipelineRunNamespace := os.Getenv("PIPELINERUN_NAMESPACE")
	if pipelineRunNamespace != "" {
		namespace = pipelineRunNamespace
	}

	pipelineRunName := os.Getenv("PIPELINERUN_NAME")
	if pipelineRunName != "" {
		namespace = pipelineRunName
	}

	log.Println("pipelineRunNamespace: ", namespace)
	log.Println("pipelineRunName: ", name)
	// log.Println("e2e test image: ", paramImg)
}

func main() {
	pipelineRun := GetPipelineRun(namespace, name)
	/*params := pipelineRun.Spec.Params
	for i := 0; i < len(params); i++ {
		if params[i].Name == "E2E_STRESS_IMAGE" {
			params[i].Value = v1beta1.ArrayOrString{Type: "string", StringVal: paramImg}
		}
	}*/
	pipelineRun.Name = name
	//pipelineRun.Spec.Params = params
	pipelineRun.ResourceVersion = ""

	// obj convert json []byte
	pipelineRunJson, _ := json.Marshal(pipelineRun)
	// log.Println(string(pipelineRunJson))

	log.Println("Delete PipelineRun: ", name)
	DeletePipelineRun(namespace, name)
	log.Println("Create PipelineRun: ", name)
	createPipelineRun(namespace, pipelineRunJson)

	// https://xieys.club/go-client-conversion/

	//yml, _ := yaml.JSONToYAML(resJson)
	//ymlName := fmt.Sprintf("%s.yml", pipelineRun.Name)
	//f, _ := os.Create(ymlName)
	//defer f.Close()
	//_, _ = f.WriteString(string(yml))
}
