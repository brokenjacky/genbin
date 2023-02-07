package parse

type MetaData struct {
    Package    string  `json:"package,omitempty"`
    StructName string  `json:"structName,omitempty"`
    Fields     []Field `json:"fields,omitempty"`
    Imports    []string
}

type Field struct {
    Name     string  `json:"name,omitempty"`
    Skip     bool    `json:"skip,omitempty"`
    Kind     string  `json:"type,omitempty"`
    Type     string  `json:"type,omitempty"`
    MaxLen   int     `json:"maxLen,omitempty"`
    MaxValue int     `json:"maxValue,omitempty"`
    Refer    string  `json:"refer,omitempty"` // 自定义数据
    Comment  string  `json:"comment,omitempty"`
    Fields   []Field `json:"fields,omitempty"`
}
