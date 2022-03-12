package caddy2proxyprotocol

import (
	"fmt"
	"net"
	"time"

	"github.com/caddyserver/caddy/v2"
	"github.com/pires/go-proxyproto"
)

// Wrapper provides PROXY protocol support to Caddy by implementing the caddy.ListenerWrapper interface. It must be loaded before the `tls` listener.
type Wrapper struct {
	// Timeout specifies an optional maximum time for the PROXY header to be received. If zero, timeout is disabled. Default is 5s.
	Timeout caddy.Duration `json:"timeout,omitempty"`
}

func (pp *Wrapper) Provision(ctx caddy.Context) error {
	return nil
}

func (pp *Wrapper) WrapListener(l net.Listener) net.Listener {
    proxyListener := &proxyproto.Listener{
		Listener:          l,
		ReadHeaderTimeout: time.Duration(pp.Timeout),
	}
	return proxyListener
}
