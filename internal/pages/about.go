package pages

import (
	"math"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"ssh-yassine/internal/view"
)

const aboutLogo = `
           .......           
          ....:++..          
          .:xXXXx+.          
          +++xxxxxx          
           +++;;x+x          
           ..++x;.           
            x;.;xx.          
         ..;$XxX&x:..        
     .;;::::++:+$;::::..     
    ..$Xx;::..::.::::::...   
   ...:x$x+:.....:;;;;;....  
  ....:xxxX$.....;+xx+;..... 
 .....+XXX+;;;;:.;+xx+:......
......;x;..:x&&..:;;;;:......
......:+...:.+$$+:;;;:.......
..........:....$XXx+.........
................+xXXX;.......
..........::.....;x+.........`

const aboutIntro = `Hey there, I'm Yassine Abassi, a Computer Engineering graduate from Polytechnique Montreal. I've been fascinated by how things work under the hood ever since I was a child.
My younger self's curiosity has blossomed into a full-blown passion for Computer Engineering.
"Why build a simple portfolio when you can build one running on a terminal?" - Me, at 3 AM.`

type aboutIntroStyle int

const (
	aboutIntroNormal aboutIntroStyle = iota
	aboutIntroBold
	aboutIntroItalic
)

type aboutIntroSegment struct {
	text  string
	style aboutIntroStyle
}

var aboutIntroSegments = []aboutIntroSegment{
	{text: "Hey there, I'm ", style: aboutIntroNormal},
	{text: "Yassine Abassi", style: aboutIntroBold},
	{text: ", a ", style: aboutIntroNormal},
	{text: "Computer Engineering", style: aboutIntroItalic},
	{text: " graduate from ", style: aboutIntroNormal},
	{text: "Polytechnique Montreal", style: aboutIntroItalic},
	{text: ". I've been fascinated by how things work ", style: aboutIntroNormal},
	{text: "under the hood", style: aboutIntroItalic},
	{text: " ever since I was a child.\nMy younger self's curiosity has blossomed into a full-blown passion for ", style: aboutIntroNormal},
	{text: "Computer Engineering", style: aboutIntroItalic},
	{text: ".\n", style: aboutIntroNormal},
	{text: "\"Why build a simple portfolio when you can build one running on a terminal?\"", style: aboutIntroItalic},
	{text: " - Me, at 3 AM.", style: aboutIntroNormal},
}

var aboutLogoRunes = []rune(aboutLogo)
var aboutIntroRunes = []rune(aboutIntro)

var scrambleRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:'\",.<>?/~")

const settleDurationTicks = 8

func aboutSettled(count int, scrambleTick int) bool {
	total := AboutRuneCount()
	if count > total {
		count = total
	}
	return count >= total && scrambleTick >= count+AboutSettleTicks()
}

func aboutStyled(styles view.ThemeStyles, contentWidth int) string {
	var b strings.Builder

	b.WriteString(centerAboutLogo(styles.Accent.Copy().Bold(true).Render(aboutLogo), contentWidth))
	b.WriteString("\n\n")
	b.WriteString(aboutIntroStyled(styles, len(aboutIntroRunes), 1_000_000))

	return b.String()
}

func aboutVisibleStyled(styles view.ThemeStyles, logoVisible string, introVisible string, contentWidth int) string {
	var b strings.Builder

	if logoVisible != "" {
		b.WriteString(centerAboutLogo(styles.Accent.Copy().Bold(true).Render(logoVisible), contentWidth))
	}
	if introVisible != "" {
		b.WriteString("\n")
		b.WriteString(introVisible)
	}

	return b.String()
}

func aboutIntroStyled(styles view.ThemeStyles, revealCount int, scrambleTick int) string {
	if revealCount <= 0 {
		return ""
	}

	var b strings.Builder
	consumed := 0
	for _, segment := range aboutIntroSegments {
		segmentRunes := []rune(segment.text)
		segmentCount := clamp(revealCount-consumed, 0, len(segmentRunes))
		if segmentCount <= 0 {
			break
		}

		visible := aboutVisible(segmentRunes, segmentCount, scrambleTick, consumed)
		b.WriteString(aboutIntroSegmentStyle(styles, segment.style).Render(visible))
		consumed += len(segmentRunes)
	}

	return b.String()
}

func aboutIntroSegmentStyle(styles view.ThemeStyles, style aboutIntroStyle) lipgloss.Style {
	switch style {
	case aboutIntroBold:
		return styles.Content.Copy().Bold(true)
	case aboutIntroItalic:
		return styles.Content.Copy().Italic(true)
	default:
		return styles.Content
	}
}

func centerAboutLogo(logo string, contentWidth int) string {
	if contentWidth <= 0 {
		return logo
	}
	return lipgloss.NewStyle().Width(contentWidth).Align(lipgloss.Center).Render(logo)
}

func aboutContentWidth(boxWidth int) int {
	if boxWidth <= 0 {
		return 60
	}
	if boxWidth > 4 {
		return boxWidth - 4
	}
	return boxWidth
}

func aboutVisible(runes []rune, count int, scrambleTick int, indexOffset int) string {
	if count <= 0 {
		return ""
	}
	total := len(runes)
	if count > total {
		count = total
	}

	out := make([]rune, count)
	i := 0
	for i < count {
		ch := runes[i]
		if isWhitespace(ch) {
			out[i] = ch
			i++
			continue
		}

		wordStart := i
		wordEnd := i
		for wordEnd < total && !isWhitespace(runes[wordEnd]) {
			wordEnd++
		}

		visibleEnd := wordEnd
		if visibleEnd > count {
			visibleEnd = count
		}

		if count < wordEnd {
			for j := wordStart; j < visibleEnd; j++ {
				out[j] = scrambleRune(scrambleTick, indexOffset+j)
			}
		} else {
			ticksSinceComplete := scrambleTick - indexOffset - wordEnd
			if ticksSinceComplete < 0 {
				ticksSinceComplete = 0
			}
			wordLen := wordEnd - wordStart
			settled := settledChars(wordLen, ticksSinceComplete)
			for j := wordStart; j < visibleEnd; j++ {
				if j-wordStart < settled {
					out[j] = runes[j]
					continue
				}
				out[j] = scrambleRune(scrambleTick, indexOffset+j)
			}
		}

		i = visibleEnd
	}

	return string(out)
}

func AboutRuneCount() int {
	return max(len(aboutLogoRunes), len(aboutIntroRunes))
}

func AboutSettleTicks() int {
	return settleDurationForWord(lastWordLength())
}

func RenderAbout(styles view.ThemeStyles, revealCount int, scrambleTick int, themeLabel string, boxWidth int) string {
	var b strings.Builder

	b.WriteString(styles.Title.Render("━━━ About Me ━━━"))
	b.WriteString("\n")
	contentWidth := aboutContentWidth(boxWidth)
	if aboutSettled(revealCount, scrambleTick) {
		b.WriteString(aboutStyled(styles, contentWidth))
	} else {
		logoVisible := aboutVisible(aboutLogoRunes, revealCount, scrambleTick, 0)
		introVisible := aboutIntroStyled(styles, revealCount, scrambleTick)
		b.WriteString(aboutVisibleStyled(styles, logoVisible, introVisible, contentWidth))
	}
	b.WriteString("\n")
	b.WriteString(styles.Help.Render(themeLabel + " • esc: back to menu"))

	return b.String()
}

func scrambleRune(scrambleTick int, index int) rune {
	if len(scrambleRunes) == 0 {
		return ' '
	}
	value := uint64(scrambleTick)*1103515245 + uint64(index)*12345 + 12345
	return scrambleRunes[int(value%uint64(len(scrambleRunes)))]
}

func isWhitespace(ch rune) bool {
	switch ch {
	case ' ', '\n', '\r', '\t':
		return true
	default:
		return false
	}
}

func settledChars(wordLen int, ticksSinceComplete int) int {
	if wordLen <= 0 {
		return 0
	}
	durationTicks := settleDurationForWord(wordLen)
	if durationTicks <= 0 {
		return 0
	}
	if ticksSinceComplete >= durationTicks {
		return wordLen
	}
	if ticksSinceComplete <= 0 {
		return 0
	}
	step := float64(durationTicks) / float64(wordLen)
	settled := int(math.Floor(float64(ticksSinceComplete) / step))
	if settled > wordLen {
		return wordLen
	}
	if settled < 0 {
		return 0
	}
	return settled
}

func settleDurationForWord(wordLen int) int {
	if wordLen <= 0 {
		return 0
	}
	if settleDurationTicks < wordLen {
		return wordLen
	}
	return settleDurationTicks
}

func lastWordLength() int {
	if len(aboutIntroRunes) == 0 {
		return 0
	}
	i := len(aboutIntroRunes) - 1
	for i >= 0 && isWhitespace(aboutIntroRunes[i]) {
		i--
	}
	length := 0
	for i >= 0 && !isWhitespace(aboutIntroRunes[i]) {
		length++
		i--
	}
	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func clamp(value int, low int, high int) int {
	if value < low {
		return low
	}
	if value > high {
		return high
	}
	return value
}
