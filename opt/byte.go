package opt

import (
	"bytes"
	"encoding/json"
)

// Byte ...
type Byte struct {
	v      *byte
	hasKey bool
}

// UnmarshalJSON ...
func (opt *Byte) UnmarshalJSON(b []byte) error {
	opt.hasKey = true
	if bytes.Equal(b, nullLiteral) {
		return nil
	}
	var v byte
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	opt.v = &v
	return nil
}

// MarshalJSON ..
func (opt Byte) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(opt.v)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Value ...
func (opt Byte) Value() *byte {
	return opt.v
}

// HasKey ...
func (opt Byte) HasKey() bool {
	return opt.hasKey
}
