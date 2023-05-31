package opt

import (
	"bytes"
	"encoding/json"
)

// Int64 ...
type Int64 struct {
	v      *int64
	hasKey bool
}

// UnmarshalJSON ...
func (opt *Int64) UnmarshalJSON(b []byte) error {
	opt.hasKey = true
	if bytes.Equal(b, nullLiteral) {
		return nil
	}
	var v int64
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.v = &v
	return nil
}

// MarshalJSON ..
func (opt Int64) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(opt.v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Value ...
func (opt Int64) Value() *int64 {
	return opt.v
}

// HasKey ...
func (opt Int64) HasKey() bool {
	return opt.hasKey
}
