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

package kobject

import (
	//"github.com/pkg/errors"
	//"github.com/spf13/cast"
	//"path/filepath"
	//"time"
	//"github.com/docker/docker/api/types"
	//"github.com/docker/docker/api/types/swarm"
)

// ConvertOptions holds all options that controls transformation process
type ConvertOptions struct {
	ServerURL                    string
	AuthToken				   string
	Out 						string
}