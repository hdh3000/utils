package checklist

// List is a template for a ListInstance. It provides a factory
// for creating ListInstances
type List interface {
	// Title returns the title of the list
	Title() string
	// NewInstance creates a new instantiation of the list
	NewInstance(id string) ListInstance
	// Elements gets all the items on the list
	Elements() []Item
	// Produce a new list from the combination of two lists
	// It will not add duplicate items (based on title)
	Combine(List) List
	// PushItem adds a new item to the end of the list
	Push(Item) bool
	// Drop Item finds the element with the corresponding title, and drops it from the list.
	Drop(Item) bool
}

func NewList(title string, items ...Item) List {
	return &list{
		Name:  title,
		Items: items,
	}
}

type list struct {
	Name  string `json:"Title"`
	Items []Item
}

func (l *list) Title() string {
	return l.Name
}

func (l *list) NewInstance(id string) ListInstance {
	items := make([]ItemInstance, len(l.Elements()))
	for i, item := range l.Elements() {
		items[i] = item.NewInstance()
	}

	return &listInstance{
		Id:     id,
		Parent: l,
		Items:  items,
	}
}

func (l *list) Elements() []Item {
	return l.Items
}

func (l *list) Combine(il List) List {
	return NewList(l.Title(), append(l.Items, il.Elements()...)...)
}

func (l *list) Push(item Item) bool {
	for _, li := range l.Items {
		if li.Title() == item.Title() {
			return false
		}
	}

	l.Items = append(l.Items, item)
	return true
}

func (l *list) Drop(item Item) bool {
	for i, li := range l.Items {
		if li.Title() == item.Title() {
			l.Items = append(l.Items[i:], l.Items[:i+1]...)
			return true
		}
	}

	return false
}

// ListInstance is an instance produced by a List. It
// can be checked, unchecked and stored.
type ListInstance interface {
	Title() string
	ID() string
	All() []ItemInstance
	Checked() []ItemInstance
	UnChecked() []ItemInstance
}

type listInstance struct {
	Parent List
	Id     string
	Items  []ItemInstance
}

func (l *listInstance) ID() string {
	return l.Id
}

func (l *listInstance) Title() string {
	return l.Parent.Title()
}

// All() returns all items on the list, sorted by add order
func (l *listInstance) All() []ItemInstance {
	return l.Items
}

// Checked returns all the items checked off the list
func (l *listInstance) Checked() []ItemInstance {
	var out []ItemInstance
	for _, item := range l.Items {
		if item.IsChecked() {
			out = append(out, item)
		}
	}
	return out
}

// Unchecked returns all the items that are yet to be checked off.
func (l *listInstance) UnChecked() []ItemInstance {
	var out []ItemInstance
	for _, item := range l.Items {
		if !item.IsChecked() {
			out = append(out, item)
		}
	}
	return out
}
