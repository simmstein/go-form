package theme

// @license GNU AGPL version 3 or any later version
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

import (
	"bytes"
	"html/template"

	"github.com/spf13/cast"
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

func (r *Renderer) RenderRowAttr(field *form.Field) template.HTMLAttr {
	var attributes map[string]string

	if field.HasOption("row_attr") {
		attributes = field.GetOption("row_attr").Value.(map[string]string)
	}

	return r.RenderAttr("raw_attr", r.Theme["attributes"], map[string]any{
		"Attributes": attributes,
	})
}

func (r *Renderer) RenderFormHelp(form *form.Form) template.HTML {
	var help string

	if form.HasOption("help") {
		help = form.GetOption("help").Value.(string)
	}

	return r.Render("help", r.Theme["help"], map[string]any{
		"Help": help,
	})
}

func (r *Renderer) RenderWidgetHelp(field *form.Field) template.HTML {
	var help string

	if field.HasOption("help") {
		help = field.GetOption("help").Value.(string)
	}

	return r.Render("help", r.Theme["help"], map[string]any{
		"Help": help,
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

func (r *Renderer) FuncMap() template.FuncMap {
	return template.FuncMap{
		"form":             r.RenderForm,
		"form_row":         r.RenderRow,
		"form_label":       r.RenderLabel,
		"form_widget":      r.RenderWidget,
		"form_error":       r.RenderError,
		"form_attr":        r.RenderFormAttr,
		"form_widget_attr": r.RenderWidgetAttr,
		"form_label_attr":  r.RenderLabelAttr,
		"form_row_attr":    r.RenderRowAttr,
		"form_help":        r.RenderFormHelp,
		"form_widget_help": r.RenderWidgetHelp,
		"sum": func(values ...any) float64 {
			res := float64(0)
			for _, value := range values {
				res += cast.ToFloat64(value)
			}

			return res
		},
	}
}

func (r *Renderer) Render(name, tpl string, args any) template.HTML {
	t, err := template.New(name).Funcs(r.FuncMap()).Parse(tpl)

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
