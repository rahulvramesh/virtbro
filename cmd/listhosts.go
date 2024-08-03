package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"virtbro/pkg/db"
)

var listHostsCmd = &cobra.Command{
	Use:   "listhosts",
	Short: "List all remote hosts",
	Long:  `List all remote hosts added to the local database.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := listHosts(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(listHostsCmd)
}

func listHosts() error {
	hosts, err := db.ListHosts()
	if err != nil {
		return fmt.Errorf("failed to list hosts: %v", err)
	}

	columns := []table.Column{
		{Title: "ID", Width: 5},
		{Title: "Name", Width: 20},
		{Title: "URI", Width: 30},
		{Title: "UUID", Width: 36},
	}

	var rows []table.Row
	for _, host := range hosts {
		rows = append(rows, table.Row{
			host["id"],
			host["name"],
			host["uri"],
			host["uuid"],
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	t.SetStyles(table.Styles{
		Header: table.DefaultStyles().Header.
			Bold(true).
			Foreground(lipgloss.Color("205")),
		Cell: table.DefaultStyles().Cell.
			Foreground(lipgloss.Color("240")),
		Selected: table.DefaultStyles().Selected.
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("57")),
	})

	p := tea.NewProgram(model{t: t})

	if err := p.Start(); err != nil {
		return fmt.Errorf("failed to start program: %v", err)
	}

	return nil
}

type model struct {
	t table.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key.Matches(msg, quitKey) {
		case true:
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.t, cmd = m.t.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return lipgloss.NewStyle().Padding(1, 2, 1, 2).Render(m.t.View())
}

var quitKey = key.NewBinding(
	key.WithKeys("q"),
	key.WithHelp("q", "quit"),
)
