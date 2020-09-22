package structbot

import "testing"

type TestStruct struct {
	ID   int    `json:"id" yaml:"id" xml:"id"`
	Data string `json:"data" yaml:"data" xml:"data"`
}

func TestBot_MakeStruct(t *testing.T) {
	type args struct {
		str interface{}
		out interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Yaml make struct test",
			args: args{
				str: []byte(`
id: 1
data: test
`),
				out: &TestStruct{},
			},
			wantErr: false,
		}, {
			name: "Json make struct test",
			args: args{
				str: []byte(`{"id":1,"data":"test"}`),
				out: &TestStruct{},
			},
			wantErr: false,
		}, {
			name: "Xml make struct test",
			args: args{
				str: `
<root>
   <data>test</data>
   <id>1</id>
</root>`,
				out: &TestStruct{},
			},
			wantErr: false,
		}, {
			name: "Map make struct test",
			args: args{
				str: map[string]interface{}{
					"id":   1,
					"data": "test",
				},
				out: &TestStruct{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := MakeStruct(tt.args.str, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("MakeStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
			check := tt.args.out.(*TestStruct)
			if check.ID != 1 || check.Data != "test" {
				t.Errorf("MakeStruct() error ID = %v, want %v ,Data = %v want %v",
					check.ID, 1, check.Data, "test")
			}
		})
	}
}
