---
title: "Introduction"
category: "docs"
tags: ["intro", "readme"]
date: "2024.08.12"
time: "11:15"
excerpt: "The basics to get you started"
---

## Goals

1. Write your templates in go templ
2. Write your content in markdown
3. Generate static HTML
4. Serve HTML with go

## Dependencies

- go 1.22.5
- air
- templ
- make

## Commands

### Develoment mode

- `make templ`: start templ file generation in watch mode
- `make dev`: start dev server with hot reload

### SSR (Server Side Rendering) mode

- `make build`: build go binary
- `make ssr`: start server

### SSG (Static Site Generation) mode

- `make gen`: generate static files

You can put these static files to any server. You can serve them with
**lytepage** as well:

- `make ssg`: start static file server

## Themes

You can put you custom theme under the `theme` folder.

```txt
theme
├── static
└── templates
```

Put static files like `css`, `js`, images etc. under `theme/static`.

Put `templ` components under `theme/templates`.

You can use any other folder which you may need, like `tailwind` folder for
generating tailwind css, or `script` folder for generating JavaScript from
Typescript. The important this is to output the generated files to
`theme/static`, and link them at your pages `<head>` element.

Every theme should have the following `templ` components exported from the root
directory of the theme, as you can see in the `pages/page.go`:

```go
type (
  NotFoundPage func() templ.Component
  IndexPage    func(files []*fileutils.FileData) templ.Component
  PostPage     func(post *fileutils.FileData) templ.Component
  TagPage      func(tag string, files []*fileutils.FileData) templ.Component
  CategoryPage func(category string, files []*fileutils.FileData) templ.Component
)
```

### Default theme

The default theme used for the [lytepage](https://lytepage.peterszarvas.hu)
site, is included in this repo, under `theme` folder.

### Switching themes

Theme switching is a manual process for now, mabye should be improved in the
future. Make sure to backup you old theme folder, or commit it to a version
control, before cloning or writing a new theme, since the old theme are not
saved anywhere else.

## Content

You should put your content in the `content` folder in any folder structure,
using `.md` files with [frontmatter](#post). There are protected routes which
you can not use, and it will cause the program to panic. These are the
following:

- /404
- /static
- /tag
- /category

as well as any [custom route](#custom-routes) you define.

## Routes

The structure of the `content` folder is refcted in the actual routes
(file-based routing). E.g: `content/post/hello.md` is served on `/post/hello`,
both in dev and SSR mode. In SSG mode, the generated folder structore is the
same, e.g. `post/index.html`, so it gets served correctly by any static file
server.

Static assets are served / generated similary from the `theme/static` folder.

### 404

There is a custom 404 (not found) page rendered. In a custom file server or CDN,
you need to implement a catch-all logic, if you want to use this page, because
the content is static. If you serve static files via the built-in file server
with the `make ssg` command, this logic is already implemented.

In dev and SSR modes, catch-all logic is handled as well.

### Index

The homepage displas the latest posts, like in other blog-styled frameworks. You
can define a [custom](#custom-routes) index page of course.

### Post {#post}

The post page should display an individual article / post.

Content is written in markdown, with some common extensions. The `.md` file
should start with the following frontmatter in `yaml` format:

```yml
title: "Title"
category: "category"
tags: ["tag1", "tag2"]
date: "2024.01.02"
time: "11:15"
excerpt: "This is an excrept from the article"
```

The date and datetime format can be configured in the `pkg/config/config.go`
file.

The files are parsed, routes are handled or files are generated as mentioned
before. In your template, you can use the following struct for any article:

```go
type FileData struct {
  Fileroute string      `json:"fileroute"`
  Matter    FrontMatter `json:"matter"`
  DateTime  time.Time   `json:"datetime"`
  Html      string      `json:"html"`
  Content   string      `json:"content"`
  Path      string      `json:"path"`
}

type FrontMatter struct {
  Title    string   `yaml:"title"`
  Category string   `yaml:"category"`
  Tags     []string `yaml:"tags"`
  Date     string   `yaml:"date"`
  Time     string   `yaml:"time"`
  Excerpt  string   `yaml:"excerpt"`
  Hidden   bool     `yaml:"hidden"`
}
```

If you set `hidden: true` in the frontmatter, your file will not be generated or
served as usual, but you can still use the generated content anywhere in your
template. E.g. the `index.md` file is hidden, so it does not appear in the
article list, but the landing page uses it`s content. See
[custom routes](#custom-routes) below.

### Category

A category page is generated, based on the frontmatter of the content files, for
every category.

E.g.: `content/hello.md` has a frontmatter like this:

```yml
category: "test"
```

In this case, a category page is generated and served on `/category/test` route
in dev an SSR mode, or `/category/test/index.html` file is generated in SSG
mode.

### Tag

Similary to category pages, tag pages are generated based on frontmatter. One
file can be one category, but can have multiple tags.

```yml
tags: ["tag1", "tag2"]
```

In the above example, `/tag/tag1` and `/tag/tag2` pages are generated.

### Custom routes {#custom-routes}

You can define custom routes in `pkg/custom/custom.go` by adding element to the
`Routes` map:

```go
var Routes = &routeMap{
  "/":       templates.CustomIndexPage(fileutils.GetFileByTitle("index")),
  "/search": templates.SearchPage(fileutils.GetFiles()),
  "/docs":   templates.PostsPage(fileutils.GetFiles()),
}
```

E.g. the landing page of [lytepage](https://lytepage.peterszarvas.hu) site is a
custom route.

## Search

Since this is a static site generator, server-side seach is not supported. A
static `files.json` file is generated from your files, so you can perform
client-side search with tools like `minisearch`, or any other way.

TODO: files.json generation should be opt-in
