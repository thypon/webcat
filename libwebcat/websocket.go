package libwebcat

import (
	"flag"
	"log"
	"net/url"

	"golang.org/x/net/websocket"
)

type WSClient struct {
	ReadWriter
	channel *websocket.Conn
}

func NewWSClient(url *url.URL, rw ReadWriter) (Client, error) {
	var origin string
	if Origin != "" {
		origin = Origin
	} else {
		origin = ""
		switch url.Scheme {
		case "ws":
			origin += "http://"
		case "wss":
			origin += "https://"
		default:
			log.Fatal("No known websocket scheme")
		}
		origin += url.Host
	}
	w, err := websocket.Dial(flag.Arg(0), "", origin)
	if err != nil {
		return nil, err
	}
	return &WSClient{rw, w}, nil
}

func (c *WSClient) Receive() (string, error) {
	var msg string
	err := websocket.Message.Receive(c.channel, &msg)
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (c *WSClient) Send(msg string) error {
	return websocket.Message.Send(c.channel, msg)
}

func (c *WSClient) Run() error {
	err := make(chan error)

	go func() {
		for {
			s, e := c.Receive()
			if e != nil {
				err <- e
				close(err)
				break
			}
			e = c.Write(s)
			if e != nil {
				err <- e
				close(err)
				break
			}
		}
	}()

	go func() {
		for {
			s, e := c.Read()
			if e != nil {
				err <- e
				close(err)
				break
			}

			e = c.Send(s)
			if e != nil {
				err <- e
				close(err)
				break
			}
		}
	}()

	return <-err
}
