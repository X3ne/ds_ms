package validator

import (
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	validator *protovalidate.Validator
)

func init() {
	var err error
	validator, err = protovalidate.New()
	if err != nil {
		fmt.Println("failed to initialize validator:", err)
	}
}

func Validate(req protoreflect.ProtoMessage) error {
	if err := validator.Validate(req); err != nil {
		return err
	}

	return nil
}
