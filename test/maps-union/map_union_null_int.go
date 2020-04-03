// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     arrays.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
	"io"
)

func writeMapUnionNullInt(r *MapUnionNullInt, w io.Writer) error {
	err := vm.WriteLong(int64(len(r.M)), w)
	if err != nil || len(r.M) == 0 {
		return err
	}
	for k, e := range r.M {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = writeUnionNullInt(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapUnionNullInt struct {
	keys   []string
	values []*UnionNullInt
	M      map[string]*UnionNullInt
}

func NewMapUnionNullInt() *MapUnionNullInt {
	return &MapUnionNullInt{
		keys:   make([]string, 0),
		values: make([]*UnionNullInt, 0),
		M:      make(map[string]*UnionNullInt),
	}
}

func (_ *MapUnionNullInt) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapUnionNullInt) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapUnionNullInt) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapUnionNullInt) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapUnionNullInt) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapUnionNullInt) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapUnionNullInt) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapUnionNullInt) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapUnionNullInt) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapUnionNullInt) SetDefault(i int)      { panic("Unsupported operation") }
func (r *MapUnionNullInt) Finalize() {
	for i := range r.keys {
		r.M[r.keys[i]] = r.values[i]
	}
	r.keys = nil
	r.values = nil
}

func (r *MapUnionNullInt) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v *UnionNullInt
	v = NewUnionNullInt()

	r.values = append(r.values, v)
	return r.values[len(r.values)-1]
}

func (_ *MapUnionNullInt) AppendArray() types.Field { panic("Unsupported operation") }