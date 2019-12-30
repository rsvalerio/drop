package drop

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/metrics"

	"github.com/caddyserver/caddy"
	"github.com/miekg/dns"
)

func init() { plugin.Register("drop", setup) }

func setup(c *caddy.Controller) error {
	d, err := parse(c)
	if err != nil {
		return plugin.Error("Drop", err)
	}

	// Register all metrics.
	c.OnStartup(func() error {
		metrics.MustRegister(c, dropCount)
		return nil
	})

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		d.Next = next
		return d
	})

	return nil
}

func parse(c *caddy.Controller) (Drop, error) {
	d := Drop{}

	c.Next() //drop
	if len(c.RemainingArgs()) == 0 {
		return d, c.Errf("its expected a valid query type to drop: [A, AAAA, TXT, ANY, ...]")
	}

	c.NextArg()
	d.qtype = dns.StringToType[c.Val()]
	if dns.TypeNone == d.qtype {
		return d, c.Errf("unexpected token %q; expect valid query type: [A, AAAA, TXT, ANY, ...]", c.Val())
	}

	return d, nil
}
