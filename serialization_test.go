package structbot

import "testing"

func TestConfigureFileType_String(t *testing.T) {
	tests := []struct {
		name string
		c    SerializationType
		want string
	}{
		{
			name: "Json",
			c:    Json,
			want: "json",
		}, {
			name: "Xml",
			c:    Xml,
			want: "xml",
		}, {
			name: "Env",
			c:    Env,
			want: "env",
		}, {
			name: "Yaml",
			c:    Yaml,
			want: "yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.getTagString(); got != tt.want {
				t.Errorf("getTagString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSerializationType_String(t *testing.T) {
	tests := []struct {
		name string
		s    SerializationType
		want string
	}{
		{
			name: "Json",
			s:    Json,
			want: "Json",
		}, {
			name: "Xml",
			s:    Xml,
			want: "Xml",
		}, {
			name: "Env",
			s:    Env,
			want: "Env",
		}, {
			name: "Yaml",
			s:    Yaml,
			want: "Yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
