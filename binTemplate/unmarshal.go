package binTemplate

var unmarshalTpl = `
func ({{.Package}} *{{.StructName}}) Unmarshal(buf *bytes.Buffer) error {

    var (
        line []byte
        err error
        c byte
    )
    _,_,_ = line,err,c
    
{{ range .Fields }}
    {{- if .Skip}} {{continue}} {{- end}}

    {{- if eq .Name .Type}}
    // read {{ .Name }}
    err = {{$.Package}}.{{ .Name }}.Unmarshal(buf)
    if err != nil {
        return err
    }
    {{- end }}

    {{- if eq .Type "string"}}
    // read {{.Name}}
    line, err = buf.ReadBytes(0x00)
    if err != nil {
        return err
    }

    {{- if gt .MaxLen 0 }} 
    if len(line) > {{.MaxLen}} {
        Logger.Infof("decode {{.Name}} failed, lineLen[%d] max[%d]", len(line),{{.MaxLen}})
        return err
    }{{end}}

    if len(line) > 1 {
        {{$.Package}}.{{.Name}} = string(line[0 : len(line)-1])
    }
    {{- else if eq .Type "uint8" }}
    // read {{.Name}}
    c, err = buf.ReadByte()
    if err != nil {
        Logger.Infof("read {{.Name}} err[%s]\n",err.Error())
        return err
    }

    {{- if .Refer}} 
    {{$.Package}}.{{.Name}} = {{.Refer}}(c) 
    {{- else}} 
    {{$.Package}}.{{.Name}} = c 
    {{- end}}

    {{- else if eq .Type "uint32" }}
    {{- if .Refer}}
    {{$.Package}}.{{.Name}} = {{ .Refer }}(BigEndian.Uint32(buf.Next(4)))
    {{- else }}
    {{$.Package}}.{{.Name}} = BigEndian.Uint32(buf.Next(4))
    {{- end }}

    {{- end}}
{{end}}
    return nil
}

`

func init() {
    Register("unmarshal", unmarshalTpl)
}
