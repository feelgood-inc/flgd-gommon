package httpclient

type ClientConfig struct {
	Host    string
	Port    string
	Headers map[string]string
}
