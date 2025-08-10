package interfaces

// Parser interface defines strcut behaviours used to parse values to other format. This could be anything like a user input to a struct.
type Parser[T any] interface {
	Parse() (T, error)
}
