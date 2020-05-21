package main

import "testing"

func Test_snake2Camel(t *testing.T) {
	tests := []struct {
		name, want string
	}{
		{"", ""},
		{"id", "ID"},
		{"user_id", "UserID"},
		{"user_id_and_url", "UserIDAndURL"},
		{"http_server", "HTTPServer"},
	}

	for _, tt := range tests {
		got := snake2Camel(tt.name)
		if got != tt.want {
			t.Errorf("Test_snake2Camel failed. got:%v, want:%v", got, tt.want)
		}
	}
}

func TestStructGenerator_String(t *testing.T) {
	tests := []StructGenerator{
		{
			Name:    "TestStruct",
			Comment: "this is a test struct",
			Fields: []Field{
				{
					Name: "ID",
					Type: "int64",
					Tags: []string{
						`json:"id"`,
						`test_tag:"id"`,
					},
					Comment: `test comment`,
				},
				{
					Name: "LongTestColumnFieldName",
					Type: "string",
					Tags: []string{
						`json:"long_test_column_field_name"`,
						`test_tag:"long_test_column_field_name"`,
					},
					Comment: `test comment xxxxxxxx`,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Log("\n", tt.String())
	}
}
