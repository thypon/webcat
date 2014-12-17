package libwebcat

type Client interface {
	Receive() (string, error)
	Send(string) error
	Run() error
}
