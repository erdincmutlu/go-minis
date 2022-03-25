package main

import "strings"

type Item struct {
	name            string
	sellIn, quality int

}

// from name to type
var itemMap map[string]string

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
	
		switch itemMap[items[i].name] {
		case "Aged brie type":

		}
		
		switch items[i].name {
		case "Aged Brie", "X":
			if items[i].name == "X"{

			}

		case "Sulfuras, Hand of Ragnaros":
		case "Backstage passes to a TAFKAL80ETC concert":

		}

		qualityFactor := 1
		if strings.Contains(items[i].name, "Conjured") {
			qualityFactor = 2
		}

		if items[i].name != "Aged Brie" && items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].quality > 0 {
				if items[i].name != "Sulfuras, Hand of Ragnaros" {
					items[i].quality = items[i].quality - qualityFactor
				}
			}
		} else {
			if items[i].quality < 50 {
				if items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
					items[i].quality = items[i].quality + qualityFactor
				} else {
					items[i].quality = items[i].quality + 1
					if items[i].sellIn < 11 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
					if items[i].sellIn < 6 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
				}
			}
		}

		if items[i].name != "Sulfuras, Hand of Ragnaros" {
			items[i].sellIn = items[i].sellIn - 1
		}

		if items[i].sellIn < 0 {
			if items[i].name != "Aged Brie" {
				if items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].quality > 0 {
						if items[i].name != "Sulfuras, Hand of Ragnaros" {
							items[i].quality = items[i].quality - qualityFactor
						}
					}
				} else {
					items[i].quality = items[i].quality - items[i].quality
				}
			} else {
				if items[i].quality < 50 {
					if items[i].name == "Backstage passes to a TAFKAL80ETC concert" || items[i].name != "Sulfuras, Hand of Ragnaros" {
						items[i].quality = items[i].quality + 1
					} else {
						items[i].quality = items[i].quality + qualityFactor
					}
				}
			}
		}
	}

}
