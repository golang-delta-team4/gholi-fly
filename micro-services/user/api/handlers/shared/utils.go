package shared

import "context"

type ServiceGetter[T any] func(context.Context) T