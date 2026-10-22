FROM golang:1.27-bookworm AS build
WORKDIR /src/scaffhold
COPY scaffhold/go.mod scaffhold/go.sum ./
RUN go mod download
COPY scaffhold/ ./
RUN go run ./cmd/static --no-serve

FROM caddy:2.11-alpine
COPY docker/Caddyfile /etc/caddy/Caddyfile
COPY --from=build /src/scaffhold/public/ /srv/
USER 10001:10001
EXPOSE 8080
