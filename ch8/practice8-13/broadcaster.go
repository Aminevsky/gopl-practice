package main

type client chan<- string // 送信用メッセージチャネル

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // クライアントから受信するすべてのメッセージ
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
