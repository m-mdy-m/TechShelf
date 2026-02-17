package command

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Catalog struct {
	Version string      `json:"version"`
	Meta    CatalogMeta `json:"meta"`
	Books   []Book      `json:"books"`
}

type CatalogMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	LastUpdated string `json:"last_updated"`
}

type Book struct {
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	Author        []string `json:"author"`
	Category      string   `json:"category"`
	Subcategory   string   `json:"subcategory,omitempty"`
	YearPublished *int     `json:"year_published,omitempty"`
	Language      string   `json:"language,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Level         string   `json:"level,omitempty"`
	Description   string   `json:"description,omitempty"`
	WhyRead       string   `json:"why_read,omitempty"`
	WhenToRead    string   `json:"when_to_read,omitempty"`
	Importance    string   `json:"importance,omitempty"`
	Prerequisites []string `json:"prerequisites,omitempty"`
	PairsWellWith []string `json:"pairs_well_with,omitempty"`
	Source        Source   `json:"source,omitempty"`
	Status        string   `json:"status,omitempty"`
	Notes         string   `json:"notes,omitempty"`
	AddedDate     string   `json:"added_date,omitempty"`
}

type Source struct {
	FilePath  *string `json:"file_path,omitempty"`
	OnlineURL *string `json:"online_url,omitempty"`
}

func Load(path string) (*Catalog, error) {
	b, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	var c Catalog
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, err
	}
	if c.Books == nil {
		c.Books = []Book{}
	}
	return &c, nil
}

func (c *Catalog) Save(path string) error {
	c.Meta.LastUpdated = time.Now().Format("2006-01-02")
	sort.Slice(c.Books, func(i, j int) bool { return c.Books[i].ID < c.Books[j].ID })
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Clean(path), append(b, '\n'), 0o644)
}

func (c *Catalog) AddBook(book Book) error {
	if strings.TrimSpace(book.Title) == "" || len(book.Author) == 0 || strings.TrimSpace(book.Category) == "" {
		return errors.New("book requires title, author, and category")
	}
	if book.ID == "" {
		book.ID = c.UniqueID(Slugify(book.Title))
	}
	if c.IDExists(book.ID) {
		return errors.New("book id already exists: " + book.ID)
	}

	book.Level = NormalizeLevel(book.Level)
	if book.Level == "" {
		book.Level = DefaultLevel
	}
	if !ContainsStringFold(Levels, book.Level) {
		return errors.New("invalid level: " + book.Level)
	}

	book.Status = NormalizeStatus(book.Status)
	if book.Status == "" {
		book.Status = DefaultStatus
	}
	if !ContainsStringFold(Statuses, book.Status) {
		return errors.New("invalid status: " + book.Status)
	}

	if book.Language == "" {
		book.Language = DefaultLanguage
	}
	if book.AddedDate == "" {
		book.AddedDate = time.Now().Format("2006-01-02")
	}

	c.Books = append(c.Books, book)
	return nil
}

func (c *Catalog) RemoveBook(id string) (Book, bool) {
	for i, b := range c.Books {
		if b.ID == id {
			c.Books = append(c.Books[:i], c.Books[i+1:]...)
			return b, true
		}
	}
	return Book{}, false
}

func (c *Catalog) IDExists(id string) bool {
	for _, b := range c.Books {
		if b.ID == id {
			return true
		}
	}
	return false
}

func (c *Catalog) UniqueID(base string) string {
	if !c.IDExists(base) {
		return base
	}
	for i := 2; ; i++ {
		if cand := base + "-" + strconv.Itoa(i); !c.IDExists(cand) {
			return cand
		}
	}
}

func (c *Catalog) BooksInCategory(category string) []Book {
	var out []Book
	for _, b := range c.Books {
		if strings.EqualFold(b.Category, category) {
			out = append(out, b)
		}
	}
	return out
}

func Slugify(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	r := strings.NewReplacer(" ", "-", "_", "-", "/", "-", ".", "", ",", "", ":", "", "'", "")
	s = r.Replace(s)
	for strings.Contains(s, "--") {
		s = strings.ReplaceAll(s, "--", "-")
	}
	return strings.Trim(s, "-")
}

func strPtr(v string) *string { return &v }
