/*! 
 * I am Karo  😊👍 
 * 
 * Contact me:  
 *     https://www.karo.link/ 
 *     https://github.com/iamkaro
 *     https://www.linkedin.com/in/iamkaro  
 * 
 * My own GO-based library 
 * https://github.com/iamkaro/GO-LIB.git
 * Copyright © 2021 developed. 
 */

package httpx

import (
	"fmt"
	"net/http"
)

func NewServer(addr string, requestTimeout, responseTimeout Seconds, printRequests bool) (out *Server) {
	out = &Server{
		routes: routes{
			Routes:        Routes{},
			PrintRequests: printRequests,
		},
		server: &http.Server{
			Addr:              addr,
			Handler:           nil,
			ReadTimeout:       requestTimeout.Duration(),
			ReadHeaderTimeout: requestTimeout.Duration(),
			WriteTimeout:      responseTimeout.Duration(),
			MaxHeaderBytes:    1 << 20,
		},
	}
	out.server.Handler = out.routes
	return
}

type (
	Server struct {
		server *http.Server
		routes routes
	}
	routes struct {
		Routes        Routes
		PrintRequests bool
	}
	Routes map[string]Handler
)

func (it routes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := it.Routes[r.URL.String()]; ok && (handler != nil) {
		w.WriteHeader(http.StatusOK)
		if it.PrintRequests {
			fmt.Println("request " + r.URL.String() + " ok")
		}
		handler(&Response{ResponseWriter: w}, &Request{Request: *r})
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Header().Add("content-type", "text/html; charset=UTF-8")
	_, _ = w.Write([]byte("404 - not found!"))
	fmt.Println("request " + r.URL.String() + " not found!")
}

func (it *Server) SetRoutes(list Routes) {
	for s, handler := range list {
		it.AddRoute(s, handler)
	}
}

func (it *Server) AddRoute(route string, handler Handler) {
	it.routes.Routes[route] = handler
}
func (it *Server) RemoveRoute(route string) {
	if _, ok := it.routes.Routes[route]; ok {
		it.routes.Routes[route] = nil
		delete(it.routes.Routes, route)
	}
}

func (it *Server) Listen() error {
	fmt.Println("sever running on http://localhost" + it.server.Addr)
	return it.server.ListenAndServe()
}

func (it *Server) Close() error {
	return it.server.Close()
}
