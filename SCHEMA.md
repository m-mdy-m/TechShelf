# TechShelf Catalog Schema

`catalog.json` is the machine-readable source of truth for your library.
The `shelves/` directory contains human-readable markdown summaries auto-generated from it.

---

## Top-level structure

```json
{
  "version": "1.0",
  "meta": {
    "name": "techShelf",
    "description": "...",
    "last_updated": "2026-02-17"
  },
  "books": [ ... ]
}
```

`meta.last_updated` is auto-set by the CLI on every write.

---

## Book object

Only **title**, **author**, and **category** are required.
Everything else is optional — fill in what you know, skip what you don't.

| Field | Type | Notes |
|---|---|---|
| `id` | string (slug) | Auto-generated from title if omitted. |
| `title` | **string** ✱ | |
| `author` | **string[]** ✱ | One or more names. |
| `category` | **string** ✱ | Corresponds to a directory under `shelves/`. |
| `subcategory` | string | Finer grouping within a category. |
| `year_published` | number | Publication year. |
| `language` | string | Default: `English`. |
| `tags` | string[] | Free-form keywords. |
| `level` | enum | `beginner` · `intermediate` · `advanced` · `general` |
| `description` | string | What the book covers. |
| `why_read` | string | Why it's worth your time. |
| `when_to_read` | string | Best moment in your learning path to pick it up. |
| `importance` | string | How significant this book is in its field. |
| `prerequisites` | string[] | Book IDs or topics to know first. |
| `pairs_well_with` | string[] | Book IDs that complement this one. |
| `source.online_url` | string | URL to the book online (PDF, page, etc.). |
| `source.file_path` | string | Relative path to a local file in the repo. |
| `status` | enum | `unread` · `reading` · `completed` · `paused` |
| `notes` | string | Personal notes. |
| `added_date` | string (YYYY-MM-DD) | Auto-set by CLI. |

> ✱ Required.

### Valid enums

**level** — `beginner`, `intermediate`, `advanced`, `general`  
Aliases accepted at input: `intro`→`beginner`, `easy`→`beginner`, `hard`→`advanced`, `all`→`general`

**status** — `unread`, `reading`, `completed`, `paused`  
Aliases accepted at input: `read`/`done`→`completed`, `in-progress`→`reading`, `todo`/`want`→`unread`

---

## Single source of truth

All valid values and defaults are defined in `internal/command/schema.go`.
Changing a valid value or default in that file propagates everywhere automatically — no need to touch commands, types, or prompts separately.

---

## Adding books

```bash
# Interactive (recommended)
shelf add

# Inline JSON
shelf add --json '{"title":"Clean Code","author":["Robert C. Martin"],"category":"computer-science"}'

# From file
shelf add --file ./book.json
```

The `add` command automatically:
1. Adds the entry to `catalog.json`
2. Updates `shelves/<category>/README.md`
3. Creates the category directory if it's new

---

## Shelf files

`shelves/<category>/README.md` is rebuilt automatically by `add`, `remove`, and `sync`.

To manually rebuild all shelves:

```bash
shelf sync
```

---

## Minimal valid book

```json
{
  "title": "The Pragmatic Programmer",
  "author": ["David Thomas", "Andrew Hunt"],
  "category": "computer-science"
}
```