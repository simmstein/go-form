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
	"fmt"
	"net/url"
)

func MapToUrlValues(values *url.Values, prefix string, data map[string]any) {
	keyFormater := "%s"

	if prefix != "" {
		keyFormater = prefix + "[%s]"
	}

	for key, value := range data {
		keyValue := fmt.Sprintf(keyFormater, key)

		switch v := value.(type) {
		case string:
			values.Add(keyValue, v)
		case []string:
		case []int:
		case []int32:
		case []int64:
		case []any:
			for _, s := range v {
				values.Add(keyValue, fmt.Sprintf("%v", s))
			}
		case bool:
			if v {
				values.Add(keyValue, "1")
			} else {
				values.Add(keyValue, "0")
			}
		case int, int64, float64:
			values.Add(keyValue, fmt.Sprintf("%v", v))
		case map[string]any:
			MapToUrlValues(values, keyValue, v)
		default:
		}
	}
}
