package opt

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestFloat64UnmarshalJSON ...
func TestFloat64UnmarshalJSON(t *testing.T) {

	type testStruct struct {
		Req float64 `json:"req"`
		Opt Float64 `json:"opt"`
	}

	toPtr := func(f float64) *float64 {
		return &f
	}

	success := map[string]struct {
		data []byte
		want testStruct
	}{
		"has key": {
			data: []byte(`{"req":100.12, "opt":200.23}`),
			want: testStruct{
				Req: 100.12,
				Opt: Float64{v: toPtr(200.23), hasKey: true},
			},
		},
		"has key(is initial value)": {
			data: []byte(`{"req":0, "opt":0}`),
			want: testStruct{
				Req: 0,
				Opt: Float64{v: toPtr(0), hasKey: true},
			},
		},
		"has key(is null)": {
			data: []byte(`{"req":0, "opt":null}`),
			want: testStruct{
				Req: 0,
				Opt: Float64{hasKey: true},
			},
		},
		"has not key": {
			data: []byte(`{"req":100.12}`),
			want: testStruct{
				Req: 100.12,
				Opt: Float64{}, // hasKey is false
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

// TestFloat64MarshalJSON ...
func TestFloat64MarshalJSON(t *testing.T) {
	type testStruct struct {
		Req float64  `json:"req"`
		Opt *Float64 `json:"opt,omitempty"`
	}

	toPtr := func(f float64) *float64 {
		return &f
	}

	success := map[string]struct {
		data testStruct
		want string
	}{
		"has key": {
			data: testStruct{
				Req: 100.12,
				Opt: &Float64{v: toPtr(200.23), hasKey: true},
			},
			want: `{"req":100.12, "opt":200.23}`,
		},
		"has key(is initial value)": {
			data: testStruct{
				Req: 0,
				Opt: &Float64{v: toPtr(0), hasKey: true},
			},
			want: `{"req":0, "opt":0}`,
		},
		"has key(is null)": {
			data: testStruct{
				Req: 0,
				Opt: &Float64{hasKey: true},
			},
			want: `{"req":0, "opt":null}`,
		},
		"has not key": {
			data: testStruct{
				Req: 100.12,
				Opt: &Float64{}, // hasKey is false
			},
			want: `{"req":100.12, "opt":null}`,
		},
		"has not opt": {
			data: testStruct{
				Req: 100.12,
				Opt: nil,
			},
			want: `{"req":100.12}`,
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
