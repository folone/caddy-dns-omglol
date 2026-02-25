# caddy-dns-omglol

[![Go Reference](https://pkg.go.dev/badge/github.com/folone/caddy-dns-omglol.svg)](https://pkg.go.dev/github.com/folone/caddy-dns-omglol)
[![CI](https://github.com/folone/caddy-dns-omglol/actions/workflows/ci.yml/badge.svg)](https://github.com/folone/caddy-dns-omglol/actions/workflows/ci.yml)

[omg.lol](https://omg.lol) DNS provider module for [Caddy](https://caddyserver.com).

This package registers a Caddy module (`dns.providers.omglol`) that allows Caddy's built-in ACME client to complete DNS-01 challenges for domains hosted on omg.lol. It wraps [`libdns-omglol`](https://github.com/folone/libdns-omglol).

## Caddy module name

```
dns.providers.omglol
```

## Building Caddy with this plugin

### With xcaddy

```bash
xcaddy build --with github.com/folone/caddy-dns-omglol
```

### With Docker

A `Dockerfile` is included in this repository. Build it with:

```bash
docker build -t caddy-omglol .
docker run --rm -v $PWD/Caddyfile:/etc/caddy/Caddyfile caddy-omglol
```

Or extract just the binary:

```bash
docker create --name tmp caddy-omglol
docker cp tmp:/usr/bin/caddy ./caddy
docker rm tmp
```

## Configuration

### Caddyfile

Global (applies to all sites):

```caddyfile
{
    acme_dns omglol {
        api_key  {env.OMGLOL_API_KEY}
        address  yourname
    }
}
```

Per-site:

```caddyfile
yourname.omg.lol {
    tls {
        dns omglol {
            api_key {env.OMGLOL_API_KEY}
            address yourname
        }
    }
}
```

`address` is your omg.lol handle — the part before `.omg.lol` (e.g. `yourname` for `yourname.omg.lol`).

### JSON

```json
{
  "module": "acme",
  "challenges": {
    "dns": {
      "provider": {
        "name": "omglol",
        "api_key": "{env.OMGLOL_API_KEY}",
        "address": "yourname"
      }
    }
  }
}
```

## Getting your API key

Your omg.lol API key is available in your [account settings](https://home.omg.lol/account). It is recommended to supply it via an environment variable rather than hard-coding it in your Caddyfile.

## Related

- [`libdns-omglol`](https://github.com/folone/libdns-omglol) — the underlying libdns provider
- [Caddy DNS challenges](https://caddyserver.com/docs/automatic-https#dns-challenge)
- [xcaddy](https://github.com/caddyserver/xcaddy)
- [omg.lol API docs](https://api.omg.lol)
