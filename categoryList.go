package main

type CategoryList struct {
	Categories []*Category
}

func (cl *CategoryList) Add(c *Category) {
	cl.Categories = append(cl.Categories, c)
}

func (cl *CategoryList) IsEmpty() bool {
	return len(cl.Categories) == 0
}

type WalkFunc func(c *Category)

func (cl *CategoryList) Walk(f WalkFunc) {
	var walk func(cl CategoryList, f WalkFunc)
	walk = func(cl CategoryList, f WalkFunc) {
		for _, c := range cl.Categories {
			if c.SubCategories.IsEmpty() {
				f(c)
			} else {
				f(c)
				walk(c.SubCategories, f)
			}
		}
	}
	walk(*cl, f)
}
