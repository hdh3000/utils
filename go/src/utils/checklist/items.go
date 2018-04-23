package checklist

func NewItem(title string) Item {
	return &item{
		Name: title,
	}
}

type Item interface {
	Title() string
	NewInstance() ItemInstance
}

type ItemInstance interface {
	Title() string
	Check() bool
	IsChecked() bool
}

type item struct {
	Name string `json:"Title"`
}

type itemInstance struct {
	Parent  Item
	Checked bool
}

func (i *item) Title() string {
	return i.Name
}

func (i *item) NewInstance() ItemInstance {
	return &itemInstance{
		Parent:  i,
		Checked: false,
	}
}

func (i *itemInstance) Title() string {
	return i.Parent.Title()
}
func (i *itemInstance) Check() bool {
	i.Checked = !i.Checked
	return i.Checked
}
func (i *itemInstance) IsChecked() bool {
	return i.Checked
}
