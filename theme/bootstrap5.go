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

var Bootstrap5 = map[string]string{
	"form": `<form action="{{ .Form.Action }}" method="{{ .Form.Method }}" {{ form_attr .Form }}>
		{{- form_error .Form nil -}}

		{{- form_help .Form -}}

		{{- range $field := .Form.Fields -}}
			{{- form_row $field -}}
		{{- end -}}
	</form>`,
	"attributes": `{{ range $key, $value := .Attributes }}{{ $key }}="{{ $value }}"{{ end }}`,
	"help": `
		{{- if gt (len .Help) 0 -}}
			<div class="form-help">{{ .Help }}</div>
		{{- end -}}
	`,
	"label": `
		{{ if .Field.HasOption "label" }}
			{{ $label := (.Field.GetOption "label").Value }}

			{{- if ne $label "" -}}
				<label for="{{ .Field.GetId }}" {{ form_label_attr .Field }}  class="form-label">{{ $label }}</label>
			{{- end -}}
		{{- end -}}
	`,
	"input": `
		{{- $type := .Field.GetOption "type" -}}
		{{- $checked := and (eq (.Field.GetOption "type").Value "checkbox") (.Field.Data) -}}
		{{- $required := and (.Field.HasOption "required") (.Field.GetOption "required").Value -}}
		{{- $value := .Field.Data -}}
		{{- $class := "form-control" }}

		{{- if eq $type.Value "checkbox" -}}
			{{- $value = 1 -}}
		{{- end -}}
		
		{{- if or (eq $type.Value "checkbox") (eq $type.Value "radio") -}}
			{{- $class = "form-check-input" -}}
		{{- end -}}

		{{- if eq $type.Value "range" -}}
			{{- $class = "form-range" -}}
		{{- end -}}

		{{- if or (eq $type.Value "submit") (eq $type.Value "reset") (eq $type.Value "button") -}}
			{{- $class = "" -}}

			{{ if .Field.HasOption "attr" }}
				{{ $class = (.Field.GetOption "attr").Value.attr.class }}
			{{ end }}
		{{- end -}}

		<input id="{{ .Field.GetId }}" {{ if $checked }}checked{{ end }} {{ if $required }}required="required"{{ end }} name="{{ .Field.GetName }}" value="{{ $value }}" type="{{ $type.Value }}" {{ form_widget_attr .Field }} class="{{ $class }}">
	`,
	"textarea": `
		<textarea id="{{ .Field.GetId }}" {{ if .Field.HasOption "required" }}{{ if (.Field.GetOption "required").Value }}required="required"{{ end }}{{ end }} name="{{ .Field.GetName }}" {{ form_widget_attr .Field }} class="form-control">{{ .Field.Data }}</textarea>
	`,
	"choice": `
		{{- $required := and (.Field.HasOption "required") (.Field.GetOption "required").Value -}}
		{{- $isExpanded := (.Field.GetOption "expanded").Value -}}
		{{- $isMultiple := (.Field.GetOption "multiple").Value -}}
		{{- $emptyChoiceLabel := (.Field.GetOption "empty_choice_label").Value -}}
		{{- $choices := (.Field.GetOption "choices").Value -}}
		{{- $field := .Field -}}
		{{- $keyAdd := 0 -}}

		{{- if and (not $required) (not $isMultiple) -}}
			{{- $keyAdd = 1 -}}
		{{- end -}}

		{{- if $isExpanded -}}
			{{- if and (not $required) (not $isMultiple) -}}
				<div class="form-check">
					<input value="" {{ if not $field.Data }}checked{{ end }} name="{{ $field.GetName }}" type="radio" id="{{ $field.GetId }}-0" class="form-check-input">
					<label for="{{ $field.GetId }}-0" class="form-check-label">{{ ($field.GetOption "empty_choice_label").Value }}</label>
				</div>
			{{- end -}}

			{{- range $key, $choice := $choices.GetChoices -}}
				<div class="form-check">
					<input name="{{ $field.GetName }}" type="{{ if $isMultiple }}checkbox{{ else }}radio{{ end }}" value="{{ $choice.Value }}" {{ if $choices.Match $field $choice.Value }}checked{{ end }} id="{{ $field.GetId }}-{{ sum $key $keyAdd }}" class="form-check-input">
					<label for="{{ $field.GetId }}-{{ sum $key $keyAdd }}" class="form-check-label">{{- $choice.Label -}}</label>
				</div>
			{{- end -}}
		{{- else -}}
			<select id="{{ .Field.GetId }}" {{ if $required }}required="required"{{ end }} {{ if $isMultiple }}multiple{{ end }} name="{{ .Field.GetName }}" {{ form_widget_attr .Field }} class="form-select">
				{{- if and (not $required) (not $isMultiple) -}}
					<option value="">{{ $emptyChoiceLabel }}</option>
				{{- end -}}
				{{- range $choice := $choices.GetChoices -}}
					<option value="{{ $choice.Value }}" {{ if $choices.Match $field $choice.Value }}selected{{ end }}>{{ $choice.Label }}</option>
				{{- end -}}
			</select>
		{{- end -}}
	`,
	"sub_form": `
		<fieldset id="{{ .Field.GetId }}">
			{{ if .Field.HasOption "label" }}
				{{ $label := (.Field.GetOption "label").Value }}

				{{- if ne $label "" -}}
					<legend {{ form_label_attr .Field }}>{{ $label }}</legend>
				{{- end -}}
			{{- end -}}

			{{ form_widget_help .Field }}

			{{- range $field := .Field.Children -}}
				{{- form_row $field -}}
			{{- end -}}
		</fieldset>
	`,
	"error": `
		{{- if gt (len .Errors) 0 -}}
			<div class="invalid-feedback d-block">
				{{- range $error := .Errors -}}
					<div>{{- $error -}}</div>
				{{- end -}}
			</div>
		{{- end -}}
	`,
	"row": `<div {{ form_row_attr .Field }}>
		{{ $labelAfterWidget := and (.Field.HasOption "type") (eq (.Field.GetOption "type").Value "checkbox") }}

		{{ if and (eq (len .Field.Children) 0) (not $labelAfterWidget) }}
			{{- form_label .Field -}}
		{{ end }}

		{{- form_widget .Field -}}
		{{- form_error nil .Field -}}

		{{ if and (eq (len .Field.Children) 0) ($labelAfterWidget) }}
			{{- form_label .Field -}}
		{{ end }}

		{{- form_widget_help .Field -}}
	</div>`,
}
