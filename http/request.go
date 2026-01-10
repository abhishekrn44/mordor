package http

type Request struct {
	Method      string
	Target      string
	Version     string
	Headers     map[string]string
	MessageBody []byte
}
