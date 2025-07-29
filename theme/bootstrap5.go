package theme

import (
	"gitnet.fr/deblan/go-form/form"
	"gitnet.fr/deblan/go-form/validation"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

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

var Bootstrap5 = ExtendTheme(Html5, func() map[string]RenderFunc {
	theme := make(map[string]RenderFunc)

	theme["form_help"] = func(parent map[string]RenderFunc, args ...any) Node {
		form := args[0].(*form.Form)

		form.GetOption("help_attr").AsAttrs().Append("class", "form-text")

		return parent["base_form_help"](parent, form)
	}

	theme["form_widget_help"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		field.GetOption("help_attr").AsAttrs().Append("class", "form-text")

		return parent["base_form_widget_help"](parent, field)
	}

	theme["input"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)
		fieldType := field.GetOption("type").AsString()

		var class string

		if fieldType == "checkbox" || fieldType == "radio" {
			class = "form-check-input"
		} else if fieldType == "range" {
			class = "form-range"
		} else if fieldType == "button" || fieldType == "submit" || fieldType == "reset" {
			class = "btn"
		} else {
			class = "form-control"
		}

		field.GetOption("attr").AsAttrs().Append("class", class)

		return parent["base_input"](parent, field)
	}

	theme["form_label"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		var class string

		if field.Widget == "choice" && field.HasOption("expanded") && field.GetOption("expanded").AsBool() {
			class = "form-check-label"
		} else {
			class = "form-label"
		}

		field.GetOption("label_attr").AsAttrs().Append("class", class)

		return parent["base_form_label"](parent, field)
	}

	theme["choice"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		if !field.HasOption("expanded") || !field.GetOption("expanded").AsBool() {
			field.GetOption("attr").AsAttrs().Append("class", "form-control")
		}

		return parent["base_choice"](parent, field)
	}

	theme["choice_expanded_item"] = func(parent map[string]RenderFunc, args ...any) Node {
		return Div(Class("form-check"), args[0].(Node))
	}

	theme["textarea"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		field.GetOption("attr").AsAttrs().Append("class", "form-control")

		return parent["base_textarea"](parent, field)
	}

	theme["form_row"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		if field.HasOption("type") {
			fieldType := field.GetOption("type").AsString()

			if fieldType == "checkbox" || fieldType == "radio" {
				field.GetOption("row_attr").AsAttrs().Append("class", "form-check")
			}
		}

		return parent["base_form_row"](parent, field)
	}

	theme["errors"] = func(parent map[string]RenderFunc, args ...any) Node {
		errors := args[0].([]validation.Error)

		var result []Node

		for _, v := range errors {
			result = append(result, Text(string(v)))
			result = append(result, Br())
		}

		return Div(
			Class("invalid-feedback d-block"),
			Group(result),
		)
	}

	return theme
})
