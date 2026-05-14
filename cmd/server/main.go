package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/muesli/termenv"

	"ssh-yassine/internal/ui"
)

const (
	host    = "0.0.0.0"
	port    = "2222"
	keyPath = "./keys/ssh_host_ed25519_key"
)

func main() {
	lipgloss.SetColorProfile(termenv.TrueColor)

	if err := os.MkdirAll("./keys", 0700); err != nil {
		log.Fatal(err)
	}

	s, err := wish.NewServer(
		wish.WithAddress(host+":"+port),
		wish.WithHostKeyPath(keyPath),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("SSH server listening on %s:%s", host, port)
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Server stopped: %v", err)
			done <- syscall.SIGTERM
		}
	}()

	<-done
	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	m := ui.NewModel()
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}
