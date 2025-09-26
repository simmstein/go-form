package form

import (
	"fmt"
	"reflect"
)

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

// Generates a sub form
func NewFieldCollection(name string) *Field {
	f := NewField(name, "collection").
		WithOptions(
			NewOption("allow_add", true),
			NewOption("allow_delete", true),
			NewOption("form", nil),
		).
		WithCollection()

	f.WithBeforeMount(func(data any) (any, error) {
		if opt := f.GetOption("form"); opt != nil {
			if src, ok := opt.Value.(*Form); ok {
				src.Name = f.GetName()
				t := reflect.TypeOf(data)

				switch t.Kind() {
				case reflect.Slice:
					slice := reflect.ValueOf(data)

					for i := 0; i < slice.Len(); i++ {
						name := fmt.Sprintf("%d", i)
						value := slice.Index(i).Interface()

						if f.HasChild(name) {
							f.GetChild(name).Mount(value)
						} else {
							form := src.Copy()
							form.Mount(value)

							field := f.Copy()
							field.Widget = "sub_form"
							field.Name = name
							field.Add(form.Fields...)
							field.
								RemoveOption("form").
								RemoveOption("label")

							for _, c := range field.Children {
								c.NamePrefix = fmt.Sprintf("[%d]", i)
							}

							f.Add(field)
						}
					}
				}
			}
		}

		return data, nil
	})

	return f
}

func NewCollection(name string) *Field {
	return NewFieldCollection(name)
}
