# CleanSource

CleanSource scans your project and reports every file and folder that is **not currently used** — i.e. never imported or referenced from your entry point. It helps you find dead files, stale assets, and leftover code that's safe to remove.

CleanSource is designed to be **language-agnostic**, but it has currently only been tested against **JavaScript** and **Go** projects. Other languages may work, but results aren't guaranteed yet.

## Installation

```bash
go install github.com/yourname/cleansource@latest
```

(Adjust this once you have a real module path / release process.)

## Usage

```bash
cleansource <path>
```

Run it against the root of your project (or any subdirectory you want scanned).

### Example

```bash
cleansource .
```

## Flags

| Flag | Description |
|------|-------------|
| `--start` | The entry file CleanSource should start tracing imports/usage from (e.g. `main.go`, `app/page.ts`). If omitted, CleanSource starts from the first file or folder it finds. |
| `--ignore` | Comma-separated list of files or folders to skip (e.g. `--ignore node_modules,package.json,go.mod`). |
| `--list` | Only print the paths that are unused — don't delete anything. |

### Examples

Start tracing from a specific entry point:

```bash
cleansource . --start main.go
```

Ignore common non-source folders/files:

```bash
cleansource . --ignore node_modules,package.json,go.mod
```

Just see what's unused, without deleting anything:

```bash
cleansource . --list
```

## Ignoring files with `.csignore`

Instead of passing `--ignore` every time, you can drop a `.csignore` file in your project root. It works the same way a `.gitignore` does — one pattern per line.

```
node_modules
package.json
go.mod
*.log
dist/
```

If both a `.csignore` file and the `--ignore` flag are present, their entries are combined.

## Ignore pattern matching

Ignore entries (from either `--ignore` or `.csignore`) support simple wildcard matching:

| Pattern | Matches |
|---------|---------|
| `gitignore` | Exact filename match |
| `*.git` | Anything ending in `.git` |
| `git*` | Anything starting with `git` |
| `*git*` | Anything containing `git` |

## How it works

CleanSource walks your project directory, builds a reference graph starting from `--start` (or the first file found if not specified), and flags any file that is never imported, required, or otherwise referenced anywhere in that graph. Files and folders matching your ignore rules are skipped entirely and never flagged.

By default, CleanSource **deletes** unused files it finds. Pass `--list` if you just want a report first — recommended before running it destructively on a project you care about.

## Supported platforms

- ✅ Go
- ✅ JavaScript
- ⚠️ Other languages: untested, use with caution

## Caution

Always run with `--list` first and review the output before letting CleanSource delete anything. Back up or commit your work beforehand.