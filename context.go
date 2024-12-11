package fp

import "context"

type TypedKey[T any] interface {
	ZeroValue() T
}

type ValueType[T any] struct {
}

func (r ValueType[T]) ZeroValue() T {
	return Zero[T]()
}

func WithContextValue[T any](ctx context.Context, key TypedKey[T], value T) context.Context {
	return context.WithValue(ctx, key, value)
}

func GetContextValue[T any](ctx context.Context, key TypedKey[T]) Option[T] {
	if ret, ok := ctx.Value(key).(T); ok {
		return Some(ret)
	}
	return None[T]()
}
