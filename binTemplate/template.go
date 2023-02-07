package binTemplate

var templateMap = map[string]string{}

func Register(name string, tpl string) {
    templateMap[name] = tpl
}

func Get(name string) (string, bool) {
    str, ok := templateMap[name]
    return str, ok
}

func GetAllFunc() []string {
    var result []string
    for k, _ := range templateMap {
        result = append(result, k)
    }

    return result
}
