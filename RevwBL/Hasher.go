package RevwBL

import (
	"fmt"
	"hash/fnv"
)

func HashString(input string) string {
	hash := fnv.New32a()
	hash.Write([]byte(input))
	hashU32 := hash.Sum32()
	return fmt.Sprintf("%x", hashU32)
}
