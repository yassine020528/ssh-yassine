package pages

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"ssh-yassine/internal/view"
)

func RenderContact(styles view.ThemeStyles, themeLabel string) string {
	var b strings.Builder

	b.WriteString(styles.Title.Render("━━━ Contact ━━━"))
	b.WriteString("\n")

	b.WriteString(styles.Content.Render("Feel free to reach out!"))
	b.WriteString("\n\n")

	githubURL := "https://github.com/yassine020528"
	mailtoURL := "mailto:yassine020528@gmail.com"
	linkedinURL := "https://www.linkedin.com/in/yassine-abassi-b9ba721a6/"
	portfolioURL := "https://yassineabassi.com"

	leftLines := []string{
		styles.Accent.Render("Contacts"),
		styles.Content.Render(fmt.Sprintf("Email      %s", view.ClickableLink("yassine020528@gmail.com", mailtoURL))),
		styles.Content.Render(fmt.Sprintf("LinkedIn   %s", view.ClickableLink("Yassine Abassi", linkedinURL))),
		styles.Content.Render(fmt.Sprintf("GitHub     %s", view.ClickableLink("@yassine020528", githubURL))),
	}

	rightLines := []string{
		styles.Accent.Render("Websites"),
		styles.Content.Render(fmt.Sprintf("%s", view.ClickableLink("yassineabassi.com", portfolioURL))),
		"",
	}

	leftColumn := lipgloss.NewStyle().Width(40).Render(strings.Join(leftLines, "\n"))
	rightColumn := lipgloss.NewStyle().Width(24).Render(strings.Join(rightLines, "\n"))
	columns := lipgloss.JoinHorizontal(lipgloss.Top, leftColumn, rightColumn)

	b.WriteString(columns)

	b.WriteString("\n")
	b.WriteString(styles.Help.Render(themeLabel + " • esc: back to menu"))

	return b.String()
}
