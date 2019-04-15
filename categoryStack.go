package main

type CategoryStack struct {
	Categories []*Category
}

func (cs *CategoryStack) Pop() *CategoryStack {
	cs.Categories = cs.Categories[:len(cs.Categories)-1]
	return cs
}

func (cs *CategoryStack) Push(c *Category) *CategoryStack {
	cs.Categories = append(cs.Categories, c)
	return cs
}

func (cs *CategoryStack) Tail() *Category {
	return cs.Categories[len(cs.Categories)-1]
}

func (cs *CategoryStack) Clear() {
	cs.Categories = []*Category{}
}
