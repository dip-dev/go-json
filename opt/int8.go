package opt

import (
	"bytes"
	"encoding/json"
)

// Int8 ...
type Int8 struct {
	v      *int8
	hasKey bool
}

// UnmarshalJSON ...
func (opt *Int8) UnmarshalJSON(b []byte) error {
	opt.hasKey = true
	if bytes.Equal(b, nullLiteral) {
		return nil
	}
	var v int8
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.v = &v
	return nil
}

// MarshalJSON ..
func (opt Int8) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(opt.v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Value ...
func (opt Int8) Value() *int8 {
	return opt.v
}

// HasKey ...
func (opt Int8) HasKey() bool {
	return opt.hasKey
}
