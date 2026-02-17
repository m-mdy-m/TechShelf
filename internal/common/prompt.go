package common

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Prompter struct {
	r *bufio.Reader
}

func NewPrompter() *Prompter {
	return &Prompter{r: bufio.NewReader(os.Stdin)}
}

func (p *Prompter) Ask(label, defaultValue string) (string, error) {
	if defaultValue != "" {
		fmt.Printf("? %s [%s]: ", label, defaultValue)
	} else {
		fmt.Printf("? %s: ", label)
	}
	line, err := p.r.ReadString('\n')
	if err != nil {
		return "", err
	}
	line = strings.TrimSpace(line)
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
		fmt.Println("  ! This field is required.")
	}
}

func (p *Prompter) AskList(label string) ([]string, error) {
	raw, err := p.Ask(label+" (comma-separated)", "")
	if err != nil {
		return nil, err
	}
	if raw == "" {
		return []string{}, nil
	}
	parts := strings.Split(raw, ",")
	res := make([]string, 0, len(parts))
	for _, item := range parts {
		item = strings.TrimSpace(item)
		if item != "" {
			res = append(res, item)
		}
	}
	return res, nil
}

func (p *Prompter) AskTags(label string) ([]string, error) {
	raw, err := p.Ask(label+" (comma-separated; spaces also accepted)", "")
	if err != nil {
		return nil, err
	}
	if raw == "" {
		return []string{}, nil
	}
	parts := []string{}
	if strings.Contains(raw, ",") {
		parts = strings.Split(raw, ",")
	} else {
		parts = strings.Fields(raw)
	}
	res := make([]string, 0, len(parts))
	for _, item := range parts {
		item = strings.TrimSpace(item)
		if item != "" {
			res = append(res, item)
		}
	}
	return res, nil
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

func (p *Prompter) AskChoice(label string, allowed []string, defaultValue string) (string, error) {
	for {
		value, err := p.Ask(label+" ("+strings.Join(allowed, "/")+")", defaultValue)
		if err != nil {
			return "", err
		}
		for _, item := range allowed {
			if strings.EqualFold(value, item) {
				return item, nil
			}
		}
		fmt.Println("  ! Invalid choice.")
	}
}

func (p *Prompter) AskIntOptional(label string) (*int, error) {
	for {
		value, err := p.Ask(label, "")
		if err != nil {
			return nil, err
		}
		if value == "" {
			return nil, nil
		}
		n, convErr := strconv.Atoi(value)
		if convErr != nil {
			fmt.Println("  ! Please enter a valid integer year.")
			continue
		}
		return &n, nil
	}
}
