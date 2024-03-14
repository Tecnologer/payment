package api

import "fmt"

type router struct {
	host string
}

func newRouter(host string) router {
	return router{
		host: host,
	}
}

func (r router) buildPath(path string) string {
	return fmt.Sprintf("%s%s", r.host, path)
}
