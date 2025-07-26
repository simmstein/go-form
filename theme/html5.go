package theme

import (
	"fmt"

	"github.com/spf13/cast"
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

var Html5 = CreateTheme(func() map[string]RenderFunc {
	theme := make(map[string]RenderFunc)

	theme["attributes"] = func(args ...any) Node {
		var result []Node

		for i, v := range args[0].(map[string]string) {
			result = append(result, Attr(i, v))
		}

		return Group(result)
	}

	theme["form_attributes"] = func(args ...any) Node {
		form := args[0].(*form.Form)

		if !form.HasOption("attr") {
			return Raw("")
		}

		return theme["attributes"](form.GetOption("attr").AsMapString())
	}

	theme["errors"] = func(args ...any) Node {
		errors := args[0].([]validation.Error)

		var result []Node

		for _, v := range errors {
			result = append(result, Li(Text(string(v))))
		}

		return Ul(
			Class("form-errors"),
			Group(result),
		)
	}

	theme["form_errors"] = func(args ...any) Node {
		form := args[0].(*form.Form)

		return If(
			len(form.Errors) > 0,
			theme["errors"](form.Errors),
		)
	}

	theme["form_widget_errors"] = func(args ...any) Node {
		field := args[0].(*form.Field)

		return If(
			len(field.Errors) > 0,
			theme["errors"](field.Errors),
		)
	}

	theme["help"] = func(args ...any) Node {
		help := args[0].(string)

		if len(help) == 0 {
			return Raw("")
		}

		return Div(
			Class("form-help"),
			Text("ok"),
		)
	}

	theme["form_help"] = func(args ...any) Node {
		form := args[0].(*form.Form)

		if !form.HasOption("help") {
			return Raw("")
		}

		return theme["help"](form.GetOption("help").AsString())
	}

	theme["form_widget_help"] = func(args ...any) Node {
		field := args[0].(*form.Field)

		if !field.HasOption("help") {
			return Raw("")
		}

		return theme["help"](field.GetOption("help").AsString())
	}

	theme["label_attributes"] = func(args ...any) Node {
		field := args[0].(*form.Field)

		if !field.HasOption("label_attr") {
			return Raw("")
		}

		return theme["attributes"](field.GetOption("label_attr").AsMapString())
	}

	theme["form_label"] = func(args ...any) Node {
		field := args[0].(*form.Field)

		if !field.HasOption("label") {
			return Raw("")
		}

		label := field.GetOption("label").AsString()

		return If(len(label) > 0, Label(
			Class("form-label"),
			For(field.GetId()),
			theme["label_attributes"](field),
			Text(label),
		))
	}

	theme["field_attributes"] = func(args ...any) Node {
		field := args[0].(*form.Field)

		if !field.HasOption("attr") {
			return Raw("")
		}

		return theme["attributes"](field.GetOption("attr").AsMapString())
	}

	theme["textarea_attributes"] = func(args ...any) Node {
		return theme["field_attributes"](args...)
	}

	theme["input_attributes"] = func(args ...any) Node {
		return theme["field_attributes"](args...)
	}

	theme["sub_form_attributes"] = func(args ...any) Node {
		return theme["field_attributes"](args...)
	}

	theme["input"] = func(args ...any) Node {
		field := args[0].(*form.Field)
		fieldType := "text"

		if field.HasOption("type") {
			fieldType = field.GetOption("type").AsString()
		}

		value := cast.ToString(field.Data)

		if fieldType == "checkbox" {
			value = "1"
		}

		return Input(
			Name(field.GetName()),
			ID(field.GetId()),
			Type(fieldType),
			Value(value),
			If(fieldType == "checkbox" && field.Data != false, Checked()),
			If(field.HasOption("required") && field.GetOption("required").AsBool(), Required()),
			theme["input_attributes"](field),
		)
	}

	theme["choice_options"] = func(args ...any) Node {
		field := args[0].(*form.Field)
		choices := field.GetOption("choices").Value.(*form.Choices)

		isRequired := field.HasOption("required") && field.GetOption("required").AsBool()
		isMultiple := field.GetOption("multiple").AsBool()

		var options []Node

		if !isMultiple && !isRequired {
			options = append(options, Option(
				Text(field.GetOption("empty_choice_label").AsString()),
			))
		}

		for _, choice := range choices.GetChoices() {
			options = append(options, Option(
				Value(choice.Value),
				Text(choice.Label),
				If(choices.Match(field, choice.Value), Selected()),
			))
		}

		return Group(options)
	}

	theme["choice_expanded"] = func(args ...any) Node {
		field := args[0].(*form.Field)
		choices := field.GetOption("choices").Value.(*form.Choices)

		isRequired := field.HasOption("required") && field.GetOption("required").AsBool()
		isMultiple := field.GetOption("multiple").AsBool()
		noneLabel := field.GetOption("empty_choice_label").AsString()

		var items []Node

		if !isMultiple && !isRequired {
			id := fmt.Sprintf("%s-%s", field.GetId(), "none")

			items = append(items, Group([]Node{
				Input(
					Name(field.GetName()),
					ID(id),
					Value(""),
					Type("radio"),
					theme["input_attributes"](field),
					If(cast.ToString(field.Data) == "", Checked()),
				),
				Label(For(id), Text(noneLabel)),
			}))
		}

		for key, choice := range choices.GetChoices() {
			id := fmt.Sprintf("%s-%d", field.GetId(), key)

			items = append(items, Group([]Node{
				Input(
					Name(field.GetName()),
					ID(id),
					Value(choice.Value),
					If(isMultiple, Type("checkbox")),
					If(!isMultiple, Type("radio")),
					theme["input_attributes"](field),
					If(choices.Match(field, choice.Value), Checked()),
				),
				Label(For(id), Text(choice.Label)),
			}))
		}

		return Group(items)
	}

	theme["choice"] = func(args ...any) Node {
		field := args[0].(*form.Field)

		isRequired := field.HasOption("required") && field.GetOption("required").AsBool()
		isExpanded := field.GetOption("expanded").AsBool()
		isMultiple := field.GetOption("multiple").AsBool()
		noneLabel := field.GetOption("empty_choice_label").AsString()

		_ = noneLabel

		if isExpanded {
			return theme["choice_expanded"](field)
		} else {
			return Select(
				ID(field.GetId()),
				If(isRequired, Required()),
				If(isMultiple, Multiple()),
				Name(field.GetName()),
				theme["choice_options"](field),
			)
		}
	}

	theme["textarea"] = func(args ...any) Node {
		field := args[0].(*form.Field)

		return Textarea(
			Name(field.GetName()),
			ID(field.GetId()),
			If(field.HasOption("required") && field.GetOption("required").AsBool(), Required()),
			theme["textarea_attributes"](field),
			Text(cast.ToString(field.Data)),
		)
	}

	theme["sub_form_label"] = func(args ...any) Node {
		field := args[0].(*form.Field)

		if !field.HasOption("label") {
			return Raw("")
		}

		label := field.GetOption("label").AsString()

		return If(len(label) > 0, Legend(
			Class("form-label"),
			theme["label_attributes"](field),
			Text(label),
		))

	}

	theme["sub_form_content"] = func(args ...any) Node {
		field := args[0].(*form.Field)

		return theme["form_fields"](field.Children)
	}

	theme["sub_form"] = func(args ...any) Node {
		field := args[0].(*form.Field)

		return FieldSet(
			ID(field.GetId()),
			theme["sub_form_label"](field),
			theme["sub_form_attributes"](field),
			theme["sub_form_content"](field),
		)
	}

	theme["form_widget"] = func(args ...any) Node {
		field := args[0].(*form.Field)

		tpl, ok := theme[field.Widget]

		if !ok {
			return Raw("Invalid field widget: " + field.Widget)
		}

		return tpl(field)
	}

	theme["form_row"] = func(args ...any) Node {
		field := args[0].(*form.Field)

		isCheckbox := field.HasOption("type") && field.GetOption("type").AsString() == "checkbox"
		hasChildren := len(field.Children) > 0
		labelAfter := isCheckbox && !hasChildren
		label := theme["form_label"](field)

		return Div(
			Class("form-row"),
			If(!labelAfter, label),
			theme["form_widget_errors"](field),
			theme["form_widget"](field),
			If(labelAfter, label),
			theme["form_widget_help"](field),
		)
	}

	theme["form_fields"] = func(args ...any) Node {
		var items []Node

		for _, item := range args[0].([]*form.Field) {
			items = append(items, theme["form_row"](item))
		}

		return Group(items)
	}

	theme["form_content"] = func(args ...any) Node {
		form := args[0].(*form.Form)

		return Div(
			theme["form_errors"](form),
			theme["form_help"](form),
			theme["form_fields"](form.Fields),
		)
	}

	theme["form"] = func(args ...any) Node {
		form := args[0].(*form.Form)

		return Form(
			Action(form.Action),
			Method(form.Method),
			theme["form_attributes"](form),
			theme["form_content"](form),
		)
	}

	return theme
})
