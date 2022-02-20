package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	//result, _ := UnPackageString("a4bc2d5e")
	//if result == "aaaabccddddde" {
	//	fmt.Println(1)
	//	fmt.Println(result)
	//}
	//result, _ = UnPackageString("abcd")
	//if result == "abcd" {
	//	fmt.Println(2)
	//	fmt.Println(result)
	//}
	//result, err := UnPackageString("45")
	//if err != nil {
	//	fmt.Println(3)
	//	fmt.Println(err)
	//	fmt.Println(result)
	//}
	//result, err = UnPackageString(``)
	//if result == "" {
	//	fmt.Println(4)
	//	fmt.Println(result)
	//}
	//result, err = UnPackageString(`qwe\4\5`)
	//if result == "qwe45" {
	//	fmt.Println(5)
	//	fmt.Println(result)
	//}
	//result, err = UnPackageString(`qwe\45`)
	//if result == `qwe44444` {
	//	fmt.Println(6)
	//	fmt.Println(result)
	//}
	//result, err = UnPackageString(`qwe\\5`)
	//if result == `qwe\\\\\` {
	//	fmt.Println(7)
	//	fmt.Println(result)
	//}
}
