package binTemplate

var structTpl = `
type {{.StructName}} struct {
{{- range .Fields }}
    {{- if .Skip}} {{continue}} {{- end}}

    {{- if eq .Name .Type}}
        {{- if eq .Kind "struct"}}

            {{- if .Refer }} 
    {{ .Refer}}     
            {{- else }}
    {{ .Name }} 
            {{- end}} 

        {{- end }}

    {{- else}}
    {{- if .Refer }} 
    {{ .Name }} {{ .Refer }}
    {{- else }}
    {{ .Name }} {{ .Type }} 
    {{- end}}
    {{- if .Comment }} // {{ .Comment }}{{end}}
    {{- end}}
{{- end}}
}

`

func init() {
    Register("struct", structTpl)
}
