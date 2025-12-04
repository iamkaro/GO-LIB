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

package websocket

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type (
	Server  struct{ handle handler }
	handler func(this *Client)
)

func NewServer(handle handler) *Server { return &Server{handle: handle} }

func (it *Server) ListenAndServe(addr string) {

	mux := http.NewServeMux()
	mux.Handle("/", websocket.Handler(func(conn *websocket.Conn) {
		fmt.Println("New client, ", conn.RemoteAddr().String())
		it.handle(&Client{Client: conn})
	}))

	serv := http.Server{Addr: addr, Handler: mux}
	err := serv.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}
}
