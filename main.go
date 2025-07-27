package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/yassinebenaid/godump"
	"gitnet.fr/deblan/go-form/example"
	"gitnet.fr/deblan/go-form/theme"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := example.ExampleData{}

		f := example.CreateDataForm()
		f.Mount(data)

		if r.Method == f.Method {
			f.HandleRequest(r)

			if f.IsSubmitted() && f.IsValid() {
				f.Bind(&data)
				godump.Dump(data)
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
					input[type="number"],
					input[type="password"],
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

					.debug {
						padding: 10px;
					}

					.debug .debug-value {
						color: #555;
						padding: 10px 0 0 10px;
					}
				</style>
			</head>
			<body>
				<div class="debug">
					<div>
						<strong>Submitted</strong>
						<span class="debug-value">{{ .Form.IsSubmitted }}</span>
					</div>
					<div>
						<strong>Valid</strong>
						<span class="debug-value">{{ .Form.IsValid }}</span>
					</div>
					<div>
						<strong>Data</strong>
						<pre class="debug-valid">{{ .Dump }}</pre>
					</div>
				</div>

				{{ form .Form }}
			</body>
			</html>
		`)

		var dump godump.Dumper
		dump.Theme = godump.Theme{}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tpl.Execute(w, map[string]any{
			"Form": f,
			"Dump": template.HTML(dump.Sprint(data)),
		})
	})

	http.HandleFunc("/bootstrap", func(w http.ResponseWriter, r *http.Request) {
		data := example.ExampleData{}

		f := example.CreateDataForm()
		f.WithAction("/bootstrap")
		f.Mount(data)

		if r.Method == f.Method {
			f.HandleRequest(r)

			if f.IsSubmitted() && f.IsValid() {
				f.Bind(&data)
				godump.Dump(data)
			}
		}

		render := theme.NewRenderer(theme.Bootstrap5)

		tpl, _ := template.New("page").Funcs(render.FuncMap()).Parse(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>Form</title>
				<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.7/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-LN+7fdVzj6u52u30Kp6M/trliBMCMKTyK833zpbD+pXdCLuTusPj697FH4R/5mcr" crossorigin="anonymous">
			</head>
			<body>
				<div class="container">
					<div class="list-group">
						<div class="list-group-item">
							<strong>Submitted</strong>
							<span class="debug-value">{{ .Form.IsSubmitted }}</span>
						</div>
						<div class="list-group-item">
							<strong>Valid</strong>
							<span class="debug-value">{{ .Form.IsValid }}</span>
						</div>
						<div class="list-group-item">
							<strong>Data</strong>
							<pre class="debug-valid">{{ .Dump }}</pre>
						</div>
					</div>

					{{ form .Form }}
				</div>
			</body>
			</html>
		`)

		var dump godump.Dumper
		dump.Theme = godump.Theme{}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tpl.Execute(w, map[string]any{
			"Form": f,
			"Dump": template.HTML(dump.Sprint(data)),
		})

		os.Stdout
	})

	log.Fatal(http.ListenAndServe(":1122", nil))
}
