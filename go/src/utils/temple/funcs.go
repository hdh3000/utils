package temple

var DefaultFuncs = map[string]interface{}{
	"renderChild": RenderChild,
}

// RenderChild is always available as a function in all components
func RenderChild(c *Component) (string, error) {
	return c.render()
}
