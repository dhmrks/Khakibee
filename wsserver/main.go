package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Tkanos/gonfig"
	"github.com/gorilla/websocket"
)

// Examples
// 1) https://dev.to/danielkun/go-asynchronous-and-safe-real-time-broadcasting-using-channels-and-websockets-4g5d
// 2) https://github.com/gorilla/websocket/blob/master/examples/chat/main.go
// 3) https://tutorialedge.net/golang/go-websocket-tutorial/

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Configuration :configuration data type
type Configuration struct {
	Port        *int   `json:"port"`
	ContextPath string `json:"context_path"`
}

// Client struct
type Client struct {
	ch chan []byte
	ws *websocket.Conn
}

// Body struct
type Body struct {
	Msg string `json:"msg"`
}

var serverChan chan *Client
var unregister chan *Client
var broadcast chan []byte

func main() {
	conf := Configuration{}
	err := gonfig.GetConf("config.json", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	var port = "3020"
	if conf.Port != nil {
		port = fmt.Sprintf(":%d", *conf.Port)
	}

	serverChan = make(chan *Client, 256)
	unregister = make(chan *Client, 256)
	broadcast = make(chan []byte, 256)

	go server(serverChan)

	http.HandleFunc(conf.ContextPath+"/ws", wsEndpoint)
	http.HandleFunc(conf.ContextPath+"/ws/broadcast", bcEndpoint)

	fmt.Println("Server started")
	fmt.Println("Port ", port)
	fmt.Println("Context path ", conf.ContextPath)
	log.Fatal(http.ListenAndServe(port, nil))
}

func bcEndpoint(w http.ResponseWriter, r *http.Request) {
	var body Body
	_ = json.NewDecoder(r.Body).Decode(&body)

	broadcast <- []byte(body.Msg)

	w.Write([]byte("ok"))
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Handler Upgrade error: %v", err)
	}

	client := &Client{ws: ws, ch: make(chan []byte, 256)}
	serverChan <- client

	go writer(client)
	go reader(client)
}

func writer(c *Client) {
	for {
		select {
		case msg, _ := <-c.ch:
			err := c.ws.WriteMessage(1, msg)
			if err != nil {
				log.Printf("WriteMessage error: %v", err)
			}
		}
	}
}

func reader(c *Client) {
	defer func() {
		unregister <- c
	}()
	for {
		_, p, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("ReadMessage error: %v", err)
			}
			return
		}
		broadcast <- p
	}
}

func server(serverChan chan *Client) {
	clients := make(map[*Client]bool)
	for {
		select {
		case client, _ := <-serverChan:
			clients[client] = true
		case msg, _ := <-broadcast:
			for c := range clients {
				c.ch <- msg
			}
		case client, _ := <-unregister:
			client.ws.Close()
			delete(clients, client)
		}
	}
}
