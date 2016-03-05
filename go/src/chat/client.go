package main

import (
    "github.com/gorilla/websocket"
    "time"
)

// clientはチャットを行っており1人のユーザーを表します
type client struct {
    socket *websocket.Conn
    send chan *message
    room *room
    userData map[string]interface{}
}

func (c *client) read() {
    for {
        /*
        if _, msg, err := c.socket.ReadMessage(); err == nil {
            c.room.forward <- msg
        } else {
            break;
        }
        */
        var msg *message
        if err := c.socket.ReadJSON(&msg); err == nil {
            msg.When = time.Now()
            msg.Name = c.userData["name"].(string)
            c.room.forward <- msg
        } else {
            break
        }
    }
    c.socket.Close()
}

func (c *client) write() {
    for msg := range c.send {
        // if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
        if err := c.socket.WriteJSON(msg); err != nil {
            break
        }
    }
    c.socket.Close()
}