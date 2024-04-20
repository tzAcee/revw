package HandlerInfo

import (
	"fmt"
)

func GetEntryOfRequestbody[V any](reqBody map[string]interface{}, entry string) (V, error) {
	obj, ok := reqBody[entry]
	if !ok {
		return getZero[V](), fmt.Errorf("the request body does not contain an '%v' entry", entry)
	}

	dataEntry, ok := obj.(V)
	if !ok {
		return getZero[V](), fmt.Errorf("the '%v' entry holds an invalid type", entry)
	}

	return dataEntry, nil
}

func getZero[T any]() T {
	var result T
	return result
}
