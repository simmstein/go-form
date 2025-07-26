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
	"fmt"
	"time"
)

func DateBeforeMount(data any, format string) (any, error) {
	if data == nil {
		return nil, nil
	}

	switch data.(type) {
	case string:
		return data, nil
	case time.Time:
		return data.(time.Time).Format(format), nil
	case *time.Time:
		v := data.(*time.Time)
		if v != nil {
			return v.Format(format), nil
		}
	}

	return nil, nil
}

// Generates an input[type=date] with default transformers
func NewFieldDate(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "date")).
		WithBeforeMount(func(data any) (any, error) {
			return DateBeforeMount(data, "2006-01-02")
		}).
		WithBeforeBind(func(data any) (any, error) {
			return time.Parse(time.DateOnly, data.(string))
		})

	return f
}

// Generates an input[type=datetime] with default transformers
func NewFieldDatetime(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "datetime")).
		WithBeforeMount(func(data any) (any, error) {
			return DateBeforeMount(data, "2006-01-02 15:04")
		}).
		WithBeforeBind(func(data any) (any, error) {
			return time.Parse("2006-01-02T15:04", data.(string))
		})

	return f
}

// Generates an input[type=datetime-local] with default transformers
func NewFieldDatetimeLocal(name string) *Field {
	f := NewField(name, "input").
		WithOptions(
			NewOption("type", "datetime-local"),
		).
		WithBeforeMount(func(data any) (any, error) {
			return DateBeforeMount(data, "2006-01-02 15:04")
		}).
		WithBeforeBind(func(data any) (any, error) {
			a, b := time.Parse("2006-01-02T15:04", data.(string))

			return a, b
		})

	return f
}

// Generates an input[type=time] with default transformers
func NewFieldTime(name string) *Field {
	f := NewField(name, "input").
		WithOptions(NewOption("type", "time")).
		WithBeforeMount(func(data any) (any, error) {
			return DateBeforeMount(data, "15:04")
		}).
		WithBeforeBind(func(data any) (any, error) {
			if data != nil {
				v := data.(string)

				if len(v) > 0 {
					return time.Parse(time.TimeOnly, fmt.Sprintf("%s:00", v))
				}
			}

			return nil, nil
		})

	return f
}
