package opt

import (
	"bytes"
	"encoding/json"
)

// String ...
type String struct {
	v      *string
	hasKey bool
}

// UnmarshalJSON ...
func (opt *String) UnmarshalJSON(b []byte) error {
	opt.hasKey = true
	if bytes.Equal(b, nullLiteral) {
		return nil
	}
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.v = &v
	return nil
}

// MarshalJSON ..
func (opt String) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(opt.v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Value ...
func (opt String) Value() *string {
	return opt.v
}

// HasKey ...
func (opt String) HasKey() bool {
	return opt.hasKey
}
