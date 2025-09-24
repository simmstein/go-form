package util

import (
	"regexp"
	"strings"

	"github.com/spf13/cast"
)

type CollectionValue struct {
	Name     string
	Value    string
	Children map[string]*CollectionValue
}

type Collection struct {
	Children map[int]*CollectionValue
}

func NewCollection() *Collection {
	return &Collection{
		Children: make(map[int]*CollectionValue),
	}
}

func NewCollectionValue(name string) *CollectionValue {
	return &CollectionValue{
		Name:     name,
		Children: make(map[string]*CollectionValue),
	}
}

func (c *Collection) Add(indexes []string, value string) {
	firstIndex := cast.ToInt(indexes[0])
	size := len(indexes)
	child := c.Children[firstIndex]

	if child == nil {
		child = NewCollectionValue(indexes[0])
		c.Children[firstIndex] = child
	}

	child.Add(indexes[1:size], value, nil)
}

func (c *Collection) Slice() []any {
	var result []any

	for _, child := range c.Children {
		result = append(result, child.Map())
	}

	return result
}

func (c *CollectionValue) Map() any {
	if len(c.Children) == 0 {
		return c.Value
	}

	results := make(map[string]any)

	for _, child := range c.Children {
		results[child.Name] = child.Map()
	}

	return results
}

func (c *CollectionValue) Add(indexes []string, value string, lastChild *CollectionValue) {
	size := len(indexes)

	if size > 0 {
		firstIndex := indexes[0]
		child := c.Children[firstIndex]

		child = NewCollectionValue(indexes[0])
		c.Children[firstIndex] = child

		child.Add(indexes[1:size], value, child)
	} else {
		lastChild.Value = value
	}
}

func ExtractDataIndexes(value string) []string {
	re := regexp.MustCompile(`\[[^\]]+\]`)
	items := re.FindAll([]byte(value), -1)
	var results []string

	for _, i := range items {
		results = append(results, strings.Trim(string(i), "[]"))
	}

	return results
}
