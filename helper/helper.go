package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"shortly/dto"
	"sort"
)

func WriteError(w io.Writer, message string, opts ...interface{}) error {
	msg := fmt.Sprintf(message, opts...)
	log.Default().Println(msg)
	return json.NewEncoder(w).Encode(dto.ErrorResponse{Message: msg})
}

func SortTopNInMap(m map[string]uint, n int) map[string]uint {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	if len(keys) > n {
		keys = keys[:n]
	}

	newMap := make(map[string]uint, n)
	for _, key := range keys {
		newMap[key] = m[key]
	}

	return newMap
}
