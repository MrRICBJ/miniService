package v1

import (
	"fmt"
	"strconv"
)

func valid(id []string) ([]int, error) {
	if id == nil {
		return nil, fmt.Errorf("bad request")
	}
	res := make([]int, len(id))
	var err error
	for i, v := range id {
		res[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
