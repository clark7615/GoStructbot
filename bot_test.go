package structbot

import (
	"testing"
)

func TestBot_MakeStruct(t *testing.T) {
	type args struct {
		out Test
	}
	var tests = []struct {
		name string
		Str string
		args args
	}{
		{
			name: "SetStruct",
			Str: `{"id":1,"text":"test"}`,
			args: args{
				out: Test{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bo := &Bot{}
			bo.MakeStruct(tt.Str,&tt.args.out)
			if tt.args.out.ID != 1{
				t.Fatal("Set struct error")
			}
			if tt.args.out.Text != "test"{
				t.Fatal("Set struct error")
			}
		})
	}
}
