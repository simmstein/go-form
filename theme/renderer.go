package theme

import (
	"bytes"
	"html/template"

	"gitnet.fr/deblan/go-form/form"
	"gitnet.fr/deblan/go-form/validation"
)

type Renderer struct {
	Theme map[string]string
}

func NewRenderer(theme map[string]string) *Renderer {
	r := new(Renderer)
	r.Theme = theme

	return r
}

func (r *Renderer) RenderForm(form *form.Form) template.HTML {
	content := ""

	for _, field := range form.Fields {
		content = content + string(r.RenderRow(field))
	}

	return r.Render("form", r.Theme["form"], map[string]any{
		"Form":    form,
		"Content": template.HTML(content),
	})
}

func (r *Renderer) RenderRow(field *form.Field) template.HTML {
	return r.Render("row", r.Theme["row"], map[string]any{
		"Field": field,
	})
}

func (r *Renderer) RenderLabel(field *form.Field) template.HTML {
	return r.Render("label", r.Theme["label"], map[string]any{
		"Field": field,
	})
}

func (r *Renderer) RenderWidget(field *form.Field) template.HTML {
	return r.Render("widget", r.Theme[field.Widget], map[string]any{
		"Field": field,
	})
}

func (r *Renderer) RenderError(form *form.Form, field *form.Field) template.HTML {
	var errors []validation.Error

	if field != nil {
		errors = field.Errors
	}

	if form != nil {
		errors = form.Errors
	}

	return r.Render("error", r.Theme["error"], map[string]any{
		"Errors": errors,
	})
}

func (r *Renderer) Render(name, tpl string, args any) template.HTML {
	t, err := template.New(name).Funcs(template.FuncMap{
		"form":        r.RenderForm,
		"form_row":    r.RenderRow,
		"form_label":  r.RenderLabel,
		"form_widget": r.RenderWidget,
		"form_error":  r.RenderError,
	}).Parse(tpl)

	if err != nil {
		return template.HTML(err.Error())
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, args)

	if err != nil {
		return template.HTML(err.Error())
	}

	return template.HTML(buf.String())
}
