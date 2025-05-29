---
title: "index"
hidden: true
---

## 1. Install

```shell
go install github.com/peterszarvas94/lytepage@latest
```

## 1. Initialize a new project

```shell
lytepage init my-app
```

## 2. Write content

`content/hello/world.md`

```txt
---
title: "Amazing article"
---

## Hello world

How awesome this site is!

Learn:

- go
- templ
```

## 3. Write template

```js
templ PostPage(post *fileutils.FileData) {
  <h1>
    { post.Matter.Title }
  </h1>
  <div class="markdown-body">
    @fileutils.HtmlString(post.Html)
  </div>
}
```

## 4. Generate (for static hosting)

```shell
make gen
```

The generated `public/hello/world/index.html`:

```html
<h1>Amazing article</h1>
<div class="markdown-body">
	<h2>Hello world!</h2>

	<p>How awesome this site is!</p>

	<p>Learn:</p>

	<ul>
		<li>go</li>
		<li>templ</li>
	</ul>
</div>
```

## 5. Run locally

SSR mode:

```shell
make ssr
```

Static mode:

```shell
make static
```

Now `localhost:8080/hello/world` is serving your content

Dev mode - soon!
