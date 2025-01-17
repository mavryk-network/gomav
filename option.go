package gomav

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/ecadlabs/pretty"
	"github.com/mavryk-network/gomav/v2/encoding"
)

type Option[T any] struct {
	some  bool
	value T
}

func Some[T any](val T) Option[T] {
	return Option[T]{
		some:  true,
		value: val,
	}
}

func None[T any]() Option[T] {
	return Option[T]{
		some: false,
	}
}

func (op Option[T]) Unwrap() T {
	if !op.some {
		panic(fmt.Sprintf("called `Unwrap()` on a `None` value of type %T", op))
	}
	return op.value
}

func (op *Option[T]) UnwrapPtr() *T {
	if !op.some {
		panic(fmt.Sprintf("called `Unwrap()` on a `None` value of type %T", op))
	}
	return &op.value
}

func (op Option[T]) CheckUnwrap() (T, bool) {
	if !op.some {
		var null T
		return null, false
	}
	return op.value, true
}

func (op *Option[T]) CheckUnwrapPtr() (*T, bool) {
	if !op.some {
		return nil, false
	}
	return &op.value, true
}

func (op Option[T]) UnwrapUnchecked() T {
	return op.value
}

func (op Option[T]) IsSome() bool { return op.some }
func (op Option[T]) IsNone() bool { return !op.some }

func (op Option[T]) Or(val Option[T]) Option[T] {
	if op.some {
		return op
	}
	return val
}

func (op Option[T]) OrElse(f func() Option[T]) Option[T] {
	if op.some {
		return op
	}
	return f()
}

func (op Option[T]) UnwrapOr(def T) T {
	if op.some {
		return op.value
	}
	return def
}

func (op Option[T]) UnwrapOrElse(f func() T) T {
	if op.some {
		return op.value
	}
	return f()
}

func (op Option[T]) UnwrapOrZero() T {
	if op.some {
		return op.value
	}
	var t T
	return t
}

func (op *Option[T]) DecodeMV(data []byte, ctx *encoding.Context) (rest []byte, err error) {
	if len(data) == 0 {
		// tail entry
		*op = Option[T]{}
		return data, nil
	}

	*op = Option[T]{
		some: data[0] != 0,
	}
	data = data[1:]

	if op.some {
		return encoding.Decode(data, &op.value, encoding.Ctx(ctx))
	}
	return data, nil
}

func (op *Option[T]) EncodeMV(ctx *encoding.Context) ([]byte, error) {
	var buf bytes.Buffer
	if op.IsSome() {
		buf.WriteByte(255)
		val := op.Unwrap()
		if err := encoding.Encode(&buf, &val, encoding.Ctx(ctx)); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	} else {
		return []byte{0}, nil
	}
}

func (op *Option[T]) MarshalJSON() ([]byte, error) {
	var v *T
	if x, ok := op.CheckUnwrapPtr(); ok {
		v = x
	}
	return json.Marshal(v)
}

func (op *Option[T]) UnmarshalJSON(data []byte) error {
	var v *T
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if v != nil {
		*op = Some(*v)
	} else {
		*op = None[T]()
	}
	return nil
}

func (op Option[T]) GoString() string {
	var tmp T
	t := reflect.TypeOf(&tmp).Elem()
	if op.some {
		return fmt.Sprintf("Some[%v](%# v)", t, pretty.Formatter(op.value, pretty.OptStringer(true)))
	}
	return fmt.Sprintf("None[%v]", t)
}

type Option1[T any] struct {
	Option[T]
}

func (op *Option1[T]) EncodeMV(ctx *encoding.Context) ([]byte, error) {
	var buf bytes.Buffer
	if op.IsSome() {
		buf.WriteByte(1)
		val := op.Unwrap()
		if err := encoding.Encode(&buf, &val, encoding.Ctx(ctx)); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	} else {
		return []byte{0}, nil
	}
}
