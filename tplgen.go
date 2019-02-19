package tplgen

import (
	"strings"
)

type Router struct {
	Hostname string
	Serial string
	LoopbackIPAddress string
}


func CreateRouterConfiguration(router Router) Router {
	return router
}

func RemoveCidr(s string) string {
	if strings.Contains(s, "/") {
		return strings.Split(s, "/")[0]
	}
	return s
}


