package encoding

import (
	"bytes"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	name    string
	data    []byte
	v       any
	opt     []Option
	rest    []byte
	expect  any
	encode  bool
	encoded []byte
}

type grabSlice []byte

func (g *grabSlice) DecodeMV(data []byte, ctx *Context) (rest []byte, err error) {
	*g = grabSlice(data)
	return data[len(data):], nil
}

func TestEncoding(t *testing.T) {
	type withOmit struct {
		//lint:ignore U1000 test skipping
		priv uint32
		X    uint32
		Omit uint32 `mv:"omit"`
		Y    uint64
	}

	type withDyn struct {
		X uint32
		S []byte `mv:"dyn"`
		Y uint32
	}

	type withOpt struct {
		X *uint32 `mv:"opt"`
	}

	type withOptDyn struct {
		X *uint32 `mv:"opt,dyn"`
		Y uint32
	}

	type withDynOpt struct {
		X *uint32 `mv:"dyn,opt"`
		Y uint32
	}

	type withCustomDyn struct {
		X grabSlice `mv:"dyn"`
	}

	tests := []testCase{
		{
			name:   "bool",
			data:   []byte{0xff},
			v:      new(bool),
			rest:   []byte{},
			expect: func() *bool { x := true; return &x }(),
			encode: true,
		},
		{
			name:   "uint8",
			data:   []byte{0xab},
			v:      new(uint8),
			rest:   []byte{},
			expect: func() *uint8 { x := uint8(0xab); return &x }(),
			encode: true,
		},
		{
			name:   "int8",
			data:   []byte{123},
			v:      new(int8),
			rest:   []byte{},
			expect: func() *int8 { x := int8(123); return &x }(),
			encode: true,
		},
		{
			name:   "*uint8",
			data:   []byte{0xab},
			v:      func() **uint8 { var x *uint8; return &x }(),
			rest:   []byte{},
			expect: func() **uint8 { x := uint8(0xab); p := &x; return &p }(),
			encode: true,
		},
		{
			name:   "uint16",
			data:   []byte{0xab, 0xcd},
			v:      new(uint16),
			rest:   []byte{},
			expect: func() *uint16 { x := uint16(0xabcd); return &x }(),
			encode: true,
		},
		{
			name:   "int16",
			data:   []byte{0xab, 0xcd},
			v:      new(int16),
			rest:   []byte{},
			expect: func() *int16 { x := int16(-21555); return &x }(),
			encode: true,
		},
		{
			name:   "*uint16",
			data:   []byte{0xab, 0xcd},
			v:      func() **uint16 { var x *uint16; return &x }(),
			rest:   []byte{},
			expect: func() **uint16 { x := uint16(0xabcd); p := &x; return &p }(),
			encode: true,
		},
		{
			name:   "uint32",
			data:   []byte{0x01, 0x23, 0xab, 0xcd},
			v:      new(uint32),
			rest:   []byte{},
			expect: func() *uint32 { x := uint32(0x0123abcd); return &x }(),
			encode: true,
		},
		{
			name:   "int32",
			data:   []byte{0x01, 0x23, 0xab, 0xcd},
			v:      new(int32),
			rest:   []byte{},
			expect: func() *int32 { x := int32(0x0123abcd); return &x }(),
			encode: true,
		},
		{
			name:   "*uint32",
			data:   []byte{0x01, 0x23, 0xab, 0xcd},
			v:      func() **uint32 { var x *uint32; return &x }(),
			rest:   []byte{},
			expect: func() **uint32 { x := uint32(0x0123abcd); p := &x; return &p }(),
			encode: true,
		},
		{
			name:   "uint64",
			data:   []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			v:      new(uint64),
			rest:   []byte{},
			expect: func() *uint64 { x := uint64(0x0123456789abcdef); return &x }(),
			encode: true,
		},
		{
			name:   "int64",
			data:   []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			v:      new(int64),
			rest:   []byte{},
			expect: func() *int64 { x := int64(0x0123456789abcdef); return &x }(),
			encode: true,
		},
		{
			name:   "*uint64",
			data:   []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			v:      func() **uint64 { var x *uint64; return &x }(),
			rest:   []byte{},
			expect: func() **uint64 { x := uint64(0x0123456789abcdef); p := &x; return &p }(),
			encode: true,
		},
		{
			name:   "[n]byte",
			data:   []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			v:      new([8]byte),
			rest:   []byte{},
			expect: &[8]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			encode: true,
		},
		{
			name:   "*[n]byte",
			data:   []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			v:      func() **[8]byte { var x *[8]byte; return &x }(),
			rest:   []byte{},
			expect: func() **[8]byte { x := &[8]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}; return &x }(),
			encode: true,
		},
		{
			name:   "[]byte",
			data:   []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			v:      new([]byte),
			rest:   []byte{},
			expect: &[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			encode: true,
		},
		{
			name:   "*[]byte",
			data:   []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			v:      func() **[]byte { var x *[]byte; return &x }(),
			rest:   []byte{},
			expect: func() **[]byte { x := &[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef}; return &x }(),
			encode: true,
		},
		{
			name:   "[]uint32",
			data:   []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
			v:      new([]uint32),
			rest:   []byte{},
			expect: &[]uint32{0x01234567, 0x89abcdef},
			encode: true,
		},
		{
			name:   "string",
			data:   []byte("abcd"),
			v:      new(string),
			rest:   []byte{},
			expect: func() *string { x := "abcd"; return &x }(),
			encode: true,
		},
		{
			name: "struct",
			data: []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x89, 0xab, 0xcd, 0xef},
			v:    new(withOmit),
			rest: []byte{},
			expect: &withOmit{
				X: 0x01234567,
				Y: 0x89abcdef89abcdef,
			},
			encode: true,
		},
		{
			name: "dynamic struct",
			data: []byte{
				0x01, 0x23, 0x45, 0x67, // X
				0x00, 0x00, 0x00, 0x08, // array length
				0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, // array
				0x89, 0xab, 0xcd, 0xef, // Y
			},
			v:    new(withDyn),
			rest: []byte{},
			expect: &withDyn{
				X: 0x01234567,
				S: []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77},
				Y: 0x89abcdef,
			},
			encode: true,
		},
		{
			name: "optional attribute some",
			data: []byte{
				0xff,                   // tag
				0x01, 0x23, 0x45, 0x67, // X
			},
			v:      new(withOpt),
			rest:   []byte{},
			expect: func() *withOpt { x := uint32(0x01234567); return &withOpt{X: &x} }(),
			encode: true,
		},
		{
			name: "optional attribute none",
			data: []byte{
				0x00, // tag
			},
			v:      new(withOpt),
			rest:   []byte{},
			expect: &withOpt{X: nil},
			encode: true,
		},
		{
			name: "optional and dynamic attributes",
			data: []byte{
				0xff,                   // tag
				0x00, 0x00, 0x00, 0x05, // X length
				0x01, 0x23, 0x45, 0x67, // X
				0x00,                   // extra byte
				0x89, 0xab, 0xcd, 0xef, // Y
			},
			encoded: []byte{
				0xff,                   // tag
				0x00, 0x00, 0x00, 0x04, // X length
				0x01, 0x23, 0x45, 0x67, // X
				0x89, 0xab, 0xcd, 0xef, // Y
			},
			v:      new(withOptDyn),
			rest:   []byte{},
			expect: func() *withOptDyn { x := uint32(0x01234567); return &withOptDyn{X: &x, Y: 0x89abcdef} }(),
			encode: true,
		},
		{
			name: "dynamic and optional attributes",
			data: []byte{
				0x00, 0x00, 0x00, 0x06, // X length
				0xff,                   // tag
				0x01, 0x23, 0x45, 0x67, // X
				0x00,                   // extra byte
				0x89, 0xab, 0xcd, 0xef, // Y
			},
			encoded: []byte{
				0x00, 0x00, 0x00, 0x05, // X length
				0xff,                   // tag
				0x01, 0x23, 0x45, 0x67, // X
				0x89, 0xab, 0xcd, 0xef, // Y
			},
			v:      new(withDynOpt),
			rest:   []byte{},
			expect: func() *withDynOpt { x := uint32(0x01234567); return &withDynOpt{X: &x, Y: 0x89abcdef} }(),
			encode: true,
		},
		{
			name: "custom dynamic",
			data: []byte{
				0x00, 0x00, 0x00, 0x04, // X length
				0x89, 0xab, 0xcd, 0xef, // X
				0x01, 0x02, 0x03, 0x04, // rest
			},
			v:      new(withCustomDyn),
			rest:   []byte{0x01, 0x02, 0x03, 0x04},
			expect: &withCustomDyn{X: grabSlice{0x89, 0xab, 0xcd, 0xef}},
		},
	}
	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			t.Run("Decode", func(t *testing.T) {
				gotRest, err := Decode(test.data, test.v, test.opt...)
				if test.expect == nil {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					require.Equal(t, test.expect, test.v)
					require.Equal(t, test.rest, gotRest)
				}
			})
			if test.encode {
				t.Run("Encode", func(t *testing.T) {
					var buf bytes.Buffer
					var expect []byte
					if test.encoded != nil {
						expect = test.encoded
					} else {
						expect = test.data
					}
					require.NoError(t, Encode(&buf, test.expect, test.opt...))
					require.Equal(t, expect, buf.Bytes())
				})
			}
		})
	}
}

type testEnum interface {
	testEnum()
}

type variant1 uint32

func (variant1) testEnum() {}

type variant2 struct {
	X uint32
	Y uint32
}

func (*variant2) testEnum() {}

type variant3 uint32

func (variant3) testEnum() {}

type testEnum1 interface {
	testEnum1()
}

type variant11 uint32

func (variant11) testEnum1() {}

type userType interface {
	UserType()
}

type stringType string

func (stringType) UserType() {}

func TestEnum(t *testing.T) {
	enums := NewEnumRegistry()
	enums.RegisterEnum(Variants[testEnum]{
		0: variant1(0),
		1: (*variant2)(nil),
	}, variant3(0))
	enums.RegisterEnum(Variants[testEnum1]{
		0: variant11(0),
	}, nil)

	tests := []testCase{
		{
			name: "var1",
			data: []byte{
				0x00,                   // tag
				0x01, 0x23, 0xab, 0xcd, // Var1
			},
			v:    new(testEnum),
			opt:  []Option{Enums(enums)},
			rest: []byte{},
			expect: func() *testEnum {
				var e testEnum = variant1(0x0123abcd)
				return &e
			}(),
			encode: true,
		},
		{
			name: "var2",
			data: []byte{
				0x01,                   // tag
				0x01, 0x23, 0x45, 0x67, // X
				0x89, 0xab, 0xcd, 0xef, // Y
			},
			v:    new(testEnum),
			opt:  []Option{Enums(enums)},
			rest: []byte{},
			expect: func() *testEnum {
				var e testEnum = &variant2{
					X: 0x01234567,
					Y: 0x89abcdef,
				}
				return &e
			}(),
			encode: true,
		},
		{
			name: "var3",
			data: []byte{
				0xff,                   // tag
				0x01, 0x23, 0xab, 0xcd, // Var3
			},
			v:    new(testEnum),
			opt:  []Option{Enums(enums)},
			rest: []byte{},
			expect: func() *testEnum {
				var e testEnum = variant3(0x0123abcd)
				return &e
			}(),
		},
		{
			name: "unknown tag",
			data: []byte{
				0xff, // tag
			},
			v:      new(testEnum1),
			opt:    []Option{Enums(enums)},
			rest:   nil,
			expect: nil,
		},
	}

	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			t.Run("Encode", func(t *testing.T) {
				gotRest, err := Decode(test.data, test.v, test.opt...)
				if test.expect == nil {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					require.Equal(t, test.expect, test.v)
					require.Equal(t, test.rest, gotRest)
				}
			})
			if test.encode {
				t.Run("Encode", func(t *testing.T) {
					var buf bytes.Buffer
					var expect []byte
					if test.encoded != nil {
						expect = test.encoded
					} else {
						expect = test.data
					}
					require.NoError(t, Encode(&buf, test.expect, test.opt...))
					require.Equal(t, expect, buf.Bytes())
				})
			}
		})
	}
}

func TestUser(t *testing.T) {
	types := NewTypeRegistry()
	types.RegisterType(func(data []byte, ctx *Context) (userType, []byte, error) {
		return stringType("text"), data[len(data):], nil
	})

	var out userType
	src := []byte{0x01, 0x23, 0xab, 0xcd}

	gotRest, err := Decode(src, &out, Types(types))
	require.NoError(t, err)
	require.Equal(t, []byte{}, gotRest)
	require.Equal(t, stringType("text"), out)
}

func TestEnumForEach(t *testing.T) {
	RegisterEnum(&Enum[testEnum]{
		Variants: Variants[testEnum]{
			0: variant1(0),
			1: (*variant2)(nil),
		},
		Default: variant3(0),
	})

	expect := []testEnum{
		variant1(0),
		(*variant2)(nil),
	}
	sort.Slice(expect, func(i, j int) bool { return reflect.TypeOf(expect[i]).Name() < reflect.TypeOf(expect[j]).Name() })

	result := ListVariants[testEnum]()
	sort.Slice(result, func(i, j int) bool { return reflect.TypeOf(result[i]).Name() < reflect.TypeOf(result[j]).Name() })

	require.Equal(t, expect, result)
}
