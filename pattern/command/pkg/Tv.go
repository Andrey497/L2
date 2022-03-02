package pkg

import "fmt"

type Tv struct {
	isOn          bool
	currentChanel int
}

func (t *Tv) tvOn() {
	t.isOn = true
	fmt.Println("Включение телевизора")
}

func (t *Tv) tvOff() {
	t.isOn = false
	fmt.Println("Выключение телевизора")
}
func (t *Tv) currentChannel() {
	fmt.Println(fmt.Sprintf("текущий канал:%d", t.currentChanel))
}
func (t *Tv) nextChannel() {
	t.currentChanel++
	fmt.Println(fmt.Sprintf("текущий канал:%d", t.currentChanel))
}

func (t *Tv) prevChannel() {
	t.currentChanel--
	fmt.Println(fmt.Sprintf("текущий канал:%d", t.currentChanel))
}
