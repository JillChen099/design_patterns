package strategy

import (
	"testing"
	"fmt"
)

func TestStrategy(t *testing.T) {
	mt := new(MovieTicker)
	mt.setPrice(60.0)
	fmt.Println("电影票原始价为：",mt.price)

	dc := new(VIPDiscount)
	mt.setDiscount(dc)

	currentPrice := mt.getPrice()
	fmt.Println("折扣价为：",currentPrice)

}
