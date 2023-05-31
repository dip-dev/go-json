package opt

import (
	"bytes"
	"encoding/json"
)

// Float64 ...
type Float64 struct {
	v      *float64
	hasKey bool
}

// UnmarshalJSON ...
func (opt *Float64) UnmarshalJSON(b []byte) error {
	opt.hasKey = true
	if bytes.Equal(b, nullLiteral) {
		return nil
	}
	var v float64
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.v = &v
	return nil
}

// MarshalJSON ..
func (opt Float64) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(opt.v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Value ...
func (opt Float64) Value() *float64 {
	return opt.v
}

// HasKey ...
func (opt Float64) HasKey() bool {
	return opt.hasKey
}
