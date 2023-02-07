package parse

import (
    "text/template"
    "fmt"
    "bytes"
    "genbin/binTemplate"
    "strings"
    "os"
)

func Generate(meta *MetaData, options ...OptionFunc) error {
    option := &Option{}
    for _, o := range options {
        o(option)
    }

    var result strings.Builder

    for _, k := range option.templateName {
        binTpl, ok := binTemplate.Get(k)
        if !ok {
            continue
        }

        str := genStruct(binTpl, meta, option)
        result.WriteString(str)
    }
    // fmt.Println(result.String())

    return writeFile(meta, result.String(), option)
}

func writeFile(tpl *MetaData, str string, o *Option) error {
    _, err := os.Stat(tpl.Package)
    if err != nil {
        if os.IsNotExist(err) {
            os.Mkdir(tpl.Package, 644)
        }
    }

    filePath := fmt.Sprintf("%s/%s.go", tpl.Package, tpl.StructName)
    if len(o.fileName) > 0 {
        filePath = fmt.Sprintf("%s/%s", tpl.Package, o.fileName)
    }

    return os.WriteFile(filePath, []byte(str), 644)
}

func genStruct(binTpl string, tpl *MetaData, o *Option) string {
    temp, err := template.New("struct").Funcs(o.funcMap).Parse(binTpl)
    if err != nil {
        fmt.Println(err.Error())
        return ""
    }

    var buff bytes.Buffer
    err = temp.Execute(&buff, tpl)
    if err != nil {
        fmt.Println(err.Error())
        return ""
    }

    return buff.String()
}
