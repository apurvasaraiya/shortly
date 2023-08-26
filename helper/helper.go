package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"shortly/dto"
)

func WriteError(w io.Writer, message string, opts ...interface{}) error {
	msg := fmt.Sprintf(message, opts...)
	log.Default().Println(msg)
	return json.NewEncoder(w).Encode(dto.ErrorResponse{Message: msg})
}
