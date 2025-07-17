package theme

var Html5 = map[string]string{
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
				<label for="{{ .Field.GetId }}" {{ form_label_attr .Field }}>{{ $label }}</label>
			{{- end -}}
		{{- end -}}
	`,
	"input": `
		{{- $type := .Field.GetOption "type" -}}
		{{- $checked := and (eq (.Field.GetOption "type").Value "checkbox") (.Field.Data) -}}
		{{- $required := and (.Field.HasOption "required") (.Field.GetOption "required").Value -}}
		{{- $value := .Field.Data -}}

		{{- if eq $type.Value "checkbox" -}}
			{{- $value = 1 -}}
		{{- end -}}

		<input id="{{ .Field.GetId }}" {{ if $checked }}checked{{ end }} {{ if $required }}required="required"{{ end }} name="{{ .Field.GetName }}" value="{{ $value }}" type="{{ $type.Value }}" {{ form_widget_attr .Field }}>
	`,
	"textarea": `
		<textarea id="{{ .Field.GetId }}" {{ if .Field.HasOption "required" }}{{ if (.Field.GetOption "required").Value }}required="required"{{ end }}{{ end }} name="{{ .Field.GetName }}" {{ form_widget_attr .Field }}>{{ .Field.Data }}</textarea>
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
				<input value="" {{ if not $field.Data }}checked{{ end }} name="{{ $field.GetName }}" type="radio" id="{{ $field.GetId }}-0">
				<label for="{{ $field.GetId }}-0">None</label>
			{{- end -}}

			{{- range $key, $choice := $choices.GetChoices -}}
				<input name="{{ $field.GetName }}" type="{{ if $isMultiple }}checkbox{{ else }}radio{{ end }}" value="{{ $choice.Value }}" {{ if $choices.Match $field $choice.Value }}checked{{ end }} id="{{ $field.GetId }}-{{ sum $key $keyAdd }}">
				<label for="{{ $field.GetId }}-{{ sum $key $keyAdd }}">{{- $choice.Label -}}</label>
			{{- end -}}
		{{- else -}}
			<select id="{{ .Field.GetId }}" {{ if $required }}required="required"{{ end }} {{ if $isMultiple }}multiple{{ end }} name="{{ .Field.GetName }}" {{ form_widget_attr .Field }}>
				{{- if not $required -}}
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
			<ul class="form-errors">
				{{- range $error := .Errors -}}
					<li class="form-error">{{- $error -}}</li>
				{{- end -}}
			</ul>
		{{- end -}}
	`,
	"row": `<div class="row">
		{{ $labelAfterWidget := and (.Field.HasOption "type") (eq (.Field.GetOption "type").Value "checkbox") }}

		{{ if and (eq (len .Field.Children) 0) (not $labelAfterWidget) }}
			{{- form_label .Field -}}
		{{ end }}

		{{- form_error nil .Field -}}
		{{- form_widget .Field -}}

		{{ if and (eq (len .Field.Children) 0) ($labelAfterWidget) }}
			{{- form_label .Field -}}
		{{ end }}

		{{- form_widget_help .Field -}}
	</div>`,
}
