package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bbtea "github.com/charmbracelet/wish/bubbletea"
	"github.com/ryuzinoh/ssh-portfolio/ui"
)

const (
	host = "0.0.0.0"
	port = "1367"
)

func main() {
	s, err := wish.NewServer(
		wish.WithAddress(host+":"+port),
		wish.WithHostKeyPath(".ssh/id_sshsafal"),
		wish.WithMiddleware(
			bbtea.Middleware(teaHandler),
		),
	)
	if err != nil {
		log.Fatalf("Could not create server: %v", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("ssh portfolio listening on %s:%s", host, port)

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()
	<-done
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	m := ui.NewModel(pty.Window.Width, pty.Window.Height)
	return m, []tea.ProgramOption{tea.WithAltScreen()}

}
