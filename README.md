# Go Pokédex

An interactive Go command-line Pokédex built around the public
[PokéAPI](https://pokeapi.co/). Browse locations, discover Pokémon, attempt catches,
and inspect the collection you build during a session.

This is a learning project in HTTP clients, JSON decoding, REPL design, and
concurrent cache access—not a production service. It makes those concepts
inspectable in a small application without a frontend or database.

## How it works

- [The REPL](repl.go) normalizes input and dispatches commands through a command
  registry. [main.go](main.go) owns the API client and session collection.
- [The API client](internal/pokeapi/client.go) uses an HTTP timeout. Location
  browsing follows PokéAPI's next/previous page URLs in
  [the location client](internal/pokeapi/location_list.go).
- [The cache](internal/pokecache/client.go) stores response bytes by key, protects
  access with a mutex, and removes old entries from a background goroutine.
- [Catching](command_catch.go) uses a randomized outcome; successful catches go
  into an in-memory map. [Inspection](command_inspect.go) reads that collection,
  not an independently persisted database.

## Decisions and tradeoffs

**Cache response bytes, not domain objects.** The API layer owns JSON decoding,
while the cache only stores bytes. Repeated cached requests avoid another HTTP
call. Expiration is periodic cleanup, not a strict freshness check on every read.

**Keep the collection in memory.** A map is enough for a single interactive
session, but catches disappear when the program exits.

**Separate commands from transport.** Command handlers format terminal output;
the API client handles HTTP requests. This keeps each command small without
introducing a framework.

## Run locally

Use the Go version declared in [go.mod](go.mod), currently **Go 1.27.2**.
Network access to PokéAPI is needed for browsing and catching; no API key is
required. Build from source rather than relying on the checked-in binary.

```bash
go run .
```

At the prompt:

```text
help
map
explore pastoria-city-area
catch pikachu
inspect pikachu
pokedex
exit
```

`mapb` returns to the previous location page. A catch can fail, so `inspect` only
works after a successful catch. Use `exit` to end the session; EOF handling needs
improvement. There is no hosted demo.

## Tests and current status

```bash
go test ./...
```

Tests cover input normalization, help/exit behavior, the command registry, and
cache storage/cleanup. **The current suite fails `TestGetCommands`: its expected
list still has four commands, while the application has eight.** The `help` and
`exit` CLI flow was exercised successfully during this documentation update;
that does not verify live API behavior.

Next improvements are bringing command coverage up to date, graceful EOF
handling, and API error-path tests before adding a CI gate. No CI workflow or
license is currently included.
