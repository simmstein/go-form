package validation

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
	"reflect"
	"strings"

	"github.com/spf13/cast"
)

type Length struct {
	Min              *int
	Max              *int
	MinMessage       string
	MaxMessage       string
	ExactMessage     string
	TypeErrorMessage string
}

func NewLength() Length {
	return Length{
		MinMessage:       "This value is too short (min: {{ min }}).",
		MaxMessage:       "This value is too long (max: {{ max }}).",
		ExactMessage:     "This value is not valid (expected: {{ min }}).",
		TypeErrorMessage: "This value can not be processed.",
	}
}

func (c Length) WithMin(v int) Length {
	c.Min = &v

	return c
}

func (c Length) WithMax(v int) Length {
	c.Max = &v

	return c
}

func (c Length) WithExact(v int) Length {
	c.Min = &v
	c.Max = &v

	return c
}

func (c Length) Validate(data any) []Error {
	if c.Min == nil && c.Max == nil {
		return []Error{}
	}

	errors := []Error{}

	t := reflect.TypeOf(data)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	var size *int

	switch t.Kind() {
	case reflect.Array:
	case reflect.Slice:
		s := reflect.ValueOf(data).Len()
		size = &s
	case reflect.String:
		s := len(data.(string))
		size = &s

	default:
		errors = append(errors, Error(c.TypeErrorMessage))
	}

	if size != nil {
		if c.Max != nil && c.Min != nil {
			if *c.Max == *c.Min && *size != *c.Max {
				errors = append(errors, Error(c.BuildMessage(c.ExactMessage)))
			}
		} else if c.Min != nil && *size < *c.Min {
			errors = append(errors, Error(c.BuildMessage(c.MinMessage)))
		} else if c.Max != nil && *size > *c.Max {
			errors = append(errors, Error(c.BuildMessage(c.MaxMessage)))
		}
	}

	return errors
}

func (c *Length) BuildMessage(message string) string {
	message = strings.ReplaceAll(message, "{{ min }}", cast.ToString(c.Min))
	message = strings.ReplaceAll(message, "{{ max }}", cast.ToString(c.Max))

	return message
}
