package httpclient

type Interface interface {
	BuildRequst() *Request
	Get()
}
