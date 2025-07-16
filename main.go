package main

import (
	"fmt"
	"log"
	"net/http"

	"gitnet.fr/deblan/go-form/example"
	"gitnet.fr/deblan/go-form/theme"
)

func main() {
	type Address struct {
		Street  string
		City    string
		ZipCode uint
	}

	type Person struct {
		Name    string
		Address Address
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := new(Person)
		data.Name = ""
		data.Address = Address{
			Street:  "rue des camélias",
			City:    "",
			ZipCode: 39700,
		}

		f := example.CreateAddressForm()
		f.Bind(data)

		if r.Method == f.Method {
			f.HandleRequest(r)

			if f.IsSubmitted() && f.IsValid() {
				fmt.Printf("%+v\n", "OK")
			} else {
				fmt.Printf("%+v\n", "KO")
			}
		}

		render := theme.NewRenderer(theme.Html5)
		v := render.RenderForm(f)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(v))
	})

	log.Fatal(http.ListenAndServe(":1122", nil))
}
