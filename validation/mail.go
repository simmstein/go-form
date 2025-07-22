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

import "net/mail"

type Mail struct {
	Message string
}

func NewMail() Mail {
	return Mail{
		Message: "This value is not a valid email address.",
	}
}

func (c Mail) Validate(data any) []Error {
	errors := []Error{}

	if len(NewNotBlank().Validate(data)) == 0 {
		_, err := mail.ParseAddress(data.(string))

		if err != nil {
			errors = append(errors, Error(c.Message))
		}
	}

	return errors
}
