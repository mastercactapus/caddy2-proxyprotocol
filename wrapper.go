package caddy2proxyprotocol

import (
	"fmt"
	"net"
	"time"

	"github.com/caddyserver/caddy/v2"
	"github.com/mastercactapus/proxyprotocol"
)

type Wrapper struct {
	Timeout caddy.Duration `json:"timeout,omitempty"`
	Allow   []string       `json:"allow,omitempty"`

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
