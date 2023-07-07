package main

type Hub struct {

	clients map[*Client] bool

	broadcast chan []byte

	register chan *Client
	
	unregister chan *Client
}

// Hub의 포인터값 (주소값 말고 들어간 값 ) return 하는 함수
func newHub() *Hub{

	//주소값으로 내보내서 
	return &Hub{
		broadcast: make(chan []byte),
		register: make(chan *Client),
		unregister: make(chan *Client),
		clients: make(map[*Client]bool),
	}
}

func (hub *Hub) run(){
	for{
		select {
		
		//client가 register일때
		case client := <-hub.register:
			hub.clients[client]=true
		
		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok{
				delete(hub.clients, client)
				close(client.send)
			}
		
		case message := <-hub.broadcast:
			for client := range hub.clients{

			}
		}

				
		
	}
}