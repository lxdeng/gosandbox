package shared

import (
	"errors"
)

// Every method that we want to export must have
// (1) the method has two arguments, both exported (or builtin) types
// (2) the method's second argument is a pointer
// (3) the method has return type error

type ArithImpl int

func (t *ArithImpl) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B

	// can check overflow and return error
	return nil
}

func (t *ArithImpl) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
