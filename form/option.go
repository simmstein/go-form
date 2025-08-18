package form

import "strings"

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

type Option struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

func NewOption(name string, value any) *Option {
	return &Option{
		Name:  name,
		Value: value,
	}
}

func (o *Option) AsBool() bool {
	return o.Value.(bool)
}

func (o *Option) AsString() string {
	return o.Value.(string)
}

func (o *Option) AsAttrs() Attrs {
	return o.Value.(Attrs)
}

type Attrs map[string]string

func (a Attrs) Append(name, value string) {
	v, ok := a[name]

	if !ok {
		v = value
	} else {
		v = value + " " + v
	}

	a[name] = v
}

func (a Attrs) Prepend(name, value string) {
	v, ok := a[name]

	if !ok {
		v = value
	} else {
		v += " " + value
	}

	a[name] = v
}

func (a Attrs) Remove(name, value string) {
	v, ok := a[name]

	if !ok {
		v = strings.ReplaceAll(v, value, "")
	}

	a[name] = v
}
