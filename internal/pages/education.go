package pages

import (
	"fmt"
	"strings"

	"ssh-yassine/internal/view"
)

type Education struct {
	Role    string
	Company string
	Period  string
	Desc    string
	URL     string
}

var educations = []Education{
	{
		Role:    "Computer Engineering, B.E.",
		Company: "Polytechnique Montréal",
		Period:  "Aug 2022 - May 2026",
		URL:     "https://www.polymtl.ca/",
	},
}

func Educations() []Education {
	return educations
}

func RenderEducation(styles view.ThemeStyles, eduCursor int, themeLabel string) string {
	var b strings.Builder

	b.WriteString(styles.Title.Render("━━━ Education ━━━"))
	b.WriteString("\n\n")

	const pageSize = 3
	start, end := educationWindow(eduCursor, len(educations), pageSize)

	for i := start; i < end; i++ {
		edu := educations[i]
		cursor := "  "
		if eduCursor == i {
			cursor = "→ "
		}

		line := fmt.Sprintf("%s%s @ %s",
			cursor,
			styles.Role.Render(edu.Role),
			styles.Company.Render(edu.Company))
		b.WriteString(line)
		b.WriteString("\n")

		b.WriteString("    ")
		b.WriteString(styles.Period.Render(edu.Period))
		b.WriteString("\n")

		if eduCursor == i {
			desc := edu.Desc
			if edu.URL != "" {
				desc = view.ClickableLink(desc, edu.URL)
			}
			b.WriteString("    ")
			b.WriteString(styles.Content.Render(desc))
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}

	help := themeLabel + " • ↑/↓: browse • esc: back to menu"
	if end < len(educations) {
		moreStyle := styles.Accent.Copy().Faint(true)
		b.WriteString(moreStyle.Render("more below!"))
		b.WriteString(styles.Help.Render(" • " + help))
	} else {
		b.WriteString(styles.Help.Render(help))
	}

	return b.String()
}

func educationWindow(cursor, total, pageSize int) (start, end int) {
	if total <= 0 {
		return 0, 0
	}
	if pageSize <= 0 || total <= pageSize {
		return 0, total
	}
	if cursor < 0 {
		cursor = 0
	}
	if cursor >= total {
		cursor = total - 1
	}
	start = cursor - (pageSize - 1)
	if start < 0 {
		start = 0
	}
	end = start + pageSize
	if end > total {
		end = total
		start = end - pageSize
	}
	return start, end
}
