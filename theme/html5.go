package theme

var Html5 = map[string]string{
	"form": `<form action="{{ .Form.Action }}" method="{{ .Form.Method }}" {{ form_attr .Form }}>
		{{- form_error .Form nil -}}

		{{- range $field := .Form.Fields -}}
			{{- form_row $field -}}
		{{- end -}}
	</form>`,
	"attributes": `{{ range $key, $value := .Attributes }}{{ $key }}="{{ $value }}"{{ end }}`,
	// "attributes": `{{ if gt (len .Attributes) 0 }}
	// 	{{ range $key, $value := .Attributes }}
	// 		{{ $key }}="{{ $value }}"
	// 	{{ end }}
	// {{ end }}`,
	"label": `
		{{ if .Field.HasOption "label" }}
			{{ $label := (.Field.GetOption "label").Value }}

			{{- if ne $label "" -}}
				<label for="{{ .Field.GetId }}" {{ label_attr .Field }}>{{ $label }}</label>
			{{- end -}}
		{{- end -}}
	`,
	"input": `
		{{ $type := .Field.GetOption "type" }}
		<input id="{{ .Field.GetId }}" {{ if .Field.HasOption "required" }}{{ if (.Field.GetOption "required").Value }}required="required"{{ end }}{{ end }} name="{{ .Field.GetName }}" value="{{ .Field.Data }}" type="{{ $type.Value }}" {{ widget_attr .Field }}>
	`,
	"textarea": `
		<textarea id="{{ .Field.GetId }}" {{ if .Field.HasOption "required" }}{{ if (.Field.GetOption "required").Value }}required="required"{{ end }}{{ end }} name="{{ .Field.GetName }}" {{ widget_attr .Field }}>{{ .Field.Data }}</textarea>
	`,
	"sub_form": `
		{{- range $field := .Field.Children -}}
			{{- form_row $field -}}
		{{- end -}}
	`,
	"error": `
		{{- if gt (len .Errors) 0 -}}
			<ul class="error">
				{{- range $error := .Errors -}}
					<li>{{- $error -}}</li>
				{{- end -}}
			</ul>
		{{- end -}}
	`,
	"row": `<div class="row">
		{{- form_label .Field -}}
		{{- form_error nil .Field -}}
		{{- form_widget .Field -}}
	</div>`,
}
