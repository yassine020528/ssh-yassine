package pages

import (
	"fmt"
	"strings"

	"ssh-yassine/internal/view"
)

type Experience struct {
	Role    string
	Company string
	Period  string
	Desc    string
}

var experiences = []Experience{
	{
		Role:    "Software Developer - Capstone Project",
		Company: "Hydro-Québec",
		Period:  "Jan 2026 - May 2026",
		Desc:    "Built a full-stack weather platform (Angular / ASP.NET Core) implementing GDAL raster processing and Docker containerization.",
	},
	{
		Role:    "Software Quality Assurance Intern",
		Company: "UpToTest",
		Period:  "May 2024 - Sep 2024",
		Desc:    "Designed and automated end-to-end functional test suites for the KOORS web app using Cypress with Cucumber/Gherkin in a BDD workflow.",
	},
	{
		Role:    "Private Tutor",
		Company: "Self-employed",
		Period:  "Aug 2021 - Present",
		Desc:    "Tutoring Mathematics, Physics, and Programming for college and university students.",
	},
}

func Experiences() []Experience {
	return experiences
}

func RenderExperience(styles view.ThemeStyles, expCursor int, themeLabel string) string {
	var b strings.Builder

	b.WriteString(styles.Title.Render("━━━ Experience ━━━"))
	b.WriteString("\n\n")

	for i, exp := range experiences {
		cursor := "  "
		if expCursor == i {
			cursor = "→ "
		}

		line := fmt.Sprintf("%s%s @ %s",
			cursor,
			styles.Role.Render(exp.Role),
			styles.Company.Render(exp.Company))
		b.WriteString(line)
		b.WriteString("\n")

		b.WriteString("    ")
		b.WriteString(styles.Period.Render(exp.Period))
		b.WriteString("\n")

		if expCursor == i {
			b.WriteString("    ")
			b.WriteString(styles.Content.Render(exp.Desc))
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}

	help := themeLabel + " • ↑/↓: browse • esc: back to menu"
	b.WriteString(styles.Help.Render(help))

	return b.String()
}
