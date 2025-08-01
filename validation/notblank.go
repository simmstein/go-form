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

	"github.com/spf13/cast"
)

type NotBlank struct {
	Message          string
	TypeErrorMessage string
}

func NewNotBlank() NotBlank {
	return NotBlank{
		Message:          "This value should not be blank.",
		TypeErrorMessage: "This value can not be processed.",
	}
}

func (c NotBlank) Validate(data any) []Error {
	isValid := true
	errors := []Error{}

	v := reflect.ValueOf(data)

	if data == nil || v.IsZero() {
		errors = append(errors, Error(c.Message))

		return errors
	}

	t := reflect.TypeOf(data)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.Bool:
		isValid = data == false
	case reflect.Array:
	case reflect.Slice:
		isValid = reflect.ValueOf(data).Len() > 0
	case reflect.String:
		isValid = len(data.(string)) > 0
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Int:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Int8:
	case reflect.Uint:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.Uint8:
		isValid = cast.ToFloat64(data.(string)) == float64(0)
	default:
		errors = append(errors, Error(c.TypeErrorMessage))
	}

	if !isValid {
		errors = append(errors, Error(c.Message))
	}

	return errors
}
