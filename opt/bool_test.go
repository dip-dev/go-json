package opt

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestBoolUnmarshalJSON ...
func TestBoolUnmarshalJSON(t *testing.T) {

	type testStruct struct {
		Req bool `json:"req"`
		Opt Bool `json:"opt"`
	}

	toPtr := func(b bool) *bool {
		return &b
	}

	success := map[string]struct {
		data []byte
		want testStruct
	}{
		"has key": {
			data: []byte(`{"req":true, "opt":true}`),
			want: testStruct{
				Req: true,
				Opt: Bool{v: toPtr(true), hasKey: true},
			},
		},
		"has key(is initial value)": {
			data: []byte(`{"req":false, "opt":false}`),
			want: testStruct{
				Req: false,
				Opt: Bool{v: toPtr(false), hasKey: true},
			},
		},
		"has key(is null)": {
			data: []byte(`{"req":false, "opt":null}`),
			want: testStruct{
				Req: false,
				Opt: Bool{hasKey: true},
			},
		},
		"has not key": {
			data: []byte(`{"req":false}`),
			want: testStruct{
				Req: false,
				Opt: Bool{}, // hasKey is false
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
			data: []byte(`{"req":"req bool", "opt":100}`),
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

// TestBoolMarshalJSON ...
func TestBoolMarshalJSON(t *testing.T) {
	type testStruct struct {
		Req bool  `json:"req"`
		Opt *Bool `json:"opt,omitempty"`
	}

	toPtr := func(b bool) *bool {
		return &b
	}

	success := map[string]struct {
		data testStruct
		want string
	}{
		"has key": {
			data: testStruct{
				Req: true,
				Opt: &Bool{v: toPtr(true), hasKey: true},
			},
			want: `{"req":true, "opt":true}`,
		},
		"has key(is initial value)": {
			data: testStruct{
				Req: false,
				Opt: &Bool{v: toPtr(false), hasKey: true},
			},
			want: `{"req":false, "opt":false}`,
		},
		"has key(is null)": {
			data: testStruct{
				Req: false,
				Opt: &Bool{hasKey: true},
			},
			want: `{"req":false, "opt":null}`,
		},
		"has not key": {
			data: testStruct{
				Req: false,
				Opt: &Bool{}, // hasKey is false
			},
			want: `{"req":false, "opt":null}`,
		},
		"has not opt": {
			data: testStruct{
				Req: false,
				Opt: nil,
			},
			want: `{"req":false}`,
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
