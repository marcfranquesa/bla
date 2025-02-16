package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/marcfranquesa/bla/pkg/config"
	"github.com/marcfranquesa/bla/pkg/db"
)

type model struct {
	urls   []db.URL
	cursor int
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v.", err)
	}

	for err := db.Connect(cfg.Database); err != nil; err = db.Connect(cfg.Database) {
		log.Printf("Failed to connect to database. Retrying... Error: %v.", err)
		time.Sleep(2 * time.Second)
	}
	log.Printf("Successfully connected to database.")
	defer db.Close()

	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	urls, err := db.GetAllURLs()
	if err != nil {
		log.Fatalf("Failed to get urls: %v.", err)
	}
	return model{
		urls: urls,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.urls)-1 {
				m.cursor++
			}
		case "d":
			if m.cursor < len(m.urls) {
				url := m.urls[m.cursor]
				db.DeleteURL(url.Id)
				m.urls = append(m.urls[:m.cursor], m.urls[m.cursor+1:]...)
				if m.cursor > 0 {
					m.cursor--
				}
			}
		case "c", "v", "m":
			if m.cursor < len(m.urls) {
				url := m.urls[m.cursor]
				if err := db.VerifyURL(url.Id); err != nil {
					log.Printf("Failed to mark URL as verified: %v", err)
				}
				m.urls[m.cursor].Verified = 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := ""
	s += fmt.Sprintf("%-1s %-5s %-60s %-9s\n", " ", "ID", "URL", "Verified")
	s += strings.Repeat("-", 77) + "\n"

	for i, choice := range m.urls {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%-1s %-5s %-60s %-9d\n", cursor, choice.Id, choice.Url, choice.Verified)
	}

	return s
}
