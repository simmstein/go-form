package form

import (
	"reflect"

	"github.com/spf13/cast"
)

type Choice struct {
	Value string
	Label string
	Data  any
}

func (c Choice) Match(value string) bool {
	return c.Value == value
}

type Choices struct {
	Data         any
	ValueBuilder func(key int, item any) string
	LabelBuilder func(key int, item any) string
}

func (c *Choices) Match(f *Field, value string) bool {
	if f.IsSlice {
		v := reflect.ValueOf(f.Data)

		for key, _ := range c.GetChoices() {
			for i := 0; i < v.Len(); i++ {
				item := v.Index(i).Interface()

				switch item.(type) {
				case string:
					if item == value {
						return true
					}
				default:
					if c.ValueBuilder(key, item) == value {
						return true
					}
				}
			}
		}

		return false
	}

	return f.Data == value
}

func (c *Choices) WithValueBuilder(builder func(key int, item any) string) *Choices {
	c.ValueBuilder = builder

	return c
}

func (c *Choices) WithLabelBuilder(builder func(key int, item any) string) *Choices {
	c.LabelBuilder = builder

	return c
}

func (c *Choices) GetChoices() []Choice {
	choices := []Choice{}

	v := reflect.ValueOf(c.Data)

	switch v.Kind() {
	case reflect.Slice, reflect.Array, reflect.String, reflect.Map:
		for i := 0; i < v.Len(); i++ {
			choices = append(choices, Choice{
				Value: c.ValueBuilder(i, v.Index(i).Interface()),
				Label: c.LabelBuilder(i, v.Index(i).Interface()),
				Data:  v.Index(i).Interface(),
			})
		}
	}

	return choices
}

func NewChoices(items any) *Choices {
	builder := func(key int, item any) string {
		return cast.ToString(key)
	}

	choices := Choices{
		ValueBuilder: builder,
		LabelBuilder: builder,
		Data:         items,
	}

	return &choices
}

func NewFieldChoice(name string) *Field {
	f := NewField(name, "choice").
		WithOptions(
			NewOption("choices", &Choices{}),
			NewOption("expanded", false),
			NewOption("multiple", false),
			NewOption("empty_choice_label", "None"),
		)

	f.WithBeforeBind(func(data any) (any, error) {
		choices := f.GetOption("choices").Value.(*Choices)

		switch data.(type) {
		case string:
			v := data.(string)
			for _, c := range choices.GetChoices() {
				if c.Match(v) {
					return c.Data, nil
				}
			}
		case []string:
			v := reflect.ValueOf(data)
			var res []interface{}

			for _, choice := range choices.GetChoices() {
				for i := 0; i < v.Len(); i++ {
					item := v.Index(i).Interface().(string)
					if choice.Match(item) {
						res = append(res, choice.Data)
					}
				}
			}

			return res, nil
		}

		return data, nil
	})

	return f
}
