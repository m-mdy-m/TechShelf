# techShelf Catalog Schema

`catalog.json` is the machine-readable source of truth.

## Top-level

- `version` (string): schema version.
- `meta.name` (string): project/library name.
- `meta.description` (string): short human description.
- `meta.last_updated` (string, YYYY-MM-DD): auto-updated by CLI.
- `books` (array): list of book entries.

## Book object

### Required

- `title` (string)
- `author` (array of strings)
- `category` (string)

### Recommended

- `id` (string, slug). If omitted, generated from title.
- `subcategory` (string)
- `year_published` (number)
- `language` (string)
- `tags` (array of strings)
- `level` (`beginner|intermediate|advanced|general`)
- `description` (string)
- `why_read` (string)
- `when_to_read` (string)
- `importance` (string)
- `prerequisites` (array of strings)
- `pairs_well_with` (array of strings)
- `status` (`unread|reading|completed|paused`)
- `notes` (string)
- `added_date` (YYYY-MM-DD)

### Source (at least one is recommended)

- `source.file_path` (string): local relative path in repo.
- `source.online_url` (string): external URL.

## Add book without manual editing

```bash
techshelf add --json '{"title":"Clean Code","author":["Robert C. Martin"],"category":"Computer Science","source":{"online_url":"https://example.com"}}'
```

or

```bash
techshelf add --file ./book.json
```


If you run `techshelf add` without `--json` or `--file`, the CLI opens an interactive prompt mode.

Interactive `add` now asks for all schema-aligned fields (including `year_published`, `when_to_read`, `importance`, `prerequisites`, `pairs_well_with`, and `notes`).
Tags prompt accepts comma-separated tags and also supports space-separated fallback for convenience.