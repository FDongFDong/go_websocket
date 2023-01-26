package p2p

import (
	"coin/test/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	// port 3000이 port 4000에서온 request를 update합니다.

	// CheckOrigin은 true나 false를 리턴하는데 upgrade할 때 이게 필요하다
	// 유효한 websocket연결인지 인증할때 쓴다.
	// HTTP의 쿠키나 토큰처럼 허가된 사람만 접속시켜야한다.
	// Upgrader안에 CheckOrigin 함수가 있다. 이것을 작동시켜야한다.
	// 해당 함수는 request를 받아서 cookie나 origin을 확인한다.
	// 여기서는 일단 항상 true를 선언하게 한다.
	// 그러면 websocket가 연결이 허가된다.
	openPort := r.URL.Query().Get("openPort")
	ip := utils.Splitter(r.RemoteAddr, ":", 0)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		// 열린 포트와 IP가 모두 존재하는 경우 true
		return openPort != "" && ip != ""
	}

	// 3000에서 4000으로 소통하는 방법
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	initPeer(conn, ip, openPort)
}

func AddPeer(address, port, openPort string) {
	// from :4000이 port:3000으로 연결하길 원해요
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort[1:]), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)
}
