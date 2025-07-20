# go-form

Creating and processing HTML forms is hard and repetitive. You need to deal with rendering HTML form fields, validating submitted data, mapping the form data into objects and a lot more. [`go-form`][go-form] includes a powerful form feature that provides all these features.

## Introduction

[`go-form`][go-form]  is heavily influenced by [Symfony Form](https://symfony.com/doc/current/forms.html). It includes:

* A form builder based on fields declarations and independent of structs
* Validation based on constraints
* Data mounting to populate a form from a struct instance
* Data binding to populate a struct instance from a submitted form
* Form renderer with customizable themes

## Installation

```shell
go get gitnet.fr/deblan/go-form
```

## Quick Start

```go
package main

import (
	"html/template"
	"log"
	"net/http"

	"gitnet.fr/deblan/go-form/form"
	"gitnet.fr/deblan/go-form/theme"
	"gitnet.fr/deblan/go-form/validation"
)

func main() {
	type Person struct {
		Name string
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := new(Person)

		f := form.NewForm(
			form.NewFieldText("Name").
				WithOptions(
					form.NewOption("label", "Your name"),
				).
				WithConstraints(
					validation.NewNotBlank(),
				),
		).
			End().
			WithMethod(http.MethodPost).
			WithAction("/")

		f.Mount(data)

		if r.Method == f.Method {
			f.HandleRequest(r)

			if f.IsSubmitted() && f.IsValid() {
				f.Bind(data)
			}
		}

		render := theme.NewRenderer(theme.Html5)
		tpl, _ := template.New("page").Funcs(render.FuncMap()).Parse(`{{ form .Form }}`)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tpl.Execute(w, map[string]any{
			"Form": f,
		})
	})

	log.Fatal(http.ListenAndServe(":1324", nil))
}
```

[go-form]: https://gitnet.fr/deblan/go-form
