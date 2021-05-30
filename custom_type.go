package mgo

import (
	"bytes"
	"fmt"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type Decimal decimal.Decimal

func (d Decimal) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	decimalType := reflect.TypeOf(decimal.Decimal{})
	if !val.IsValid() || !val.CanSet() || val.Type() != decimalType {
		return bsoncodec.ValueDecoderError{
			Name:     "decimalDecodeValue",
			Types:    []reflect.Type{decimalType},
			Received: val,
		}
	}

	var value decimal.Decimal
	switch vr.Type() {
	case bsontype.Decimal128:
		dec, err := vr.ReadDecimal128()
		if err != nil {
			return err
		}
		value, err = decimal.NewFromString(dec.String())
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("received invalid BSON type to decode into decimal.Decimal: %s", vr.Type())
	}

	val.Set(reflect.ValueOf(value))
	return nil
}

func (d Decimal) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	decimalType := reflect.TypeOf(decimal.Decimal{})
	if !val.IsValid() || val.Type() != decimalType {
		return bsoncodec.ValueEncoderError{
			Name:     "decimalEncodeValue",
			Types:    []reflect.Type{decimalType},
			Received: val,
		}
	}

	dec := val.Interface().(decimal.Decimal)
	dec128, err := primitive.ParseDecimal128(dec.String())
	if err != nil {
		return err
	}

	return vw.WriteDecimal128(dec128)
}

type BufferType bytes.Buffer


func (d BufferType) DecodeValue(dc bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	bufferType := reflect.TypeOf(&bytes.Buffer{})
	if !val.IsValid() || !val.CanSet() || val.Type() != bufferType {
		return bsoncodec.ValueDecoderError{
			Name:     "bufferDecodeValue",
			Types:    []reflect.Type{bufferType},
			Received: val,
		}
	}

	var value *bytes.Buffer
	switch vr.Type() {
	case bsontype.Binary:
		dec,_, err := vr.ReadBinary()
		if err != nil {
			return err
		}
		value = bytes.NewBuffer(dec)
	default:
		return fmt.Errorf("received invalid BSON type to decode into decimal.Decimal: %s", vr.Type())
	}

	val.Set(reflect.ValueOf(value))
	return nil
}

func (d BufferType) EncodeValue(ec bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	bufferType := reflect.TypeOf(&bytes.Buffer{})
	if !val.IsValid() || val.Type() != bufferType {
		return bsoncodec.ValueEncoderError{
			Name:     "bufferEncodeValue",
			Types:    []reflect.Type{bufferType},
			Received: val,
		}
	}

	dec := val.Interface().(*bytes.Buffer)
	return vw.WriteBinary(dec.Bytes())
}
