package main

import (
	"flag"
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// Add more data to this type if needed
type client struct {
	isClosing bool
	mu        sync.Mutex
}

var (
	clients    = make(map[*websocket.Conn]*client) // Note: although large maps with pointer-like types (e.g. strings) as keys are slow, using pointers themselves as keys is acceptable and fast
	register   = make(chan *websocket.Conn)        // 새 클라이언트 등록용 채널
	broadcast  = make(chan string)                 // 메시지 브로드캐스트용 채널
	unregister = make(chan *websocket.Conn)        // 클라이언트 연결 해제용 채널
)

func runHub() {
	for {
		select {
		case connection := <-register: // 새 클라이언트 등록
			clients[connection] = &client{} // 오호라 포인터를 키로 사용하네
			log.Println("connection registered")

		case message := <-broadcast: // 모든 클라이언트에게 메시지 전송
			log.Println("message received:", message)
			// Send the message to all clients
			for connection, c := range clients {
				go func(connection *websocket.Conn, c *client) { // send to each client in parallel so we don't block on a slow client
					c.mu.Lock()
					defer c.mu.Unlock()
					if c.isClosing {
						return
					}
					if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
						c.isClosing = true
						log.Println("write error:", err)

						connection.WriteMessage(websocket.CloseMessage, []byte{})
						connection.Close()
						unregister <- connection
					}
				}(connection, c)
			}

		case connection := <-unregister: // 클라이언트 연결 해제
			// Remove the client from the hub
			delete(clients, connection)

			log.Println("connection unregistered")
		}
	}
}

func main() {
	app := fiber.New()

	app.Static("/", "./home.html")

	app.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) { // Returns true if the client requested upgrade to the WebSocket protocol
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	go runHub()

	app.Get("/ws", websocket.New(func(c *websocket.Conn) { // Client가 ws://localhost:3000/ws 로 웹소켓 입장함.
		// When the function returns, unregister the client and close the connection
		defer func() {
			unregister <- c
			c.Close()
		}()

		// Register the client
		register <- c

		// 무한 루프가 존재함 -> 알아서 websocket middleware에서 이 함수를 go루틴으로 실행한다는 뜻
		for {
			messageType, message, err := c.ReadMessage() // blocking
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}

				return // Calls the deferred function, i.e. closes the connection on error
			}

			if messageType == websocket.TextMessage {
				// Broadcast the received message
				broadcast <- string(message)
			} else {
				log.Println("websocket message received of type", messageType)
			}
		}
	}))

	addr := flag.String("addr", ":8080", "http service address")
	flag.Parse()
	log.Fatal(app.Listen(*addr))
}
