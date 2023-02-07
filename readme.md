### fastbin 
- 可以根据定义的结构体，快速生成二进制解析文件
- 可以自定义模板
- 可以指定生成的文件名和方法名


### how to use
```go

type Test struct {
    CommandLength uint32
    ByteID        byte
    CustomField   string `refer:"CustomField"` // 需要提前定义好
    MaxLen        string `maxlen:"64"`
}

func register(st ...interface{}) {

    genList := []string{"header", "struct", "new", "marshal", "unmarshal"}

    for _, v := range st {

        tpl, err := parse.Parse(v)

        if err != nil {
            fmt.Println("err", err.Error())
            continue
        }

        j, _ := json.Marshal(tpl)
        fmt.Println(string(j))

        tpl.Package = "pdu"

        parse.Generate(tpl, parse.WithTemplateName(genList), parse.WithFuncMap(funcMap))
    }
}

func main() {

    register(Test{})
    tpl := parse.MetaData{
        Package:    "pdu",
        StructName: "log",
    }
    parse.Generate(&tpl, parse.WithTemplateName([]string{tpl.StructName}))
    
    tpl.StructName = "mod"
    parse.Generate(&tpl, parse.WithTemplateName([]string{tpl.StructName}), parse.WithFileName("go.mod"))
}

```
