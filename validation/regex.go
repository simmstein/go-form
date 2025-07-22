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
	"regexp"
)

type Regex struct {
	Message          string
	TypeErrorMessage string
	Match            bool
	Expression       string
}

func NewRegex(expr string) Regex {
	return Regex{
		Message:          "This value is not valid.",
		TypeErrorMessage: "This value can not be processed.",
		Match:            true,
		Expression:       expr,
	}
}

func (c Regex) MustMatch() Regex {
	c.Match = true

	return c
}

func (c Regex) MustNotMatch() Regex {
	c.Match = false

	return c
}

func (c Regex) Validate(data any) []Error {
	errors := []Error{}

	if len(NewNotBlank().Validate(data)) == 0 {
		t := reflect.TypeOf(data)

		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		switch t.Kind() {
		case reflect.String:
			matched, _ := regexp.MatchString(c.Expression, data.(string))

			if !matched && c.Match || matched && !c.Match {
				errors = append(errors, Error(c.Message))
			}

		default:
			errors = append(errors, Error(c.TypeErrorMessage))
		}
	}

	return errors
}
