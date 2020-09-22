package caddy2proxyprotocol

import "github.com/caddyserver/caddy/v2"

func init() {
	caddy.RegisterModule(ProxyProtocolSupport{})
}

func (ProxyProtocolSupport) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "caddy.listeners.proxyprotocol",
		New: func() caddy.Module { return new(ProxyProtocolSupport) },
	}
}
