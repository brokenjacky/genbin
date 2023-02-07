package binTemplate

var marshalTpl = `
func ({{.Package}} *{{.StructName}}) Marshal() []byte {
    var b bytes.Buffer

{{- range .Fields }}
    {{- if .Skip}} {{continue}} {{- end}}

    {{- test .Name -}}
    {{- if eq .Name .Type}}

    b.Write({{$.Package}}.{{ .Name }}.Marshal())

    {{- else}}
        {{- if eq .Kind "string" }}

    b.WriteString({{$.Package}}.{{ .Name }})
    b.WriteByte(0x00)

    {{- else if eq .Kind "uint8" }}

    {{- if .Refer}} 
    b.WriteByte({{.Type}}({{$.Package}}.{{ .Name }}))
    {{- else}} 
    b.WriteByte({{$.Package}}.{{ .Name }})
    {{- end}}

    {{- else if eq .Kind "uint32" }}
    Write(&b,BigEndian,{{$.Package}}.{{ .Name }})

        {{- end}}
    {{- end}}
{{- end}}

    return b.Bytes()
}

`

func init() {
    Register("marshal", marshalTpl)
}
