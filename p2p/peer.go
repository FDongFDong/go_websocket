package p2p

import (
	"fmt"

	"github.com/gorilla/websocket"
)

var Peers map[string]*peer = make(map[string]*peer)

// 연결을 저장하기 위함
type peer struct {
	conn  *websocket.Conn
	inbox chan []byte
}

func (p *peer) read() {
	// Error이 나오면 해당 peer를 삭제한다.
	for {
		_, m, err := p.conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%s", m)
	}
}
func (p *peer) write() {
	for {
		// 메세지가 오기를 기다린다.
		m := <-p.inbox
		p.conn.WriteMessage(websocket.TextMessage, m)
	}
}
func initPeer(conn *websocket.Conn, address, port string) *peer {
	key := fmt.Sprintf("%s:%s", address, port)
	p := &peer{
		conn,
		make(chan []byte),
	}
	go p.read()
	go p.write()
	Peers[key] = p
	return p
}

// :3000 포트의 peers
// {
// 	"127.0.0.1:4000": conn
// }
// :4000 포트의 peers
// {
// 	"127.0.0.1:4000": conn
// }
