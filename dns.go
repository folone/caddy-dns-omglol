// Package omglol implements a Caddy DNS provider module for omg.lol.
// It wraps github.com/libdns/omglol so that Caddy can use the omg.lol DNS
// API to complete ACME DNS-01 challenges.
package omglol

import (
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	omglol "github.com/folone/libdns-omglol"
)

// Provider wraps the libdns omg.lol provider as a Caddy module.
type Provider struct{ *omglol.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.omglol",
		New: func() caddy.Module { return &Provider{new(omglol.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()
	p.Provider.APIKey = strings.TrimSpace(repl.ReplaceAll(p.Provider.APIKey, ""))
	p.Provider.Address = strings.TrimSpace(repl.ReplaceAll(p.Provider.Address, ""))
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens.
//
// Syntax:
//
//	omglol {
//	    api_key  <api_key>
//	    address  <omglol_address>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_key":
				if d.NextArg() {
					p.Provider.APIKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "address":
				if d.NextArg() {
					p.Provider.Address = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.ArgErr()
			}
		}
	}
	if p.Provider.APIKey == "" {
		return d.Err("api_key is required")
	}
	if p.Provider.Address == "" {
		return d.Err("address is required")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
