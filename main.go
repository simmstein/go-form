package main

import (
	"fmt"

	"gitnet.fr/deblan/go-form/form"
	"gitnet.fr/deblan/go-form/theme"
	"gitnet.fr/deblan/go-form/validation"
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

	data := new(Person)
	data.Name = ""
	data.Address = Address{
		Street:  "rue des camélias",
		City:    "",
		ZipCode: 39700,
	}

	f := form.NewForm(
		form.NewFieldText("Name").
			WithOptions(
				form.Option{Name: "required", Value: true},
			).
			WithConstraints(
				validation.NotBlank{},
			),
		form.NewSubForm("Address").
			Add(
				form.NewFieldText("Street"),
				form.NewFieldText("City").
					WithConstraints(
						validation.NotBlank{},
					),
				form.NewFieldText("ZipCode"),
			),
	).WithMethod("POST").WithAction("")

	f.Bind(data)

	fmt.Printf("%+v\n", f.IsValid())

	render := theme.NewRenderer(theme.Html5)
	v := render.RenderForm(f)

	fmt.Print(v)

	// r, e := theme.RenderForm(f, theme.Html5)

	// fmt.Printf("%+v\n", e)
	// fmt.Printf("%+v\n", r)

	// fmt.Printf("%+v\n", e)
	//
	// fmt.Printf("%+v\n", f)
	//
	// for _, field := range f.Fields {
	// 	fmt.Printf("%+v\n", *field)
	//
	// 	if len(field.Children) > 0 {
	// 		for _, c := range field.Children {
	// 			fmt.Printf("C %+v\n", *c)
	// 		}
	// 	}
	// }
}
