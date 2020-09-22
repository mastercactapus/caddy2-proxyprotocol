package caddy2proxyprotocol

import (
	"fmt"
	"net"
	"time"

	"github.com/caddyserver/caddy/v2"
	"github.com/mastercactapus/proxyprotocol"
)

type ProxyProtocolSupport struct {
	Timeout string   `json:"timeout,omitempty"`
	Allow   []string `json:"allow,omitempty"`

	timeout time.Duration
	rules   []proxyprotocol.Rule
}

func (pp *ProxyProtocolSupport) parse() (time.Duration, []proxyprotocol.Rule, error) {
	var timeout time.Duration
	var err error
	switch pp.Timeout {
	case "":
		timeout = 5 * time.Second
	case "none", "0":
	default:
		timeout, err = time.ParseDuration(pp.Timeout)
	}
	if err != nil {
		return 0, nil, fmt.Errorf("invalid timeout value: %w", err)
	}

	rules := make([]proxyprotocol.Rule, 0, len(pp.Allow))
	for _, s := range pp.Allow {
		_, n, err := net.ParseCIDR(s)
		if err != nil {
			return 0, nil, fmt.Errorf("invalid subnet '%s': %w", s, err)
		}
		rules = append(rules, proxyprotocol.Rule{Timeout: timeout, Subnet: n})
	}

	return timeout, rules, nil
}

func (pp *ProxyProtocolSupport) Validate() error {
	_, _, err := pp.parse()
	return err
}

func (pp *ProxyProtocolSupport) Provision(ctx caddy.Context) error {
	timeout, rules, err := pp.parse()
	if err != nil {
		return err
	}

	pp.timeout = timeout
	pp.rules = rules

	return nil
}

func (pp *ProxyProtocolSupport) WrapListener(l net.Listener) net.Listener {
	pL := proxyprotocol.NewListener(l, pp.timeout)
	pL.SetFilter(pp.rules)

	return pL
}
