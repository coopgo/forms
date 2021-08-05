package forms

// Transport is the transport interface for communication with the outside world
type Transport interface {
	Run(*Service)
}

// NoTransport is a fake transport
type NoTransport struct {
}

func (t NoTransport) Run(s *Service) {}
