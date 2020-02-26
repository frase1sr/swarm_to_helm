/*
Copyright 2017 The Kubernetes Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package app

import (
	//"github.com/frase1sr/swarm_to_helm/cmd"
	//"github.com/spf13/cobra"
	"github.com/frase1sr/swarm_to_helm/pkg/kobject"
	"fmt"
	"net/http"
	"crypto/tls"
	"io/ioutil"
	"path/filepath"
	"strings"
	//"github.com/moby/moby/api/types/swarm"
	//"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"encoding/json"
	"helm.sh/helm/v3/pkg/chartutil"
	//"github.com/kubernetes/kompose/pkg/transformer/kubernetes"
	"os"
	yaml "gopkg.in/yaml.v2"

	//"github.com/ompluscator/dynamic-struct"



)

const (
)

type Values struct {
	ReplicaCount *uint64 `yaml:"replicaCount"`
	Image        struct {
		Repository string `yaml:"repository"`
		PullPolicy string `yaml:"pullPolicy"`
	} `yaml:"image"`
	//ImagePullSecrets []interface{} `yaml:"imagePullSecrets"`
	//ImagePullSecrets map[string]string `yaml:"imagePullSecrets"`
	//ImagePullSecrets []string `yaml:"imagePullSecrets"`

	ImagePullSecrets []Secret `yaml:"imagePullSecrets"`

	NameOverride     string        `yaml:"nameOverride"`
	FullnameOverride string        `yaml:"fullnameOverride"`
	ServiceAccount   struct {
		Create      bool `yaml:"create"`
		Annotations struct {
		} `yaml:"annotations"`
		Name interface{} `yaml:"name"`
	} `yaml:"serviceAccount"`
	PodSecurityContext struct {
	} `yaml:"podSecurityContext"`
	SecurityContext struct {
	} `yaml:"securityContext"`
	Service struct {
		Type string `yaml:"type"`
		Port uint32    `yaml:"port"`
	} `yaml:"service"`
	Ingress struct {
		Enabled     bool `yaml:"enabled"`
		Annotations struct {
		} `yaml:"annotations"`
		Hosts []struct {
			Host  string        `yaml:"host"`
			Paths []interface{} `yaml:"paths"`
		} `yaml:"hosts"`
		TLS []interface{} `yaml:"tls"`
	} `yaml:"ingress"`
	Resources struct {
	} `yaml:"resources"`
	NodeSelector struct {
	} `yaml:"nodeSelector"`
	Tolerations []interface{} `yaml:"tolerations"`
	Affinity    struct {
	} `yaml:"affinity"`
	///Envars map[string]string `yaml:"envars"`
	Env []Envs`yaml:"env"`

}
type Secret struct {
	Name  string        `yaml:"name"`
}
type Envs struct {
	Name  string        `yaml:"name"`
	Value string		`yaml:"value"`
}
// ValidateFlags validates all command line flags
/*
func ValidateFlags(bundle string, args []string, cmd *cobra.Command, opt *kobject.ConvertOptions) {


}*/

// Convert transforms
func Convert(opt kobject.ConvertOptions) {
	//TODO::ADD SERVICE TYPE | REGISTRY SECRET 

	services := GetServicesFromCluster(opt.ServerURL,opt.AuthToken)
	service := FindServiceFromFilter(services, opt.Filter)
	fmt.Println(service.Spec.Name)

	directory := CreateChart(service.Spec.Name)

	values := ReadFile(directory+"/values.yaml")
	
	converted := Map(service,values, "NodePort", "regcred")

	WriteFile(directory+"/values.yaml", converted)
	fmt.Println(opt.DeployService)
	if opt.DeployService {
		fmt.Println("Deploying service.....")
		DeployService(service.Spec.Name)
	}

}

func DeployService(name string) {

}

func Map(service swarm.Service, values Values, serviceType string, registrySecret string) Values {
	//TODO::ADD HASH VALIDATION | REGISTRY SECRET VALUE CHECK

	serviceSpec := service.Spec
	endpointSpec := serviceSpec.EndpointSpec
	containerSpec := serviceSpec.TaskTemplate.ContainerSpec

	splitImage := strings.Split(containerSpec.Image, ":")
    strippedImage := splitImage[0]
	values.Image.Repository = strippedImage
	
	values.ImagePullSecrets =[]Secret{
		Secret{Name: registrySecret}}
	//values.ImagePullSecrets =[]string{registrySecret}
	//values.ImagePullSecrets = map[string]string{"name": registrySecret}

	values.ReplicaCount = serviceSpec.Mode.Replicated.Replicas
	
	//TODO::Multiple Ports
	values.Service.Port = endpointSpec.Ports[0].TargetPort
	
	
	values.Service.Type = serviceType

	//envMap := make(map[string]string, len(containerSpec.Env))
	values.Env = make([]Envs,len(containerSpec.Env))

	for i,envar := range containerSpec.Env {
		s := strings.Split(envar, "=")
		name, value := s[0], s[1]
		values.Env[i] = Envs{Name: name, Value: value}
		//envMap[name] = value
	}
	//values.Envar = envMap
	return values
}



func CreateChart(service string) string {
	err := os.Mkdir(service, 0755)
	x_, err := chartutil.Create(service, filepath.Dir(service))
	if err != nil {
		fmt.Println(err)
	}
	return x_
}

func GetServicesFromCluster(server string, token string) []swarm.Service {

	response := MakeRequest(server,token,"/services", true)

	if response != nil {
		services := []swarm.Service{}
		err := json.Unmarshal(response, &services)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return services
	}
	return nil
}

func FindServiceFromFilter(services []swarm.Service, filter string) swarm.Service {
	var foundService swarm.Service
	for _,service := range services {
		if filter == service.Spec.Name {
			//fmt.Printf("%#v\n", service);
			foundService = service
		}
	}
	return foundService
}


func MakeRequest(server string, token string, endpoint string, isInsecure bool) []byte {
	url := server + endpoint
	method := "GET"
  
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: isInsecure},
	}
	
	client := &http.Client {Transport: tr}
	req, err := http.NewRequest(method, url, nil)
  
	req.Header.Add("Authorization", "Bearer " + token)
  
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	return body
}


func WriteFile(location string, data Values) {
	d, err := yaml.Marshal(&data)
	if err != nil {
		fmt.Println("error: %v", err)
	}
	err = ioutil.WriteFile(location, d, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func ReadFile(location string) Values {
	data, err := ioutil.ReadFile(location)
	if err != nil {
        fmt.Println("File reading error", err)
	}
	
	var values Values
	err = yaml.Unmarshal(data, &values)
	if err != nil {
        fmt.Println("File reading error", err)
	}

	return values
}