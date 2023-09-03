//revive:disable
package foo

import (
	"context"
)

// AllowedBeforeType is a type that is configured to be allowed before context.Context
type AllowedBeforeType string

// AllowedBeforeStruct is a type that is configured to be allowed before context.Context
type AllowedBeforeStruct struct{}

// AllowedBeforePtrStruct is a type that is configured to be allowed before context.Context
type AllowedBeforePtrStruct struct{}

// A proper context.Context location
//revive:enable:context-as-argument
func x(ctx context.Context, s string) { // ok
}

// An invalid context.Context location with more than 2 args
//revive:enable:context-as-argument
func y(ctx context.Context, s string, r int, x int) { // MATCH /context.Context should be the first parameter of a function/
}
