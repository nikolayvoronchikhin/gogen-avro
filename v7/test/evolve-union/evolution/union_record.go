// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type UnionRecord struct {
	A *UnionNullString `json:"a"`

	Name string `json:"name"`
}

const UnionRecordAvroCRC64Fingerprint = "\xf1\xaa\xd1\x1b\x17fj\xae"

func NewUnionRecord() *UnionRecord {
	return &UnionRecord{}
}

func DeserializeUnionRecord(r io.Reader) (*UnionRecord, error) {
	t := NewUnionRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeUnionRecordFromSchema(r io.Reader, schema string) (*UnionRecord, error) {
	t := NewUnionRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeUnionRecord(r *UnionRecord, w io.Writer) error {
	var err error
	err = writeUnionNullString(r.A, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.Name, w)
	if err != nil {
		return err
	}
	return err
}

func (r *UnionRecord) Serialize(w io.Writer) error {
	return writeUnionRecord(r, w)
}

func (r *UnionRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"a\",\"type\":[\"null\",\"string\"]},{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"UnionRecord\",\"type\":\"record\"}"
}

func (r *UnionRecord) SchemaName() string {
	return "UnionRecord"
}

func (_ *UnionRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *UnionRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *UnionRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *UnionRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *UnionRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *UnionRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *UnionRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ *UnionRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *UnionRecord) Get(i int) types.Field {
	switch i {
	case 0:
		r.A = NewUnionNullString()

		return r.A
	case 1:
		return &types.String{Target: &r.Name}
	}
	panic("Unknown field index")
}

func (r *UnionRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *UnionRecord) NullField(i int) {
	switch i {
	case 0:
		r.A = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ *UnionRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionRecord) Finalize()                        {}

func (_ *UnionRecord) AvroCRC64Fingerprint() []byte {
	return []byte(UnionRecordAvroCRC64Fingerprint)
}
