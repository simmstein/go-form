package theme

func CreateTheme(generator func() map[string]RenderFunc) map[string]RenderFunc {
	return generator()
}

func ExtendTheme(base map[string]RenderFunc, generator func() map[string]RenderFunc) map[string]RenderFunc {
	extended := CreateTheme(generator)

	for i, v := range base {
		extended[i] = v
		extended["base_"+i] = v
	}

	return extended
}
