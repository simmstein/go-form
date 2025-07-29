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

	theme["attributes"] = func(parent map[string]RenderFunc, args ...any) Node {
		var result []Node

		for i, v := range args[0].(form.Attrs) {
			result = append(result, Attr(i, v))
		}

		return Group(result)
	}

	theme["form_attributes"] = func(parent map[string]RenderFunc, args ...any) Node {
		form := args[0].(*form.Form)

		if !form.HasOption("attr") {
			return Raw("")
		}

		return parent["attributes"](parent, form.GetOption("attr").AsAttrs())
	}

	theme["errors"] = func(parent map[string]RenderFunc, args ...any) Node {
		errors := args[0].([]validation.Error)

		var result []Node

		for _, v := range errors {
			result = append(result, Li(Text(string(v))))
		}

		return Ul(
			Group(result),
		)
	}

	theme["form_errors"] = func(parent map[string]RenderFunc, args ...any) Node {
		form := args[0].(*form.Form)

		return If(
			len(form.Errors) > 0,
			parent["errors"](parent, form.Errors),
		)
	}

	theme["form_widget_errors"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		return If(
			len(field.Errors) > 0,
			parent["errors"](parent, field.Errors),
		)
	}

	theme["help"] = func(parent map[string]RenderFunc, args ...any) Node {
		help := args[0].(string)
		var extra Node

		if len(help) == 0 {
			return Raw("")
		}

		if len(args) == 2 {
			extra = args[1].(Node)
		}

		return Div(
			Text(help),
			extra,
		)
	}

	theme["form_help"] = func(parent map[string]RenderFunc, args ...any) Node {
		form := args[0].(*form.Form)

		if !form.HasOption("help") {
			return Raw("")
		}

		return parent["help"](
			parent,
			form.GetOption("help").AsString(),
			parent["attributes"](parent, form.GetOption("help_attr").AsAttrs()),
		)
	}

	theme["form_widget_help"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		if !field.HasOption("help") {
			return Raw("")
		}

		return parent["help"](
			parent,
			field.GetOption("help").AsString(),
			parent["attributes"](parent, field.GetOption("help_attr").AsAttrs()),
		)
	}

	theme["label_attributes"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		return parent["attributes"](parent, field.GetOption("label_attr").AsAttrs())
	}

	theme["form_label"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		if !field.HasOption("label") {
			return Raw("")
		}

		label := field.GetOption("label").AsString()

		return If(len(label) > 0, Label(
			For(field.GetId()),
			parent["label_attributes"](parent, field),
			Text(label),
		))
	}

	theme["field_attributes"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		return parent["attributes"](parent, field.GetOption("attr").AsAttrs())
	}

	theme["textarea_attributes"] = func(parent map[string]RenderFunc, args ...any) Node {
		return parent["field_attributes"](parent, args...)
	}

	theme["input_attributes"] = func(parent map[string]RenderFunc, args ...any) Node {
		return parent["field_attributes"](parent, args...)
	}

	theme["sub_form_attributes"] = func(parent map[string]RenderFunc, args ...any) Node {
		return parent["field_attributes"](parent, args...)
	}

	theme["input"] = func(parent map[string]RenderFunc, args ...any) Node {
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
			If(fieldType == "checkbox" && field.Data != nil && field.Data != false, Checked()),
			If(field.HasOption("required") && field.GetOption("required").AsBool(), Required()),
			parent["input_attributes"](parent, field),
		)
	}

	theme["choice_options"] = func(parent map[string]RenderFunc, args ...any) Node {
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

	theme["choice_expanded_item"] = func(parent map[string]RenderFunc, args ...any) Node {
		return args[0].(Node)
	}

	theme["choice_attributes"] = func(parent map[string]RenderFunc, args ...any) Node {
		return parent["field_attributes"](parent, args...)
	}

	theme["choice_expanded"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)
		choices := field.GetOption("choices").Value.(*form.Choices)

		isRequired := field.HasOption("required") && field.GetOption("required").AsBool()
		isMultiple := field.GetOption("multiple").AsBool()
		noneLabel := field.GetOption("empty_choice_label").AsString()

		var items []Node

		if !isMultiple && !isRequired {
			id := fmt.Sprintf("%s-%s", field.GetId(), "none")

			items = append(items, parent["choice_expanded_item"](parent, Group([]Node{
				Input(
					Name(field.GetName()),
					ID(id),
					Value(""),
					Type("radio"),
					parent["choice_attributes"](parent, field),
					If(cast.ToString(field.Data) == "", Checked()),
				),
				Label(
					For(id),
					Text(noneLabel),
					parent["label_attributes"](parent, field),
				),
			})))
		}

		for key, choice := range choices.GetChoices() {
			id := fmt.Sprintf("%s-%d", field.GetId(), key)

			items = append(items, parent["choice_expanded_item"](parent, Group([]Node{
				Input(
					Name(field.GetName()),
					ID(id),
					Value(choice.Value),
					If(isMultiple, Type("checkbox")),
					If(!isMultiple, Type("radio")),
					parent["choice_attributes"](parent, field),
					If(choices.Match(field, choice.Value), Checked()),
				),
				Label(
					For(id),
					Text(choice.Label),
					parent["label_attributes"](parent, field),
				),
			})))
		}

		return Group(items)
	}

	theme["choice"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		isRequired := field.HasOption("required") && field.GetOption("required").AsBool()
		isExpanded := field.GetOption("expanded").AsBool()
		isMultiple := field.GetOption("multiple").AsBool()
		noneLabel := field.GetOption("empty_choice_label").AsString()

		_ = noneLabel

		if isExpanded {
			return parent["choice_expanded"](parent, field)
		} else {
			return Select(
				ID(field.GetId()),
				If(isRequired, Required()),
				If(isMultiple, Multiple()),
				Name(field.GetName()),
				parent["choice_attributes"](parent, field),
				parent["choice_options"](parent, field),
			)
		}
	}

	theme["textarea"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		return Textarea(
			Name(field.GetName()),
			ID(field.GetId()),
			If(field.HasOption("required") && field.GetOption("required").AsBool(), Required()),
			parent["textarea_attributes"](parent, field),
			Text(cast.ToString(field.Data)),
		)
	}

	theme["sub_form_label"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		if !field.HasOption("label") {
			return Raw("")
		}

		label := field.GetOption("label").AsString()

		return If(len(label) > 0, Legend(
			parent["label_attributes"](parent, field),
			Text(label),
		))

	}

	theme["sub_form_content"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		return parent["form_fields"](parent, field.Children)
	}

	theme["sub_form"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		return FieldSet(
			ID(field.GetId()),
			parent["sub_form_label"](parent, field),
			parent["sub_form_attributes"](parent, field),
			parent["sub_form_content"](parent, field),
		)
	}

	theme["form_widget"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		tpl, ok := parent[field.Widget]

		if !ok {
			return Raw("Invalid field widget: " + field.Widget)
		}

		return tpl(parent, field)
	}

	theme["form_row"] = func(parent map[string]RenderFunc, args ...any) Node {
		field := args[0].(*form.Field)

		isCheckbox := field.HasOption("type") && field.GetOption("type").AsString() == "checkbox"
		hasChildren := len(field.Children) > 0
		labelAfter := isCheckbox && !hasChildren
		label := parent["form_label"](parent, field)
		attrs := Raw("")

		if field.HasOption("row_attr") {
			attrs = parent["attributes"](parent, field.GetOption("row_attr").AsAttrs())
		}

		return Div(
			attrs,
			If(!labelAfter, label),
			parent["form_widget_errors"](parent, field),
			parent["form_widget"](parent, field),
			If(labelAfter, label),
			parent["form_widget_help"](parent, field),
		)
	}

	theme["form_fields"] = func(parent map[string]RenderFunc, args ...any) Node {
		var items []Node

		for _, item := range args[0].([]*form.Field) {
			items = append(items, parent["form_row"](parent, item))
		}

		return Group(items)
	}

	theme["form_content"] = func(parent map[string]RenderFunc, args ...any) Node {
		form := args[0].(*form.Form)

		return Group([]Node{
			parent["form_errors"](parent, form),
			parent["form_help"](parent, form),
			parent["form_fields"](parent, form.Fields),
		})
	}

	theme["form"] = func(parent map[string]RenderFunc, args ...any) Node {
		form := args[0].(*form.Form)

		return Form(
			Action(form.Action),
			Method(form.Method),
			parent["form_attributes"](parent, form),
			parent["form_content"](parent, form),
		)
	}

	return theme
})
