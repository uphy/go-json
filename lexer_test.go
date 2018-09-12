package json

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseObject(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    *Object
		wantErr bool
	}{
		{
			name: "empty object",
			args: `{}`,
			want: &Object{
				members: []Member{},
			},
		},
		{
			name: "one member",
			args: `{"a":1}`,
			want: &Object{
				members: []Member{
					{
						Key:   "a",
						Value: Value{int64(1)},
					},
				},
			},
		},
		{
			name: "two members",
			args: `{"a":1,"b":true}`,
			want: &Object{
				members: []Member{
					{
						Key:   "a",
						Value: Value{int64(1)},
					},
					{
						Key:   "b",
						Value: Value{true},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseObject(strings.NewReader(tt.args))
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseArray(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    Array
		wantErr bool
	}{
		{
			name: "empty array",
			args: `[]`,
			want: Array{},
		},
		{
			name: "one element",
			args: `[0]`,
			want: Array{
				Value{int64(0)},
			},
		},
		{
			name: "two elements",
			args: `[0,null]`,
			want: Array{
				Value{int64(0)},
				Value{nil},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseArray(strings.NewReader(tt.args))
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
