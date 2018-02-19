package helpers

import (
	"errors"
	"strconv"
)

var (
	errInvalidArgumentType = errors.New("Invalid argument type")
)

func Atoi(fromI interface{}) (interface{}, error) {
	switch from := fromI.(type) {
	case string:
		return strconv.Atoi(from)

	case []string:
		result := []int{}
		for _, itemStr := range from {
			item, err := strconv.Atoi(itemStr)
			if err != nil {
				return nil, err
			}
			result = append(result, item)
		}
		return result, nil
	}

	return nil, errInvalidArgumentType
}
