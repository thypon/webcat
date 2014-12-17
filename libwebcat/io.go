package libwebcat

import (
	"bufio"
	"fmt"
)

type ReadWriter interface {
	Reader
	Writer
}

type Reader interface {
	Read() (string, error)
}

type Writer interface {
	Write(string) error
}

type STDIO struct {
	rw *bufio.ReadWriter
}

func NewSTDIO() *STDIO {
	return &STDIO{}
}

func (stdio *STDIO) Read() (string, error) {
	var str string
	_, err := fmt.Scanln(&str)
	return str, err
}

func (stdio *STDIO) Write(s string) error {
	_, err := fmt.Println(s)
	return err
}
