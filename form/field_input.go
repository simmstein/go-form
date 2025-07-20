package form

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
	"github.com/spf13/cast"
)

// Generates an input[type=text]
func NewFieldText(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "text"))

	return f
}

// Generates an input[type=number] with default transformers
func NewFieldNumber(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "number")).
		WithBeforeBind(func(data any) (any, error) {
			return cast.ToFloat64(data), nil
		})

	return f
}

// Generates an input[type=email]
func NewFieldMail(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "email"))

	return f
}

// Generates an input[type=range]
func NewFieldRange(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "range")).
		WithBeforeBind(func(data any) (any, error) {
			return cast.ToFloat64(data), nil
		})

	return f
}

// Generates an input[type=password]
func NewFieldPassword(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "password"))

	return f
}

// Generates an input[type=hidden]
func NewFieldHidden(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "hidden"))

	return f
}

// Generates an input[type=submit]
func NewSubmit(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "submit"))

	f.Data = "Submit"

	return f
}
