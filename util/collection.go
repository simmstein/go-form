package util

// @license GNU AGPL version 3 or any later version
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
