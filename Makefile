# dev mode:

dev/typescript:
	cd theme/script && \
	npm run dev

dev/templ:
	templ generate --watch --proxy="http://localhost:8080" --open-browser=false -v

dev/server:
	air -c .air.server.toml

dev/assets:
	air -c .air.assets.toml

dev:
	make -j4 dev/typescript dev/templ dev/server dev/assets

# ssr mode:

build:
	templ generate
	go build -o ./bin/ssr ./cmd/ssr

ssr:
	./bin/ssr
	
# static mode:

gen:
	templ generate
	go run ./cmd/gen

static:
	go run ./cmd/static
