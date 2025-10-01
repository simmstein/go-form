package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/yassinebenaid/godump"
	"gitnet.fr/deblan/go-form/example"
	"gitnet.fr/deblan/go-form/theme"
)

//go:embed example/view/*.html
var templates embed.FS

func handler(view, action string, formRenderer *theme.Renderer, w http.ResponseWriter, r *http.Request) {
	entity := example.ExampleData{}
	form := example.CreateDataForm(action)

	form.Mount(entity)

	if r.Method == form.Method {
		form.HandleRequest(r)

		if form.IsSubmitted() && form.IsValid() {
			form.Bind(&entity)
		}
	}

	content, _ := templates.ReadFile(view)

	formAsJson, _ := json.MarshalIndent(form, "  ", "  ")

	tpl, _ := template.New("page").
		Funcs(formRenderer.FuncMap()).
		Parse(string(content))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var dump godump.Dumper
	dump.Theme = godump.Theme{}

	tpl.Execute(w, map[string]any{
		"isSubmitted": form.IsSubmitted(),
		"isValid":     form.IsValid(),
		"form":        form,
		"json":        string(formAsJson),
		"dump":        template.HTML(dump.Sprint(entity)),
	})
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(
			"example/view/html5.html",
			"/",
			theme.NewRenderer(theme.Html5),
			w,
			r,
		)
	})

	http.HandleFunc("/bootstrap", func(w http.ResponseWriter, r *http.Request) {
		handler(
			"example/view/bootstrap.html",
			"/bootstrap",
			theme.NewRenderer(theme.Bootstrap5),
			w,
			r,
		)
	})

	log.Fatal(http.ListenAndServe(":1122", nil))
}
