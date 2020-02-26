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

package deploy

import (
	//"github.com/frase1sr/swarm_to_helm/cmd"
	//"github.com/spf13/cobra"
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


)