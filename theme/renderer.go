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
	return r.Render("form", r.Theme["form"], map[string]any{
		"Form": form,
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

func (r *Renderer) RenderLabelAttr(field *form.Field) template.HTMLAttr {
	var attributes map[string]string

	if field.HasOption("label_attr") {
		attributes = field.GetOption("label_attr").Value.(map[string]string)
	}

	return r.RenderAttr("label_attr", r.Theme["attributes"], map[string]any{
		"Attributes": attributes,
	})
}

func (r *Renderer) RenderWidgetAttr(field *form.Field) template.HTMLAttr {
	var attributes map[string]string

	if field.HasOption("attr") {
		attributes = field.GetOption("attr").Value.(map[string]string)
	}

	return r.RenderAttr("widget_attr", r.Theme["attributes"], map[string]any{
		"Attributes": attributes,
	})
}

func (r *Renderer) RenderFormAttr(form *form.Form) template.HTMLAttr {
	var attributes map[string]string

	if form.HasOption("attr") {
		attributes = form.GetOption("attr").Value.(map[string]string)
	}

	return r.RenderAttr("form_attr", r.Theme["attributes"], map[string]any{
		"Attributes": attributes,
	})
}

func (r *Renderer) RenderAttr(name, tpl string, args any) template.HTMLAttr {
	t, err := template.New(name).Parse(tpl)

	if err != nil {
		return template.HTMLAttr("")
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, args)

	if err != nil {
		return template.HTMLAttr("")
	}

	return template.HTMLAttr(buf.String())
}

func (r *Renderer) Render(name, tpl string, args any) template.HTML {
	t, err := template.New(name).Funcs(template.FuncMap{
		"form":        r.RenderForm,
		"form_row":    r.RenderRow,
		"form_label":  r.RenderLabel,
		"form_widget": r.RenderWidget,
		"form_error":  r.RenderError,
		"form_attr":   r.RenderFormAttr,
		"widget_attr": r.RenderWidgetAttr,
		"label_attr":  r.RenderLabelAttr,
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
