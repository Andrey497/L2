package main

import (
	"fmt"
	"time"
)

//Функция обьединение каналов
func merge(ch ...<-chan interface{}) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		for item := range ch {
			go func(val <-chan interface{}) {
				for {
					select {
					case v := <-val:
						c <- v
					}
				}
			}(ch[item])
		}
	}()
	return c
}

//=== Or channel ===
//
//Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
//Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
//однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
//В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.
//
//Определение функции:

//Пример использования функции:
var or = func(channels ...<-chan interface{}) <-chan interface{} {

	flagClose := false
	c := make(chan interface{})
	time.Now()
	for item := range channels {
		go func(channel <-chan interface{}, item int) {
		Loop:
			for {
				select {
				case val, ok := <-channel:
					if !ok {
						break Loop
					}
					if flagClose {
						c <- val
					}
				}
			}
			//попадаем суда при закрытии канала
			flagClose = true
			fmt.Println(fmt.Sprintf(`close channel %v after %v`, item, time.Now()))
		}(channels[item], item)
	}

	return c
}

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {

			defer close(c)
			time.Sleep(after)
			c <- 1
			c <- 2
		}()
		return c
	}

	start := time.Now()
	go func() {
		for val := range or(
			sig(1*time.Second),
			sig(2*time.Second),
			sig(28*time.Second),
			sig(13*time.Second),
			sig(5*time.Second),
		) {
			fmt.Println(val)
		}
	}()

	time.Sleep(100 * time.Second)

	fmt.Printf(`fone after %v`, time.Since(start))
}
