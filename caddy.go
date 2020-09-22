package caddy2proxyprotocol

import "github.com/caddyserver/caddy/v2"

var (
	_ = caddy.Provisioner(&ProxyProtocolSupport{})
	_ = caddy.Validator(&ProxyProtocolSupport{})
	_ = caddy.Module(&ProxyProtocolSupport{})
	_ = caddy.ListenerWrapper(&ProxyProtocolSupport{})
)

func init() {
	caddy.RegisterModule(ProxyProtocolSupport{})
}

func (ProxyProtocolSupport) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "caddy.listeners.proxyprotocol",
		New: func() caddy.Module { return new(ProxyProtocolSupport) },
	}
}
