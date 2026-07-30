package types

import (
	"fmt"
	"runtime/debug"
)

// GetPrimitiveTypeByName returns the TypeInfo for a given primitive TypeKind or nil if the kind is not a primitive type.
func GetPrimitiveTypeByName(name string) *TypeInfo {
	if t, ok := builtinTypes[name]; ok {
		return t
	}
	return nil
}

/* Utility functions */

func (t *TypeInfo) checkNotNil() bool {
	if t == nil {
		panic(fmt.Sprintf("TypeInfo is nil. Stack trace:\n%s", debug.Stack()))
	}

	return true
}

// IsUsable checks if the TypeInfo represents an usable type (not nil, not unknown, and not error).
func (t *TypeInfo) IsUsable() bool {
	return t.checkNotNil() && t.Kind != Unknown && t.Kind != Error
}

// IsUnknown checks if the TypeInfo represents an unknown type (needs type inference).
func (t *TypeInfo) IsUnknown() bool {
	return t.checkNotNil() && t.Kind == Unknown
}

// IsKnown checks if the TypeInfo is known (not unknown and not nil).
func (t *TypeInfo) IsKnown() bool {
	return t.checkNotNil() && t.Kind != Unknown
}

// IsError checks if the TypeInfo represents an error type.
func (t *TypeInfo) IsError() bool {
	return t.checkNotNil() && t.Kind == Error
}

// IsValid checks if the TypeInfo is valid (not nil and not error).
func (t *TypeInfo) IsValid() bool {
	return t.checkNotNil() && t.Kind != Error
}

// IsUnit checks if the TypeInfo represents the unit type.
func (t *TypeInfo) IsUnit() bool {
	return t.checkNotNil() && t.Kind == Unit
}

// IsPrimitive checks if the type is a primitive type (bool, int, float, string, char).
func (t *TypeInfo) IsPrimitive() bool {
	return t.checkNotNil() && GetPrimitiveTypeByName(t.Name) != nil
}

// IsNominal checks if the type is a nominal (user-defined) type.
func (t *TypeInfo) IsNominal() bool {
	return t.checkNotNil() && t.Kind == Nominal
}

// IsInt checks if the type is an integer type.
func (t *TypeInfo) IsInt() bool {
	return t.checkNotNil() && t.Kind == Int
}

// IsFloat checks if the type is a floating-point type.
func (t *TypeInfo) IsFloat() bool {
	return t.checkNotNil() && t.Kind == Float
}

// IsNumeric checks if the type is either an integer or a floating-point type.
func (t *TypeInfo) IsNumeric() bool {
	return t.checkNotNil() && (t.Kind == Int || t.Kind == Float)
}

// IsBool checks if the type is a boolean type.
func (t *TypeInfo) IsBool() bool {
	return t.checkNotNil() && t.Kind == Bool
}

// IsString checks if the type is a string type.
func (t *TypeInfo) IsString() bool {
	return t.checkNotNil() && t.Kind == String
}

// IsChar checks if the type is a character type.
func (t *TypeInfo) IsChar() bool {
	return t.checkNotNil() && t.Kind == Char
}
