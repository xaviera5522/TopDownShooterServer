package router

import (
	"sync"
)

type Router struct {
	routes sync.Map
}

func newRouter() *Router {
	router := &Router{}
	//Bins Services here: router.routes.Store()

	return router
}

func (r *Router) Find(service string) (string, bool) {
	v, found := r.routes.Load(service)
	if !found {
		return "", false
	}
	return v.(string), true
}
