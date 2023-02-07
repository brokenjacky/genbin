package binTemplate

var modTpl = `
module {{.Package}}
go 1.18

`

func init() {
    Register("mod", modTpl)
}
