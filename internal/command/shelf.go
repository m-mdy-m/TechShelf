package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/m-mdy-m/TechShelf/internal/common"
)

func SyncShelfCategory(shelvesDir, category string, books []Book) error {
	slug := Slugify(category)
	dir := filepath.Join(shelvesDir, slug)
	if err := common.EnsureDir(dir); err != nil {
		return err
	}
	var sb strings.Builder
	sb.WriteString("# " + category + "\n\n")
	for _, b := range books {
		writeBookMD(&sb, b)
	}
	return os.WriteFile(filepath.Join(dir, "README.md"), []byte(sb.String()), 0o644)
}

func SyncAllShelves(shelvesDir string, books []Book) error {
	groups := map[string][]Book{}
	for _, b := range books {
		groups[b.Category] = append(groups[b.Category], b)
	}
	for category, grp := range groups {
		if err := SyncShelfCategory(shelvesDir, category, grp); err != nil {
			return err
		}
	}
	return nil
}

func writeBookMD(sb *strings.Builder, b Book) {
	sb.WriteString(fmt.Sprintf("## %s\n\n", b.Title))
	sb.WriteString(fmt.Sprintf("- **Authors**: %s\n", strings.Join(b.Author, ", ")))
	if b.YearPublished != nil {
		sb.WriteString(fmt.Sprintf("- **Year**: %d\n", *b.YearPublished))
	}
	if b.Language != "" && b.Language != DefaultLanguage {
		sb.WriteString(fmt.Sprintf("- **Language**: %s\n", b.Language))
	}
	sb.WriteString(fmt.Sprintf("- **Level**: %s\n", b.Level))
	sb.WriteString(fmt.Sprintf("- **Status**: %s\n", b.Status))
	if len(b.Tags) > 0 {
		sb.WriteString(fmt.Sprintf("- **Tags**: %s\n", strings.Join(b.Tags, ", ")))
	}
	if b.Subcategory != "" {
		sb.WriteString(fmt.Sprintf("- **Subcategory**: %s\n", b.Subcategory))
	}
	if b.Description != "" {
		sb.WriteString(fmt.Sprintf("\n> %s\n", b.Description))
	}
	if b.WhyRead != "" {
		sb.WriteString(fmt.Sprintf("\n**Why read**: %s\n", b.WhyRead))
	}
	if b.WhenToRead != "" {
		sb.WriteString(fmt.Sprintf("\n**When to read**: %s\n", b.WhenToRead))
	}
	if b.Importance != "" {
		sb.WriteString(fmt.Sprintf("\n**Importance**: %s\n", b.Importance))
	}
	if len(b.Prerequisites) > 0 {
		sb.WriteString(fmt.Sprintf("\n**Prerequisites**: %s\n", strings.Join(b.Prerequisites, ", ")))
	}
	if len(b.PairsWellWith) > 0 {
		sb.WriteString(fmt.Sprintf("**Pairs well with**: %s\n", strings.Join(b.PairsWellWith, ", ")))
	}
	if b.Source.OnlineURL != nil {
		sb.WriteString(fmt.Sprintf("\n🔗 [Online](%s)\n", *b.Source.OnlineURL))
	}
	if b.Source.FilePath != nil {
		sb.WriteString(fmt.Sprintf("📄 [Local file](%s)\n", *b.Source.FilePath))
	}
	if b.Notes != "" {
		sb.WriteString(fmt.Sprintf("\n_Notes: %s_\n", b.Notes))
	}
	sb.WriteString("\n---\n\n")
}
