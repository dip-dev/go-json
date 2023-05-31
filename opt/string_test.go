package opt

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestStringUnmarshalJSON ...
func TestStringUnmarshalJSON(t *testing.T) {

	type testStruct struct {
		Req string `json:"req"`
		Opt String `json:"opt"`
	}

	toPtr := func(s string) *string {
		return &s
	}

	success := map[string]struct {
		data []byte
		want testStruct
	}{
		"has key": {
			data: []byte(`{"req":"req string", "opt":"opt string"}`),
			want: testStruct{
				Req: "req string",
				Opt: String{v: toPtr("opt string"), hasKey: true},
			},
		},
		"has key(is initial value)": {
			data: []byte(`{"req":"", "opt":""}`),
			want: testStruct{
				Req: "",
				Opt: String{v: toPtr(""), hasKey: true},
			},
		},
		"has key(is null)": {
			data: []byte(`{"req":"", "opt":null}`),
			want: testStruct{
				Req: "",
				Opt: String{hasKey: true},
			},
		},
		"has not key": {
			data: []byte(`{"req":"req string"}`),
			want: testStruct{
				Req: "req string",
				Opt: String{}, // hasKey is false
			},
		},
	}
	t.Run("success", func(t *testing.T) {
		for tt, tp := range success {
			t.Run(tt, func(t *testing.T) {
				var got testStruct
				err := json.Unmarshal(tp.data, &got)
				assert.NoError(t, err)
				assert.Equal(t, tp.want, got)
			})
		}
	})

	fail := map[string]struct {
		data []byte
	}{
		"cannot unmarshal": {
			data: []byte(`{"req":"req string", "opt":100}`),
		},
	}
	t.Run("fail", func(t *testing.T) {
		for tt, tp := range fail {
			t.Run(tt, func(t *testing.T) {
				var got testStruct
				if err := json.Unmarshal(tp.data, &got); err == nil {
					assert.Fail(t, "No error")
				}
			})
		}
	})
}

// TestStringMarshalJSON ...
func TestStringMarshalJSON(t *testing.T) {
	type testStruct struct {
		Req string  `json:"req"`
		Opt *String `json:"opt,omitempty"`
	}

	toPtr := func(s string) *string {
		return &s
	}

	success := map[string]struct {
		data testStruct
		want string
	}{
		"has key": {
			data: testStruct{
				Req: "req string",
				Opt: &String{v: toPtr("opt string"), hasKey: true},
			},
			want: `{"req":"req string", "opt":"opt string"}`,
		},
		"has key(is initial value)": {
			data: testStruct{
				Req: "",
				Opt: &String{v: toPtr(""), hasKey: true},
			},
			want: `{"req":"", "opt":""}`,
		},
		"has key(is null)": {
			data: testStruct{
				Req: "",
				Opt: &String{hasKey: true},
			},
			want: `{"req":"", "opt":null}`,
		},
		"has not key": {
			data: testStruct{
				Req: "req string",
				Opt: &String{}, // hasKey is false
			},
			want: `{"req":"req string", "opt":null}`,
		},
		"has not opt": {
			data: testStruct{
				Req: "req string",
				Opt: nil,
			},
			want: `{"req":"req string"}`,
		},
	}

	t.Run("success", func(t *testing.T) {
		for tt, tp := range success {
			t.Run(tt, func(t *testing.T) {
				got, err := json.Marshal(tp.data)
				assert.NoError(t, err)
				require.JSONEq(t, tp.want, string(got))
			})
		}
	})

}
