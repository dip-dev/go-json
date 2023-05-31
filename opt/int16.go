package opt

import (
	"bytes"
	"encoding/json"
)

// Int16 ...
type Int16 struct {
	v      *int16
	hasKey bool
}

// UnmarshalJSON ...
func (opt *Int16) UnmarshalJSON(b []byte) error {
	opt.hasKey = true
	if bytes.Equal(b, nullLiteral) {
		return nil
	}
	var v int16
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.v = &v
	return nil
}

// MarshalJSON ..
func (opt Int16) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(opt.v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Value ...
func (opt Int16) Value() *int16 {
	return opt.v
}

// HasKey ...
func (opt Int16) HasKey() bool {
	return opt.hasKey
}
