package main

// NewCatalog create store
func NewCatalog() Catalog {
	items := map[string]Item{
		"mouse":    Item{ID: "mouse", Name: "マウス", Price: 18000},
		"keyboard": Item{ID: "keyboard", Name: "キーボード", Price: 22000},
	}
	return Catalog{ItemMap: items}
}

// Item provides item name and ID
type Item struct {
	ID    string
	Name  string
	Price int
}

// Catalog has shopping items
type Catalog struct {
	ItemMap map[string]Item
}
