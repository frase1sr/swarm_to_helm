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
	//"github.com/moby/moby/api/types/swarm"
	//"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"encoding/json"
	"helm.sh/helm/v3/pkg/chartutil"
	//"os"

)

const (
	// DefaultComposeFile name of the file that kompose will use if no file is explicitly set
	DefaultComposeFile = "docker-compose.yml"
)


// ValidateFlags validates all command line flags
/*
func ValidateFlags(bundle string, args []string, cmd *cobra.Command, opt *kobject.ConvertOptions) {


}*/

// Convert transforms docker compose or dab file to k8s objects
func Convert(opt kobject.ConvertOptions) {
	services := GetServicesFromCluster(opt.ServerURL,opt.AuthToken)
	service := FindServiceFromFilter(services, opt.Filter)
	fmt.Println(service.Spec.Name)
	fmt.Println(create.defaultValues)
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



