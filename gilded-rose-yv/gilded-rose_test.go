package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Foo(t *testing.T) {

	// tests:=[]struct{
	// 	testName string 
	// 	name string
	// 	sellin int
	// 	quality int
	// 	finalSellin int
	// 	finalQuality int
	// } {
	// 	{
	// 		testName:"Aged brie improves",
	// 		name: "Aged Brie",
	// 		sellin : 1,
	// 		quality: 1,
	// 		finalSellin: 0,
	// 		finalQuality: 2,
	// 	},
	// }

	var items = []*Item{
		&Item{"Aged Brie", 1, 1},
		&Item{"Aged Brie", 0, 0},
		&Item{"Aged Brie", -1, 2},
		&Item{"Aged Brie", 10, 49},
		&Item{"Aged Brie", 10, 50},
		&Item{"Sulfuras, Hand of Ragnaros", 5, 4},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 5, 4},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 12, 4},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 8, 4},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 0, 4},
		&Item{"Elixir of the Mongoose", 5, 7},
		&Item{"Elixir of the Mongoose", 5, 0},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Conjured Elixir", 2, 5},
		&Item{"Conjured Elixir", 2, 0},

		&Item{"Elixir", 0, 10},
		&Item{"Conjured Elixir", -1, 10},
	}

	expected := []*Item{
		&Item{"Aged Brie", 0, 2},
		&Item{"Aged Brie", -1, 2},
		&Item{"Aged Brie", -2, 4},
		&Item{"Aged Brie", 9, 50},
		&Item{"Aged Brie", 9, 50},
		&Item{"Sulfuras, Hand of Ragnaros", 5, 4},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 4, 7},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 11, 5},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 7, 6},
		&Item{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
		&Item{"Elixir of the Mongoose", 4, 6},
		&Item{"Elixir of the Mongoose", 4, 0},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Conjured Elixir", 1, 3},
		&Item{"Conjured Elixir", 1, 0},
		&Item{"Elixir", -1, 8},
		&Item{"Conjured Elixir", -2, 6},
	}

	UpdateQuality(items)

	// if items[0].name != "fixme" {
	// 	t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].name)
	// }
	require.Equal(t, expected, items)
}
