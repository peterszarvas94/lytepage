FROM golang:1.27-bookworm AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
WORKDIR /src/scaffhold
RUN go mod edit -replace github.com/peterszarvas94/lytepage=.. \
    && go mod tidy \
    && go run ./cmd/static --no-serve

FROM caddy:2.11-alpine
COPY docker/Caddyfile /etc/caddy/Caddyfile
COPY --from=build /src/scaffhold/public/ /srv/
USER 10001:10001
EXPOSE 8080
