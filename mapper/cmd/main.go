package main

import (
	"fmt"
	"v2/mapper"
)

func NewSkipString(skipCount int, s string) mapper.Mapper {
	return mapper.NewMapper(skipCount, s)
}

func main() {
	s := NewSkipString(3, "Aspiration.com")
	mapper.MapString(&s)
	fmt.Println(s)
}
