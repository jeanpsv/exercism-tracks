package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
	numberOfCows int
}

func (err *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", err.numberOfCows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()

	switch {
	case cows < 0:
		return 0, &SillyNephewError{numberOfCows: cows}
	case cows == 0:
		return 0, errors.New("division by zero")
	case fodder < 0 && (err == ErrScaleMalfunction || err == nil):
		return 0, errors.New("negative fodder")
	case fodder < 0:
		return 0, errors.New("non-scale error")
	case err == ErrScaleMalfunction && fodder > 0:
		fodder = fodder * 2
	case err != nil:
		return 0, err
	}
	return fodder / float64(cows), nil
}
