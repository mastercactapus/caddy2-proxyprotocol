package caddy2proxyprotocol

import (
	"fmt"
	"net"
	"time"

	"github.com/caddyserver/caddy/v2"
	"github.com/mastercactapus/proxyprotocol"
)

// Wrapper provides PROXY protocol support to Caddy by implementing the caddy.ListenerWrapper interface. It must be loaded before the `tls` listener.
type Wrapper struct {
	// Timeout specifies an optional maximum time for the PROXY header to be received. If zero, timeout is disabled. Default is 5s.
	Timeout caddy.Duration `json:"timeout,omitempty"`

	// Allow is an optional list of CIDR ranges to allow/require PROXY headers from.
	Allow []string `json:"allow,omitempty"`

	rules []proxyprotocol.Rule
}

func (pp *Wrapper) parseRules() ([]proxyprotocol.Rule, error) {
	rules := make([]proxyprotocol.Rule, 0, len(pp.Allow))
	for _, s := range pp.Allow {
		_, n, err := net.ParseCIDR(s)
		if err != nil {
			return nil, fmt.Errorf("invalid subnet '%s': %w", s, err)
		}
		rules = append(rules, proxyprotocol.Rule{Timeout: time.Duration(pp.Timeout), Subnet: n})
	}

	return rules, nil
}

func (pp *Wrapper) Provision(ctx caddy.Context) error {
	rules, err := pp.parseRules()
	if err != nil {
		return err
	}

	pp.rules = rules

	return nil
}

func (pp *Wrapper) WrapListener(l net.Listener) net.Listener {
	pL := proxyprotocol.NewListener(l, time.Duration(pp.Timeout))
	pL.SetFilter(pp.rules)

	return pL
}
