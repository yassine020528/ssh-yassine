# Terminal Site

An SSH-based terminal portfolio built with Go, Wish, Bubble Tea, and Lip Gloss.

Instead of serving a traditional web page, this project runs an SSH server that opens an interactive terminal UI. Visitors connect with `ssh`, navigate the portfolio with keyboard controls, and browse sections for about, projects, education, experience, and contact.

## Features

- Interactive SSH portfolio served on port `2222` locally
- Bubble Tea TUI with full-screen rendering
- Animated splash screen with terminal-style rain
- Keyboard navigation for portfolio sections
- Theme switching with multiple color palettes
- Clickable terminal links using OSC 8 escape sequences
- Docker and Fly.io deployment support
- Persistent SSH host key support through `./keys`

## Tech Stack

- Go 1.24
- Charm Bubble Tea
- Charm Wish
- Charm Lip Gloss
- Fly.io
- Docker

## Requirements

- Go `1.24.2` or compatible
- An SSH client
- Optional: Docker
- Optional: Fly CLI for deployment

## Run Locally

Start the SSH server:

```bash
go run ./cmd/server
```

The server listens on:

```text
0.0.0.0:2222
```

In another terminal, connect with:

```bash
ssh localhost -p 2222
```

If your SSH client warns about a new host key during local development, that is expected. The app creates or reuses a host key at:

```text
./keys/ssh_host_ed25519_key
```

## Controls

```text
enter / space        continue or select
up / k               move up
down / j             move down
esc / backspace      return to menu
t                    change theme
q / ctrl+c           quit
```

## Project Structure

```text
cmd/server/          SSH server entrypoint
internal/ui/         Bubble Tea application state and routing
internal/pages/      Individual portfolio pages
internal/view/       Shared theme, logo, and terminal helpers
.github/workflows/   Fly.io deployment workflow
Dockerfile           Production container build
fly.toml             Fly.io service configuration
```

## Editing Content

Most portfolio content lives in `internal/pages`.

```text
about.go        About page copy, logo, and typewriter animation
projects.go     Project list, tech stacks, and links
education.go    Education entries
experience.go   Experience entries
contact.go      Contact links
menu.go         Main menu labels and descriptions
splash.go       Welcome screen text
```

Theme colors are defined in:

```text
internal/view/theme.go
```

## Links

Project and contact links use terminal hyperlinks. Terminals that support OSC 8 links will make them clickable. Terminals that do not support OSC 8 will still show the label text.

For projects, an empty `Link` field is skipped, so projects without public URLs do not display a placeholder URL.

## Colors In Deployment

The server forces a truecolor Lip Gloss profile at startup so colors render correctly when deployed in Docker/Fly environments where stdout color detection may otherwise fail.

For best results, connect from a terminal with 256-color or truecolor support. Common values include:

```text
xterm-256color
screen-256color
tmux-256color
```

## Docker

Build the image:

```bash
docker build -t ssh-yassine .
```

Run it locally:

```bash
docker run --rm -p 2222:2222 ssh-yassine
```

Then connect:

```bash
ssh localhost -p 2222
```

## Fly.io Deployment

This repo includes `fly.toml` and a GitHub Actions workflow for deployment.

The Fly service maps external TCP port `22` to the app's internal port `2222`, so deployed users can connect over standard SSH:

```bash
ssh ssh-yassine.fly.dev
```

If using a custom domain, connect to that host instead:

```bash
ssh your-domain.example
```

The Fly config mounts a volume at:

```text
/app/keys
```

That keeps the SSH host key stable across deploys/restarts.

Manual deploy:

```bash
flyctl deploy --remote-only
```

GitHub Actions deploys automatically on pushes to `main` when `FLY_API_TOKEN` is configured as a repository secret.

## Development Commands

Run all tests:

```bash
go test ./...
```

Format Go files:

```bash
gofmt -w cmd internal
```

Tidy dependencies:

```bash
go mod tidy
```

## Notes

This is intentionally not a web app. The primary interface is SSH. The Dockerfile exposes port `2222`, while Fly.io exposes it publicly through port `22`.

