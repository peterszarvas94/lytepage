---
title: "index"
hidden: true
---

## 1. Install

```shell
git clone github.com/peterszarvas94/lytepage my-site
cd my-site
```

## 2. Write content

`content/hello.md`

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

## 4. Generate

```shell
make gen
```

The generated `public/hello/index.html`:

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

```shell
make ssg
```

Now `localhost:8080/hello` is serving your content
