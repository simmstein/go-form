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

	"gitnet.fr/deblan/go-form/form"
	"maragu.dev/gomponents"
)

type RenderFunc func(parent map[string]RenderFunc, args ...any) gomponents.Node

type Renderer struct {
	Theme map[string]RenderFunc
}

func NewRenderer(theme map[string]RenderFunc) *Renderer {
	r := new(Renderer)
	r.Theme = theme

	return r
}

func toTemplateHtml(n gomponents.Node) template.HTML {
	var buf bytes.Buffer

	n.Render(&buf)

	return template.HTML(buf.String())
}

func (r *Renderer) RenderForm(form *form.Form) template.HTML {
	return toTemplateHtml(r.Theme["form"](r.Theme, form))
}

func (r *Renderer) FuncMap() template.FuncMap {
	funcs := template.FuncMap{}

	for _, name := range []string{"form", "form_errors"} {
		funcs[name] = func(form *form.Form) template.HTML {
			return toTemplateHtml(r.Theme[name](r.Theme, form))
		}
	}

	for _, name := range []string{"form_row", "form_widget", "form_label", "form_widget_errors"} {
		funcs[name] = func(field *form.Field) template.HTML {
			return toTemplateHtml(r.Theme[name](r.Theme, field))
		}
	}

	return funcs
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
