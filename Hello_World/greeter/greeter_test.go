package greeter_test

import (
	"hello_world/greeter"
	"testing"
)

func TestGreet(t *testing.T) {
	type testCase struct {
		lang     greeter.Language
		want     string
		wantErr  bool
		errValue string
	}

	var tests = map[string]testCase{
		"English": {
			lang: "en",
			want: "Hello world",
		},
		"French": {
			lang: "fr",
			want: "Bonjour le monde",
		},
		"Spanish": {
			lang: "es",
			want: "Hola mundo",
		},
		"Latin": {
			lang: "la",
			want: "Salve mundi",
		},
		"Empty": {
			lang:     "",
			want:     "Unsupported language",
			wantErr:  true,
			errValue: `"": unsupported language`,
		},
		"UnknownLang": {
			lang:     "jp",
			want:     "Unsupported language",
			wantErr:  true,
			errValue: `"jp": unsupported language`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := greeter.Greet(test.lang)

			if got != test.want {
				t.Errorf("Expected greeting: %q, got: %q", test.want, got)
			}

			if test.wantErr {
				if err == nil {
					t.Errorf("Expected error but got none")
				} else if err.Error() != test.errValue {
					t.Errorf("Expected error: %q, got: %q", test.errValue, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect error, but got: %v", err)
				}
			}
		})
	}
}
