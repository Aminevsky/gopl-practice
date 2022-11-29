package main

//type client chan string // 送信用メッセージチャネル

type client struct {
	Channel    chan string
	ClientName string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // クライアントから受信するすべてのメッセージ
)

func broadcaster() {
	clients := make(map[chan string]string)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli.Channel] = cli.ClientName
			for _, name := range clients {
				cli.Channel <- name
			}

		case cli := <-leaving:
			delete(clients, cli.Channel)
			close(cli.Channel)
		}
	}
}
