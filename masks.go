package main

import (
	"fmt"
)

/*
1. "preview_description": "Покупка масок",
Никита собирается открывать свой пункт выдачи заказов. Для обеспечения безопасности посетителей Никита решил обеспечить
	ПВЗ одноразовыми медицинскими масками, которые будут поставляться каждый месяц в количестве n штук.

Местные поставщики продают маски по цене 0.55 рублей за одну штуку, но можно сэкономить, купив упаковку из 24 масок
	за 11 рублей или коробку из 16 упаковок за 160 рублей.
Помогите Никите, определив, сколько коробок,
	упаковок и отдельных масок он должен купить, чтобы заплатить как можно меньше денег.&nbsp;<br>\n<br>\n<em>
Примечание: Никита готов купить больше масок, чем ему нужно, если это позволит сэкономить."

Формат входных данных
В одной строке вводится одно целое число N(1<= N <= 50000)
  "is_show_input_format": true,
  "is_ready_for_solution": false,
Формат выходных данных
Требуется вывести три числа в одну строку через пробел
- количество отдельных масок, упаковок и коробок, которые надо купить.
*/

func first(n int) {
	priceOne := 0.55
	pricePackage := 11
	priceBox := 160

	countInPackage := 24
	countInBox := countInPackage * 16 //384

	cBox := n / countInBox
	if cBox > 0 {
		n %= cBox * countInBox
	}

	cPackage := n / countInPackage
	if cPackage > 0 {
		n %= cPackage * countInPackage
	}

	pPackage := cPackage * pricePackage
	if pPackage > priceBox {
		cBox++
		cPackage = 0
		n = 0
	}
	cOne := n
	pOne := float64(cOne) * priceOne
	if pOne > float64(pricePackage) {
		cPackage++
		cOne = 0
	}
	fmt.Println(cOne, cPackage, cBox)
}

func main() {
	n := 0
	fmt.Scan(&n)
	first(n)
}
