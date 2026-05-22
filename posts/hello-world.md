---
title: "Hello, World"
date: "2026-01-01"
description: "My first post."
---

This is your first post. Edit or delete this file, then add your own in the `posts/` directory.

## Writing a post

Each post is a Markdown file with a frontmatter block at the top:

```
---
title: "Post Title"
date: "YYYY-MM-DD"
description: "One-line summary shown in the post list."
---

Your content here.
```

The filename (without `.md`) becomes the URL slug. For example, `posts/hello-world.md` is served at `/writing/hello-world`.

## Markdown support

Standard Markdown works: **bold**, *italic*, `inline code`, [links](https://example.com).

### Code blocks

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

### Blockquotes

> A quote or callout can go here.

Posts are sorted by date, newest first.
