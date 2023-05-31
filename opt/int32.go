package opt

import (
	"bytes"
	"encoding/json"
)

// Int32 ...
type Int32 struct {
	v      *int32
	hasKey bool
}

// UnmarshalJSON ...
func (opt *Int32) UnmarshalJSON(b []byte) error {
	opt.hasKey = true
	if bytes.Equal(b, nullLiteral) {
		return nil
	}
	var v int32
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.v = &v
	return nil
}

// MarshalJSON ..
func (opt Int32) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(opt.v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Value ...
func (opt Int32) Value() *int32 {
	return opt.v
}

// HasKey ...
func (opt Int32) HasKey() bool {
	return opt.hasKey
}
