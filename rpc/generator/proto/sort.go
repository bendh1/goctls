package proto

import (
	"github.com/bendh1/goctls/rpc/parser"
	"github.com/duke-git/lancet/v2/slice"
)

func SortImport(data map[string]parser.Import) (result []string) {
	result = []string{}
	for k, _ := range data {
		result = append(result, k)
	}

	slice.Sort(result)

	return result
}

func SortEnum(data map[string]parser.Enum) (result []string) {
	result = []string{}
	for k, _ := range data {
		result = append(result, k)
	}

	slice.Sort(result)

	return result
}

func SortMessage(data map[string]parser.Message) (result []string) {
	result = []string{}
	for k, _ := range data {
		result = append(result, k)
	}

	slice.Sort(result)

	return result
}

func SortService(data map[string]parser.Service) (result []string) {
	result = []string{}
	for k, _ := range data {
		result = append(result, k)
	}

	slice.Sort(result)

	return result
}
