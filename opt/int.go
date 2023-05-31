package opt

import (
	"bytes"
	"encoding/json"
)

// Int ...
type Int struct {
	v      *int
	hasKey bool
}

// UnmarshalJSON ...
func (opt *Int) UnmarshalJSON(b []byte) error {
	opt.hasKey = true
	if bytes.Equal(b, nullLiteral) {
		return nil
	}
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.v = &v
	return nil
}

// MarshalJSON ..
func (opt Int) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(opt.v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Value ...
func (opt Int) Value() *int {
	return opt.v
}

// HasKey ...
func (opt Int) HasKey() bool {
	return opt.hasKey
}
