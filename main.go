package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"gitnet.fr/deblan/go-form/example"
	"gitnet.fr/deblan/go-form/theme"
)

func main() {
	type Address struct {
		Street  string
		City    string
		ZipCode uint
	}

	type Foo struct {
		Name     string
		Address  Address
		Date     time.Time
		DateTime time.Time
		Time     time.Time
		Checkbox bool
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// now := time.Now()
		// data := new(Foo)
		// data.Name = ""
		// data.Date = now
		// data.DateTime = now
		// data.Time = now
		// data.Address = Address{
		// 	Street:  "",
		// 	City:    "",
		// 	ZipCode: 39700,
		// }
		//
		// f := example.CreateExampleForm()
		// f.Mount(data)

		data := example.Post{
			// Tags: []example.Tag{example.Tag{"tag1"}, example.Tag{"tag2"}, example.Tag{"tag3"}},
			Tag: example.Tag{"tag1"},
		}

		f := example.CreateExampleForm2()
		f.Mount(data)

		if r.Method == f.Method {
			f.HandleRequest(r)

			if f.IsSubmitted() && f.IsValid() {
				f.Bind(&data)
				fmt.Printf("BIND=%+v\n", data)
			}
		}

		render := theme.NewRenderer(theme.Html5)

		tpl, _ := template.New("page").Funcs(render.FuncMap()).Parse(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>Form</title>
				<style>
					input[type="text"],
					input[type="date"],
					input[type="datetime"],
					input[type="time"],
					input[type="range"],
					input[type="email"],
					select,
					input[type="datetime-local"],
					textarea {
						box-sizing: border-box;
						padding: 9px;
						margin: 10px 0;
						display: block;
						width: 100%;
						border: 1px solid black;
					}

					.form-errors {
						margin: 0;
						padding: 5px 0 0 0;
						color: red;
						list-style: none;
					}

					.form-errors li {
						padding: 0;
						margin: 0;
					}

					.form-help {
						color: blue;
						font-size: 9px;
					}
				</style>
			</head>
			<body>
				{{ form .Form }}
			</body>
			</html>
		`)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tpl.Execute(w, map[string]any{
			"Form": f,
			// "Post": data,
		})
	})

	log.Fatal(http.ListenAndServe(":1122", nil))
}
