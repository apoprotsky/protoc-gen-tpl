package options

import (
	"regexp"
	"strconv"

	"github.com/apoprotsky/protoc-gen-tpl/internal/registry"
)

var matchOption = regexp.MustCompile("^(\\d+): \"(.*)\"$")

// New creates Model from option string
func New(option string) *Model {
	matches := matchOption.FindStringSubmatch(option)
	if len(matches) != 3 {
		println(option, "not recognazed")
		return nil
	}
	number, err := strconv.Atoi(matches[1])
	if err != nil {
		println(option, "has not numeric id")
		return nil
	}
	registryService := registry.GetService()
	extension := registryService.GetExtensionByNumber(int32(number))
	if extension == nil {
		return nil
	}
	return &Model{
		Id:    extension.GetNumber(),
		Name:  extension.GetName(),
		Type:  extension.GetType(),
		Value: matches[2],
	}
}
