package clitypes

import "fmt"

type OptionalBool struct {
	IsSet bool
	Value bool
}

func (b *OptionalBool) Set(s string) error {
	b.IsSet = true
	b.Value = (s == "true")
	return nil
}

func (b *OptionalBool) String() string {
	if !b.IsSet {
		return ""
	}
	return fmt.Sprintf("%v", b.Value)
}
