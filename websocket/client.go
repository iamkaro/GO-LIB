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

	"golang.org/x/net/websocket"
)

type (
	Client struct{ Client *websocket.Conn }
)

func NewClient(socket, myHost string) *Client {
	conn, err := websocket.Dial(socket, "", myHost)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &Client{Client: conn}
}

func (it *Client) Close() {
	if it.Client.IsServerConn() {
		fmt.Println("Exit client, ", it.Client.RemoteAddr().String())
	}

	err := it.Client.Close()
	if err != nil {
		fmt.Println("\t\tError close: ", err.Error())
	}
}

func (it *Client) Receive(v interface{}) bool {
	err := websocket.JSON.Receive(it.Client, v)
	if err != nil {
		fmt.Println("Error receive: ", err)
		return false
	}
	return true
}

func (it *Client) Send(v interface{}) bool {
	err := websocket.JSON.Send(it.Client, v)
	if err != nil {
		fmt.Println("Error send: ", err)
		return false
	}
	return true
}
