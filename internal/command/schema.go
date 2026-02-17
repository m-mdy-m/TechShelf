package command

import (
	"strings"
)

var Levels = []string{"beginner", "intermediate", "advanced", "general"}

var Statuses = []string{"unread", "reading", "completed", "paused"}

const (
	DefaultLevel    = "general"
	DefaultStatus   = "unread"
	DefaultLanguage = "English"
	ShelvesDir      = "shelves"
	CatalogVersion  = "1.0"
)

var statusAliases = map[string]string{
	"read":        "completed",
	"done":        "completed",
	"in-progress": "reading",
	"in_progress": "reading",
	"todo":        "unread",
	"want":        "unread",
}

var levelAliases = map[string]string{
	"intro": "beginner",
	"easy":  "beginner",
	"hard":  "advanced",
	"all":   "general",
}

func NormalizeStatus(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	if alias, ok := statusAliases[s]; ok {
		return alias
	}
	return s
}

func NormalizeLevel(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	if alias, ok := levelAliases[s]; ok {
		return alias
	}
	return s
}

func ContainsStringFold(values []string, target string) bool {
	for _, v := range values {
		if strings.EqualFold(v, target) {
			return true
		}
	}
	return false
}
