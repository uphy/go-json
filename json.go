package json

import (
	"bytes"
	"fmt"
)

type (
	Object struct {
		members []Member
	}
	Array  []Value
	Member struct {
		Key   string
		Value Value
	}
	Value struct {
		v interface{}
	}
)

func (o *Object) Get(name string) *Value {
	for _, m := range o.members {
		if m.Key == name {
			return &m.Value
		}
	}
	return nil
}

func (o *Object) Remove(name string) *Value {
	for i, m := range o.members {
		if m.Key == name {
			o.members = append(o.members[:i], o.members[i+1:]...)
			return &m.Value
		}
	}
	return nil
}

func (v *Value) Interface() interface{} {
	return v.v
}

func (v *Value) isNull() bool {
	return v.v == nil
}

func (v Value) String() string {
	if v.v == nil {
		return "null"
	}
	switch casted := v.v.(type) {
	case string:
		return fmt.Sprintf(`"%s"`, casted)
	}
	return fmt.Sprint(v.v)
}

func (a Array) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString("[")
	for i, e := range a {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(e.String())
	}
	buf.WriteString("]")
	return buf.String()
}

func (o Object) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString("{")
	for i, e := range o.members {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteRune('"')
		buf.WriteString(e.Key)
		buf.WriteRune('"')

		buf.WriteString(":")
		buf.WriteString(e.Value.String())
	}
	buf.WriteString("}")
	return buf.String()
}
