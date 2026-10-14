package main

import (
	"fmt"

	"github.com/omo-ri/x5-go/lessons/lesson03/order"
)

func main() {
	o, err := order.NewOrder(1)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	items := []order.Item{
		{Name: "Хлеб", PriceKopecks: 5000},
		{Name: "Молоко", PriceKopecks: 9000},
		{Name: "", PriceKopecks: 100},
		{Name: "Сыр", PriceKopecks: 30000},
	}
	for _, item := range items {
		if err := order.AddItem(o, item); err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Printf("added %s: len=%d cap=%d\n", item.Name, len(o.Items), cap(o.Items))
	}

	fmt.Println("status:", order.StatusNames[o.Status])
	fmt.Println("total:", order.Total(*o))
}
