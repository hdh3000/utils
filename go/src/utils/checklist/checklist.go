package checklist

type ListCollection struct {
	Name  string
	Lists []*List
}

type List struct {
	Name  string
	Items []*ListItem
}

type ListItem struct {
	Name    string
	Checked bool
}
