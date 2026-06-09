package pages

import (
	"strings"

	"ssh-yassine/internal/view"
)

type Project struct {
	Name string
	Desc string
	Tech string
	Link string
}

var projects = []Project{
	{
		Name: "Hydro-Quebec Weather Forecasting Platform",
		Desc: "Full-stack weather visualization platform with interactive geospatial mapping and real-time updates.",
		Tech: "Angular, .NET, EF Core, GDAL, Leaflet, SignalR, Docker",
		Link: "",
	},
	{
		Name: "3D Interactive Portfolio",
		Desc: "3D portfolio with an interactive room featuring a functional OS, custom camera rigging, and state-driven transitions.",
		Tech: "React.js, Three.js, Upstash Redis",
		Link: "https://yassineabassi.com",
	},
	{
		Name: "Terminal Portfolio",
		Desc: "SSH-based interactive portfolio application served over SSH, built with Go and Charm, deployed on Fly.io.",
		Tech: "Go, Charm, Fly.io",
		Link: "https://yassineabassi.com/ssh",
	},
	{
		Name: "Corporate LinkedIn Translator ",
		Desc: "Chrome extension that translates LinkedIn buzzwords into plain satirical language.",
		Tech: "Node.js/Express, Fly.io",
		Link: "https://github.com/yassine020528/corporate-translator",
	},
	{
		Name: "WildGuard - AI Wildlife Monitoring System",
		Desc: "AI-powered ecological surveillance platform in Python for anti-poaching and endangered species protection.",
		Tech: "Python, Socket.IO, Twilio, YOLOv8, Flask",
		Link: "https://github.com/yassine020528/wildguard",
	},
	{
		Name: "Lost Woods Browser Horror Game",
		Desc: "Canvas-based horror game with tile generation, reachable item placement using pathfinding logic and enemy AI behaviour.",
		Tech: "React.js, Typescript, Web Audio API",
		Link: "https://lost-woods.netlify.app",
	},
	{
		Name: "Multi-Robot Exploration System",
		Desc: "Coordinated multi-robot exploration system with AgileX Limo robots.",
		Tech: "Python, ROS2, Gazebo, Docker",
		Link: "",
	},
	{
		Name: "Real-Time Multiplayer Combat Game ",
		Desc: "Real-time multiplayer game featuring seamless bidirectional communication.",
		Tech: "Angular, NestJS, WebSockets",
		Link: "",
	},
	{
		Name: "Embedded Obstacle Detection System",
		Desc: "Autonomous obstacle detection and search system with low-level I/O control and real-time decision logic.",
		Tech: "C++, ATmega164A",
		Link: "",
	},
	{
		Name: "Android Mobile Applications (3 Apps)",
		Desc: "Finance tracker, Bluetooth device tracker, and Location-based running-walking app with Retrofit/OkHttp networking, Firebase cloud synchronization.",
		Tech: "Kotlin, Jetpack Compose, Room, Firebase",
		Link: "",
	},
	{
		Name: "Desktop Chess Application ",
		Desc: "Fully interactive chess application implementing legal move validation with a custom QWidget board renderer.",
		Tech: "C++, Qt",
		Link: "https://github.com/yassine020528/chess-cpp-qt",
	},
}

func Projects() []Project {
	return projects
}

func RenderProjects(styles view.ThemeStyles, projectCursor int, themeLabel string) string {
	var b strings.Builder

	b.WriteString(styles.Title.Render("━━━ Projects ━━━"))
	b.WriteString("\n\n")

	const pageSize = 3
	start, end := projectWindow(projectCursor, len(projects), pageSize)

	for i := start; i < end; i++ {
		p := projects[i]
		cursor := "  "
		if projectCursor == i {
			cursor = "→ "
		}

		name := cursor + p.Name
		if projectCursor == i {
			b.WriteString(styles.ProjectName.Render(name))
		} else {
			b.WriteString(styles.Menu.Render(name))
		}
		b.WriteString("\n")

		// Expands project section
		if projectCursor == i {
			b.WriteString(styles.Subtle.Render("    " + p.Desc))
			b.WriteString("\n")
			b.WriteString("    ")
			b.WriteString(styles.Tech.Render(p.Tech))
			b.WriteString("\n")
			projectURL := strings.TrimSpace(p.Link)
			if projectURL != "" {
				if !strings.HasPrefix(projectURL, "http://") && !strings.HasPrefix(projectURL, "https://") {
					projectURL = "https://" + projectURL
				}
				b.WriteString("    ")
				b.WriteString(styles.Accent.Render(view.ClickableLink(projectURL, projectURL)))
				b.WriteString("\n")
			}
		}
		b.WriteString("\n")
	}

	help := themeLabel + " • ↑/↓: browse • esc: back to menu"
	if end < len(projects) {
		moreStyle := styles.Accent.Copy().Faint(true)
		b.WriteString(moreStyle.Render("more below!"))
		b.WriteString(styles.Help.Render(" • " + help))
	} else {
		b.WriteString(styles.Help.Render(help))
	}

	return b.String()
}

func projectWindow(cursor, total, pageSize int) (start, end int) {
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
