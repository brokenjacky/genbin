package parse

import (
    "reflect"
    "errors"
    "strings"
    "strconv"
)

func Parse(st interface{}) (*MetaData, error) {
    var MetaData MetaData
    kind := reflect.TypeOf(st).Kind()
    v := reflect.ValueOf(st)
    t := reflect.TypeOf(st)
    if kind == reflect.Ptr {
        v = reflect.ValueOf(st).Elem()
        t = reflect.TypeOf(st).Elem()
    }

    strs := strings.Split(t.String(), ".")
    if len(strs) >= 2 {
        MetaData.StructName = strs[1]
    } else {
        MetaData.StructName = t.String()
    }

    if v.Kind() != reflect.Struct {
        return nil, errors.New(" must be a struct")
    }

    for i := 0; i < v.NumField(); i++ {
        tag := t.Field(i).Tag
        var maxLen, maxValue int
        maxLenStr := tag.Get("maxlen")
        if len(maxLenStr) > 0 {
            maxLen, _ = strconv.Atoi(maxLenStr)
        }

        maxValueStr := tag.Get("maxvalue")
        if len(maxValueStr) > 0 {
            maxValue, _ = strconv.Atoi(maxValueStr)
        }

        field := Field{
            Name:     t.Field(i).Name,
            Skip:     tag.Get("gen") == "-",
            Type:     v.Field(i).Type().Name(),
            Kind:     v.Field(i).Kind().String(),
            MaxLen:   maxLen,
            MaxValue: maxValue,
            Refer:    tag.Get("refer"),
            Comment:  tag.Get("comment"),
            Fields:   nil,
        }

        if len(field.Refer) > 0 {
            strs := strings.Split(field.Refer, ".")
            if len(strs) > 1 {
                MetaData.Imports = append(MetaData.Imports, strs[0])
            }
        }

        if v.Field(i).Kind() == reflect.Struct && !field.Skip {
            t, er := Parse(v.Field(i).Interface())
            if er != nil {
                return nil, er
            }

            field.Fields = t.Fields
        }

        MetaData.Fields = append(MetaData.Fields, field)
    }
    return &MetaData, nil
}
