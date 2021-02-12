package util

import (
	"fmt"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewObjectID() string {
	objID := primitive.NewObjectID()
	return objID.Hex()
}

func BsonToStruct(input primitive.M, outputPointer interface{}) error {
	inputByte, err := bson.Marshal(input)
	if err != nil {
		return err
	}
	err = bson.Unmarshal(inputByte, outputPointer)
	if err != nil {
		return err
	}
	return nil
}
func BsonArrToStruct(input interface{}, outputPointer interface{}) error {
	bsonRaw, err := bsonArrToRaw(input)
	if err != nil {
		return err
	}
	return bsonRaw.Unmarshal(outputPointer)
}
func bsonArrToRaw(input interface{}) (bson.RawValue, error) {
	in := struct{ Data interface{} }{Data: input}
	inByte, err := bson.Marshal(in)
	if err != nil {
		return bson.RawValue{}, err
	}
	var outRaw struct{ Data bson.RawValue }
	err = bson.Unmarshal(inByte, &outRaw)
	if err != nil {
		return bson.RawValue{}, err
	}
	return outRaw.Data, nil
}

func PrimitiveM(input interface{}, outputPointer *primitive.M) error {
	inputBtyte, err := bson.Marshal(input)
	if err != nil {
		return err
	}
	output := bson.M{}
	err = bson.Unmarshal(inputBtyte, &output)
	if err != nil {
		return err
	}
	*outputPointer = output
	return nil
}

func PrimitiveMSlice(inputs []interface{}, outputPointer *[]primitive.M) error {
	var outputList []primitive.M
	for _, input := range inputs {
		data := bson.M{}
		err := PrimitiveM(input, &data)
		if err != nil {
			return err
		}
		outputList = append(outputList, data)
	}
	*outputPointer = outputList
	return nil
}

func PrimitiveFloatTo64(input interface{}, outputPointer *int64) error {
	inputString := fmt.Sprintf("%v", input)
	inputFloat64, err := strconv.ParseFloat(inputString, 64)
	if err != nil {
		return err
	}
	outputInt64 := int64(inputFloat64)
	*outputPointer = outputInt64
	return nil
}

func PrimitiveStringToInt64(input interface{}, outputPointer *int64) error {
	inputString := fmt.Sprintf("%v", input)
	outputInt64, err := strconv.ParseInt(inputString, 10, 64)
	if err != nil {
		return err
	}
	*outputPointer = outputInt64
	return nil
}

func PrimitiveToString(input interface{}) string {
	return fmt.Sprintf("%v", input)
}

func PrimitiveTo(input interface{}) string {
	return fmt.Sprintf("%v", input)
}
