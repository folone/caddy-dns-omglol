# Build a custom Caddy binary that includes the omg.lol DNS provider plugin.
#
# Usage:
#   docker build -t caddy-omglol .
#   docker run --rm -v $PWD/Caddyfile:/etc/caddy/Caddyfile caddy-omglol
#
# Or extract the binary only:
#   docker create --name tmp caddy-omglol
#   docker cp tmp:/usr/bin/caddy ./caddy
#   docker rm tmp

FROM caddy:builder AS builder

RUN xcaddy build \
    --with github.com/folone/caddy-dns-omglol

FROM caddy:latest

COPY --from=builder /usr/bin/caddy /usr/bin/caddy
