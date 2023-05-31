package opt

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestIntUnmarshalJSON ...
func TestIntUnmarshalJSON(t *testing.T) {

	type testStruct struct {
		Req int `json:"req"`
		Opt Int `json:"opt"`
	}

	toPtr := func(s int) *int {
		return &s
	}

	success := map[string]struct {
		data []byte
		want testStruct
	}{
		"has key": {
			data: []byte(`{"req":100, "opt":200}`),
			want: testStruct{
				Req: 100,
				Opt: Int{v: toPtr(200), hasKey: true},
			},
		},
		"has key(is initial value)": {
			data: []byte(`{"req":0, "opt":0}`),
			want: testStruct{
				Req: 0,
				Opt: Int{v: toPtr(0), hasKey: true},
			},
		},
		"has key(is null)": {
			data: []byte(`{"req":0, "opt":null}`),
			want: testStruct{
				Req: 0,
				Opt: Int{hasKey: true},
			},
		},
		"has not key": {
			data: []byte(`{"req":100}`),
			want: testStruct{
				Req: 100,
				Opt: Int{}, // hasKey is false
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
			data: []byte(`{"req":"req string", "opt":"string"}`),
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

// TestIntMarshalJSON ...
func TestIntMarshalJSON(t *testing.T) {
	type testStruct struct {
		Req int  `json:"req"`
		Opt *Int `json:"opt,omitempty"`
	}

	toPtr := func(i int) *int {
		return &i
	}

	success := map[string]struct {
		data testStruct
		want string
	}{
		"has key": {
			data: testStruct{
				Req: 100,
				Opt: &Int{v: toPtr(200), hasKey: true},
			},
			want: `{"req":100, "opt":200}`,
		},
		"has key(is initial value)": {
			data: testStruct{
				Req: 0,
				Opt: &Int{v: toPtr(0), hasKey: true},
			},
			want: `{"req":0, "opt":0}`,
		},
		"has key(is null)": {
			data: testStruct{
				Req: 0,
				Opt: &Int{hasKey: true},
			},
			want: `{"req":0, "opt":null}`,
		},
		"has not key": {
			data: testStruct{
				Req: 100,
				Opt: &Int{}, // hasKey is false
			},
			want: `{"req":100, "opt":null}`,
		},
		"has not opt": {
			data: testStruct{
				Req: 100,
				Opt: nil,
			},
			want: `{"req":100}`,
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
