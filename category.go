package main

import "time"

type Category struct {
	CategoryName  string
	SubCategories CategoryList
	Projects      []Project
}

func (cat *Category) AddSubCategory(outerCat *Category) {
	cat.SubCategories.Add(outerCat)
}

type Project struct {
	StarsAmount int
	Description string
	LastCommit  time.Time
}
