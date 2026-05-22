# YAPST — Yet Another Personal Site Template

A minimal personal website template built with [Go](https://go.dev) and [HTMX](https://htmx.org). Write posts in Markdown. There are no build steps, no frameworks, no bundler, and no JavaScript.

To see a live version, click [here](https://tahashafiei.com).

**Features:**

- **Go binary**
- **HTMX navigation**
- Markdown posts with **frontmatter**
- **Dark/light theme toggle** (Kanagawa Dragon / Lotus palette as default)
- **Mobile-responsive**

## Prerequisites

- [Go 1.22 or later](https://go.dev/dl/)

Verify your installation:

```sh
go version
```

## Quick start

```sh
git clone https://github.com/yourusername/yapst my-site
cd my-site
go run .
```

> Change `my-site` to your site's name.
> You can also fork this repo and clone is normally.

Open [http://localhost:8080](http://localhost:8080).

## Customization

### 1. Personal details

Open `config.go` and update these fields:

```go
var siteConfig = SiteConfig{
    Name:        "Your Name",
    GitHubURL:   "https://github.com/yourusername",
    LinkedInURL: "https://www.linkedin.com/in/yourusername",
    ResumeFile:  "resume.pdf", // filename inside static/; leave empty to hide
}
```

These values are populated into the site name, page titles, and footer links automatically. To hide a footer link, set its field to an empty string `""`.

### 2. About page

Edit `templates/home.html`. Given that we are using HTMX, basic HTML is all we need. Make any and all changes between the `{{define "content"}}` and `{{end}}` tags.

### 3. Misc page

Edit `templates/misc.html` the same way. Use it for anything that doesn't belong on the About page, or feel free to change it into whatever page you need.

### 4. Adding a post

Create a new `.md` file in `posts/`:

```
posts/my-post-title.md
```

Every post needs a frontmatter block at the top:

```
---
title: "Post Title"
date: "2026-01-15"
description: "One sentence shown in the post list."
---

Your pos`t content here.
```

| Field         | Required | Notes                                              |
| ------------- | -------- | -------------------------------------------------- |
| `title`       | Yes      | Displayed in the list and at the top of the post   |
| `date`        | Yes      | ISO format `YYYY-MM-DD`. Posts sort newest first.  |
| `description` | No       | Short summary; shown in the post list (future use) |

The filename (without `.md`) becomes the URL. `posts/my-post-title.md` → `/writing/my-post-title`.

Restart the server after adding a post, as assets are embedded at compile time.

### 5. Resume / static files

Place any static files in `static/`. To link your resume, drop your PDF in `static/` and set `ResumeFile` in `config.go` to its filename.

## Project structure

```
.
├── config.go           # your personal details live here
├── main.go             # HTTP server, routing, post loading
├── go.mod
├── go.sum
├── templates/
│   ├── layout.html     # shared header, nav, footer, theme toggle
│   ├── home.html       # About page content
│   ├── writing.html    # post list
│   ├── post.html       # individual post
│   └── misc.html       # Misc page content
├── static/
│   └── style.css       # Kanagawa Dragon/Lotus theme
└── posts/
    └── hello-world.md  # example post — edit or delete this
```

## Adding or renaming pages

To add a new page (e.g. `/projects`):

1. Create `templates/projects.html` with `{{define "content"}}...{{end}}`
2. Add a handler in `main.go`:
   ```go
   mux.HandleFunc("/projects", handleProjects)
   ```
   ```go
   func handleProjects(w http.ResponseWriter, r *http.Request) {
       render(w, r, "projects", PageData{Title: "Projects | " + siteConfig.Name, Active: "projects"})
   }
   ```
3. Register `"projects"` in the `partialNames` slice in `main()`:
   ```go
   partialNames := []string{"home", "writing", "post", "misc", "projects"}
   ```
4. Add a nav link in `templates/layout.html`:
   ```html
   <a href="/projects"
      hx-get="/projects"
      hx-target="main"
      hx-push-url="true"
      {{if eq .Active "projects"}}aria-current="page"{{end}}>Projects</a>
   ```

## Running on a custom port

```sh
PORT=3000 go run .
```

## Theming

Colors are defined as CSS variables at the top of `static/style.css`. The dark theme uses the [Kanagawa Dragon](https://github.com/rebelot/kanagawa.nvim) palette; the light theme uses Kanagawa Lotus. Edit the `:root` and `[data-theme="light"]` blocks to use whatever colors you want.
