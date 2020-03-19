package route

import (
	"net/http"

	"github.com/jhunt/go-log"
)

type Handler func(r *Request)

type route struct {
	matcher matcher
	handler Handler
}

type Router struct {
	Name   string
	Debug  bool
	routes []route
}

func (r *Router) Dispatch(match string, handler Handler) {
	r.routes = append(r.routes, route{
		matcher: newMatch(match),
		handler: handler,
	})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	request := NewRequest(w, req, r.Debug)

	for _, rt := range r.routes {
		if args, ok := rt.matcher(req); ok {
			w.Header().Set("Content-Type", "application/json")

			request.Args = args
			rt.handler(request)
			if !request.Done() {
				log.Errorf("%s handler bug: failed to call either OK() or Fail()", request)
				request.Fail(Oops(nil, "an unknown error has occurred"))
			}
			return
		}
	}

	name := r.Name
	if name != "" {
		name = name + " "
	}
	request.Fail(NotFound(nil, name+"API endpoint `%s' not found", request))
}
