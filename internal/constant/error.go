package constant

import "errors"

var (
	// ErrRequiredDB will throw if database address is empty.
	ErrRequiredDB = errors.New("required database address")
	// ErrInvalidDB will throw if database address format is invalid.
	ErrInvalidDB = errors.New("invalid database address")
	// ErrRequiredUser will throw if user id is empty.
	ErrRequiredUser = errors.New("required user id")
	// ErrRequiredName will throw if name is empty.
	ErrRequiredName = errors.New("required item name")
	// ErrInvalidTaxCode will throw if tax code is invalid.
	ErrInvalidTaxCode = errors.New("invalid tax code")
	// ErrInvalidPrice will throw if price is 0 or negative.
	ErrInvalidPrice = errors.New("price must be positive")
	// ErrRequiredID will throw if item id is empty.
	ErrRequiredID = errors.New("required id")
	// ErrNotFound will throw if data not found.
	ErrNotFound = errors.New("data not found")
)
