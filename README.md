# Ply

---

## What is **Ply**

Ply is a small Go library made to easily create and **compose** your HTML page by just writing... well... HTML !!!

# :warning: WORK IN PROGRESS

So don't get your expectations to high !

## Install

```bash
go get github.com/dnazirso/ply
```

## Usage

### Example

folder structure example:

```
.
├── main.go
├── components/
│   └── hello.html
└── pages/
    └── index.html
```

main.go

```go
package main

import (
	"html/template"
	"net/http"

	"github.com/dnazirso/ply"
)

func hello(w http.ResponseWriter, r *http.Request) {
	type Hello struct {
		Hello string
	}

	tmpl, _ := template.New("index").Parse(ply.Fold("pages/index", ""))

	data := Hello{
		Hello: "Hello"
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", hello)

	http.ListenAndServe(":8080", nil)
}
```

pages/index.html

```html
<!doctype html>
<html lang="en">
	<head>
		<title>Your APP</title>
		<meta charset="UTF-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1" />
		<link href="css/style.css" rel="stylesheet" />
	</head>
	<body>
		<ply as="components/hello">World !</ply>
	</body>
</html>
```

components/hello.html

```html
<p>{{.Hello}} {{.Children}}</p>
```

### How does it works ?

#### Ply Tag

Pretty straight forward :

You need a `<ply as="path/to/your/component"></ply>` tag in your template!

and call the `ply.Fold()` method to parse your template : 
```go
tmpl, _ := template.New("index").Parse(ply.Fold("path/to/my/template", ""))
```

#### Children

You can pass a children to your ply.

```html
<ply as="path/to/your/component">
	<div class="whatever">This as child</div>
</ply>
```

Do not forget to place the `{{.Children}}` in your component (or don't place it but the children won't apprear)

```html
<p>{{.Children}} place</p>
```

#### Composition

You can compose you plies as you wish

```html
<ply as="path/to/a/component">
	<ply as="another/one">
		<ply as="another/one">but with a child<ply>
		<ply as="yet/another/component">
		<ply as="you/get/the">idea</ply>
	<ply>
</ply>
```
