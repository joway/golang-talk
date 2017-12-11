// https://github.com/kubernetes/kubernetes/blob/master/pkg/util/env/env.go

package main

import (
	eventServiceLib "github.com/the-control-group/event-service-lib"
	"fmt"
)

func main() {
	hostname := eventServiceLib.GetDiamondHostname()
	fmt.Println(hostname)
}
