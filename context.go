package fp

import "context"

type TypedKey[T any] interface {
	ZeroValue() T
	comparable
}

type ValueType[T any] struct {
}

func (r ValueType[T]) ZeroValue() T {
	return Zero[T]()
}

func WithContextValue[K TypedKey[T], T any](ctx context.Context, value T) context.Context {
	var key K
	return context.WithValue(ctx, key, value)
}

func GetContextValue[K TypedKey[T], T any](ctx context.Context) Option[T] {
	var key K

	if ret, ok := ctx.Value(key).(T); ok {
		return Some(ret)
	}
	return None[T]()
}
