package theme

var Html5 = map[string]string{
	"form": `<form action="" method="">
		{{ form_error .Form nil }}
		{{ .Content }}
	</form>`,
	"label": `<label for="">Label</label>`,
	"input": `<input name="{{ .Field.Name }}" value="{{ .Field.Data }}" type="text">`,
	"sub_form": `
		{{ form_label .Field }}

		{{ range $field := .Field.Children }}
			{{ form_row $field }}
		{{ end }}
	`,
	"error": `<div class="error">
		{{ range $error := .Errors }}
			{{ $error }}<br>
		{{ end }}
	</div>`,
	"row": `<div class="row">
		{{ form_label .Field }}
		{{ form_error nil .Field }}
		{{ form_widget .Field }}
	</div>`,
}
