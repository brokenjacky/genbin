package binTemplate

var newTpl = `
func New{{.StructName}}() *{{.StructName}} {
    return &{{.StructName}}{
    }
}
`

// `
// func New{{.StructName}}() *{{.StructName}} {
//     if header == nil {
//         header = &Header{
//         //    CommandID: commandid.BindTransceiver,  // TODO
//         }
//     }
//     return &{{.StructName}}{
//         Header: *header,
//     }
// }
// `

func init() {
    Register("new", newTpl)
}
