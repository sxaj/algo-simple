package main

import "fmt"

/**

两幅牌

按顺序打出，当有一样的时候，回收手牌

最后谁手牌光，谁输


todo 什么情况下 a b 永远无法终止

*/

func main() {

	a := queue{}
	a.data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 6, 5, 4, 3, 4, 5, 6, 5, 4, 6, 7, 6, 4, 6, 7, 8, 2}
	a.head = 0
	a.tail = len(a.data)

	b := queue{}
	b.data = []int{2, 2, 2, 2}
	b.head = 0
	b.tail = len(b.data)

	fmt.Println(cardGame(a, b, 100000))
}

type queue struct {
	data []int
	head int
	tail int
}

type stack struct {
	data [10]int
	top  int
}

func cardGame(a queue, b queue, max int) bool {
	s := stack{}
	book := make(map[int]struct{}, 0)
	//循环a 和 b
	loop := 1
	for a.head < a.tail && b.head < b.tail {
		loop++
		if loop > max {
			fmt.Println("打牌次数过大 终止")
			return false
		}
		//a 出牌
		t := a.data[a.head]
		fmt.Println("a出牌", t)
		if _, ok := book[t]; !ok { //没有标记牌
			a.head++ //出队
			//加入牌堆
			s.top++
			s.data[s.top] = t
			//标记
			book[t] = struct{}{}

			fmt.Println("牌堆", s.data)
		} else { //有对应标记
			a.head++ //出队
			//赢牌 入队
			a.data = append(a.data, t)
			fmt.Println("a赢牌", t)
			a.tail++
			//遍历牌堆 赢到标记位置
			for s.data[s.top] != t {
				delete(book, s.data[s.top])
				//赢牌 入队
				fmt.Println("a赢牌", s.data[s.top])
				a.data = append(a.data, s.data[s.top])
				a.tail++
				s.data[s.top] = 0
				s.top--
			}
			//赢得一样的牌
			delete(book, s.data[s.top])
			//赢牌 入队
			fmt.Println("a赢牌", s.data[s.top])
			a.data = append(a.data, s.data[s.top])
			a.tail++
			s.data[s.top] = 0
			s.top--
		}

		//b 出牌
		t = b.data[b.head]
		fmt.Println("b出牌", t)
		if _, ok := book[t]; !ok { //没有标记牌
			b.head++ //出队
			//加入牌堆
			s.top++
			s.data[s.top] = t
			//标记
			book[t] = struct{}{}

			fmt.Println("牌堆", s.data)
		} else { //有对应标记
			b.head++ //出队
			//赢牌 入队
			b.data = append(b.data, t)
			fmt.Println("b赢牌", t)
			b.tail++
			//遍历牌堆 赢到标记位置
			for s.data[s.top] != t {
				delete(book, s.data[s.top])
				//赢牌 入队
				fmt.Println("b赢牌", s.data[s.top])
				b.data = append(b.data, s.data[s.top])
				b.tail++
				s.data[s.top] = 0
				s.top--
			}
			delete(book, s.data[s.top])
			//赢牌 入队
			fmt.Println("b赢牌", s.data[s.top])
			b.data = append(b.data, s.data[s.top])
			b.tail++
			s.data[s.top] = 0
			s.top--
		}
	}

	//某一个人牌出完

	if a.head == a.tail {
		fmt.Println("a win")
		fmt.Println("a 牌", a.data)
		fmt.Println("s 牌堆", s.data)
		return true
	} else {
		fmt.Println("b win")
		fmt.Println("b 牌", b.data)
		fmt.Println("s 牌堆", s.data)
		return false
	}
}
