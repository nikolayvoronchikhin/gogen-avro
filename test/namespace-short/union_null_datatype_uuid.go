// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"errors"
	"io"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)


type UnionNullDatatypeUUIDTypeEnum int

const (
	UnionNullDatatypeUUIDTypeEnumDatatypeUUID UnionNullDatatypeUUIDTypeEnum = 1
)

type UnionNullDatatypeUUID struct { 
	DatatypeUUID DatatypeUUID

	UnionType UnionNullDatatypeUUIDTypeEnum
}

func writeUnionNullDatatypeUUID(r *UnionNullDatatypeUUID, w io.Writer) error { 
	if r == nil {
		return vm.WriteLong(int64(0), w)
	} 
	if err := vm.WriteLong(int64(r.UnionType), w); err != nil {
		return err
	}
	switch r.UnionType{ 
	case UnionNullDatatypeUUIDTypeEnumDatatypeUUID:
		return writeDatatypeUUID(r.DatatypeUUID, w)
	}
	return errors.New("invalid value for *UnionNullDatatypeUUID")
}

func (_ *UnionNullDatatypeUUID) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) SetString(v string) { panic("Unsupported operation") }

func (r *UnionNullDatatypeUUID) SetLong(v int64) { 
	r.UnionType = (UnionNullDatatypeUUIDTypeEnum)(v)
}

func (r *UnionNullDatatypeUUID) Get(i int) types.Field {
	switch (i) { 
	case 1:
		
		r.DatatypeUUID = DatatypeUUID{}
		
		
		return &r.DatatypeUUID
		
	
	}
	panic("Unknown field index")
}

func (r *UnionNullDatatypeUUID) Clear(i int) { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) ClearMap(key string) { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) Finalize()  { }
