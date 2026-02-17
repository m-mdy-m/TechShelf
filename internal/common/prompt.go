package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var (
	clrPrompt  = color.New(color.FgCyan, color.Bold)
	clrOption  = color.New(color.FgWhite)
	clrHint    = color.New(color.FgHiBlack)
	clrWarn    = color.New(color.FgYellow)
	clrSuccess = color.New(color.FgGreen)
	clrLabel   = color.New(color.FgHiWhite, color.Bold)
)

type Prompter struct {
	r *bufio.Reader
}

func NewPrompter() *Prompter {
	return &Prompter{r: bufio.NewReader(os.Stdin)}
}

func (p *Prompter) readLine() (string, error) {
	line, err := p.r.ReadString('\n')
	return strings.TrimSpace(line), err
}

func (p *Prompter) Ask(label, defaultValue string) (string, error) {
	if defaultValue != "" {
		clrPrompt.Printf("  ❯ %s ", label)
		clrHint.Printf("[%s]", defaultValue)
		fmt.Print(": ")
	} else {
		clrPrompt.Printf("  ❯ %s: ", label)
	}
	line, err := p.readLine()
	if err != nil {
		return "", err
	}
	if line == "" {
		return defaultValue, nil
	}
	return line, nil
}

func (p *Prompter) AskRequired(label string) (string, error) {
	for {
		v, err := p.Ask(label, "")
		if err != nil {
			return "", err
		}
		if strings.TrimSpace(v) != "" {
			return v, nil
		}
		clrWarn.Println("    ! This field is required.")
	}
}

func (p *Prompter) AskList(label string) ([]string, error) {
	raw, err := p.Ask(label+" (comma-separated)", "")
	if err != nil {
		return nil, err
	}
	return splitTrimmed(raw, ","), nil
}

func (p *Prompter) AskTags(label string) ([]string, error) {
	raw, err := p.Ask(label+" (comma or space separated)", "")
	if err != nil {
		return nil, err
	}
	if raw == "" {
		return []string{}, nil
	}
	if strings.Contains(raw, ",") {
		return splitTrimmed(raw, ","), nil
	}
	return strings.Fields(raw), nil
}

func (p *Prompter) AskConfirm(label string, defaultYes bool) (bool, error) {
	def := "y"
	if !defaultYes {
		def = "n"
	}
	v, err := p.Ask(label+" [y/n]", def)
	if err != nil {
		return false, err
	}
	v = strings.ToLower(strings.TrimSpace(v))
	return v == "y" || v == "yes", nil
}

func (p *Prompter) AskIntOptional(label string) (*int, error) {
	for {
		v, err := p.Ask(label+" (optional)", "")
		if err != nil {
			return nil, err
		}
		if v == "" {
			return nil, nil
		}
		n, convErr := strconv.Atoi(v)
		if convErr != nil {
			clrWarn.Println("    ! Please enter a valid integer.")
			continue
		}
		return &n, nil
	}
}

func (p *Prompter) AskChoice(label string, allowed []string, defaultValue string) (string, error) {
	for {
		v, err := p.Ask(label+" ("+strings.Join(allowed, "/")+")", defaultValue)
		if err != nil {
			return "", err
		}
		for _, item := range allowed {
			if strings.EqualFold(v, item) {
				return item, nil
			}
		}
		clrWarn.Printf("    ! Must be one of: %s\n", strings.Join(allowed, ", "))
	}
}

func (p *Prompter) SelectOrNew(label string, options []string) (string, error) {
	if len(options) > 0 {
		fmt.Println()
		clrLabel.Printf("  %s:\n", label)
		for i, opt := range options {
			clrHint.Printf("    %d) ", i+1)
			clrOption.Println(opt)
		}
		fmt.Println()
	}

	hint := "enter number or new value"
	if len(options) == 0 {
		hint = "enter value"
	}

	for {
		clrPrompt.Printf("  ❯ %s ", label)
		clrHint.Printf("(%s)", hint)
		fmt.Print(": ")
		line, err := p.readLine()
		if err != nil {
			return "", err
		}
		line = strings.TrimSpace(line)
		if line == "" {
			clrWarn.Println("    ! This field is required.")
			continue
		}

		if n, err := strconv.Atoi(line); err == nil {
			if n >= 1 && n <= len(options) {
				clrSuccess.Printf("    ✔ %s\n", options[n-1])
				return options[n-1], nil
			}
			clrWarn.Printf("    ! Enter 1–%d or type a new value.\n", len(options))
			continue
		}
		return line, nil
	}
}

func Section(title string) {
	fmt.Println()
	color.New(color.FgMagenta, color.Bold).Printf("  ── %s ──\n", title)
	fmt.Println()
}

func splitTrimmed(s, sep string) []string {
	parts := strings.Split(s, sep)
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if v := strings.TrimSpace(p); v != "" {
			out = append(out, v)
		}
	}
	return out
}
