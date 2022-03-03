package main

import (
	"bytes"
	"encoding/json"
	http2 "github.com/cloudevents/sdk-go/v2/protocol/http"
	. "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime"
	"log"
	"net/http"
)

var (
	apiPre          = "/apis/tekton.dev/v1beta1"
	namespacePre    = "/namespaces/"
	pipelineRunsPre = "/pipelineruns/"
	authKey         = "Tekton-Client"
	authValue       = "tektoncd/dashboard"
)

// GetPipelineRun get pipelineRun resource by params namespace and name
func GetPipelineRun(namespace, name string) *PipelineRun {
	url := tektonServer + apiPre + namespacePre + namespace + pipelineRunsPre + name
	response, err := http.Get(url)
	if err != nil {
		return nil
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}

	pipelineRun := &PipelineRun{}
	// json str convert obj
	json.Unmarshal(body, pipelineRun)

	return pipelineRun
}

// createPipelineRun create pipelineRun resource by namespace and pipelineRun definition
func createPipelineRun(namespace string, pipelineRun []byte) {
	url := tektonServer + apiPre + namespacePre + namespace + pipelineRunsPre + name
	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(pipelineRun))
	request.Header.Add(authKey, authValue)
	request.Header.Add(http2.ContentType, runtime.ContentTypeJSON)

	res, _ := http.DefaultClient.Do(request)
	defer res.Body.Close()
	ioutil.ReadAll(res.Body)
	log.Println(res.Status)
	// log.Println(string(body))
}

// DeletePipelineRun delete pipelineRun resource by params namespace and name
func DeletePipelineRun(namespace, name string) {
	url := tektonServer + apiPre + namespacePre + namespace + pipelineRunsPre + name
	request, _ := http.NewRequest(http.MethodDelete, url, nil)
	request.Header.Add(authKey, authValue)
	res, _ := http.DefaultClient.Do(request)
	defer res.Body.Close()
	ioutil.ReadAll(res.Body)

	log.Println(res.Status)
	// log.Println(string(body))
}
