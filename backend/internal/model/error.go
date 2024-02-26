package model

type constError string

func (err constError) Error() string {
	return string(err)
}

const (
	ErrIncorrectPassword = constError("incorrect password")
)
