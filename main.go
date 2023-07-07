package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
}

var clients []websocket.Conn

func main(){
	// ENDPOINT 설계
	http.HandleFunc("/echo",func(w http.ResponseWriter, r *http.Request) {

		//config 초기화

		conn,_ := upgrader.Upgrade(w,r,nil);

		clients =append(clients, *conn);

		//client루프
		for{

			//메세지 읽기
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Fatal(err)
			}

			//메시지 노출
			fmt.Println("%s send: %s\n", conn.RemoteAddr(),string(msg))

			//브라우저에 메[시지 노출
			for _,client := range clients {
				 if err = client.WriteMessage(msgType,msg);err != nil{return}
			}
		}
	})

	//화면으로 보내기

}