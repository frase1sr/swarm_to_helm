
package templateutils

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
	Filter                      string
}