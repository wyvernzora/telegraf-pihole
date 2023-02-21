package scanner

// A Scanner is a function that takes a dynamic value and attempts to convert it to the
// output type. By default, it returns errors when input is nil or not of the expected type.
type Scanner[T any] func(*T, interface{}) error

// New returns a Scanner[T] for type T
func New[T any]() Scanner[T] {
	return func(out *T, value interface{}) error {
		if value == nil {
			return nilValueError()
		}
		if t, ok := value.(T); ok {
			*out = t
			return nil
		}
		return unexpectedTypeError[T](value)
	}
}

// WithDefault modifies a Scanner instance to instead of returning error on nil input, assign
// the specified default value to the output.
func WithDefault[T any](defaultValue T, scanner Scanner[T]) Scanner[T] {
	return func(out *T, value interface{}) error {
		err := scanner(out, value)
		if isNilValueError(err) {
			*out = defaultValue
			return nil
		}
		return err
	}
}

// WithMapping modifies a Scanner instance to instead of returning result directly, look up
// in the provided map and return the map value.
// Returns an error if the map does not contain the expected value.
func WithMapping[T comparable, K any](mapping map[T]K, scanner Scanner[T]) Scanner[K] {
	return func(out *K, value interface{}) error {
		var key T
		if err := scanner(&key, value); err != nil {
			return err
		}
		if result, ok := mapping[key]; ok {
			*out = result
			return nil
		}
		return unexpectedValueError(key)
	}
}
