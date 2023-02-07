package parse

import "text/template"

type Option struct {
    templateName []string
    funcMap      template.FuncMap
    fileName     string
}

type OptionFunc func(option *Option)

func WithTemplateName(names []string) OptionFunc {
    return func(option *Option) {
        option.templateName = names
    }
}

func WithFileName(name string) OptionFunc {
    return func(option *Option) {
        option.fileName = name
    }
}

func WithFuncMap(fmap template.FuncMap) OptionFunc {
    return func(option *Option) {
        option.funcMap = fmap
    }
}
