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

// Generates an input[type=checkbox]
func NewFieldCheckbox(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "checkbox")).
		WithBeforeMount(func(data any) (any, error) {
			switch data.(type) {
			case string:
				data = data == "1"
			case bool:
				return data, nil
			}

			return cast.ToInt(data), nil
		}).
		WithBeforeBind(func(data any) (any, error) {
			switch data.(type) {
			case string:
				return data == "1", nil
			case bool:
				return data, nil
			}

			return cast.ToBool(data), nil
		})

	return f
}
