package main

import (
	"fmt"
)

// Принцип S - Принцип единственной ответственности (Single Responsibility Principle)
// Класс должен отвечать только за одно действие. Лучшая практика - разбить разный функционал на отдельные классы.
type BookPrint struct {
	Title  string
	Author string
}

func (b BookPrint) PrintDetails() {
	fmt.Printf("Title: %s, Author: %s\n", b.Title, b.Author)
}

// Принцип О - Принцип открытости/закрытости (Open/Closed Principle)
// Класс должен быть открыт для расширения, но закрыт для изменения.
// Стоит сделать общий интерфейс с требуемым методом и классы которые будут реализовывать его в зависимости от функционала.
type Discount interface {
	ApplyDiscount(price float64) float64
}

type RegularDiscount struct{}

func (r RegularDiscount) ApplyDiscount(price float64) float64 {
	return price * 0.9
}

type HolidayDiscount struct{}

func (h HolidayDiscount) ApplyDiscount(price float64) float64 {
	return price * 0.8
}

// Принцип L - Принцип подстановки Барбары Лисков (Liskov Substitution Principle)
// Наследующий класс должен дополнять, а не замещать поведение базового класса. Т.е. при замене базового класса на наследуемый, программа должна работать так же.
type Shape interface {
	Area() float64
}

type Square struct {
	Width float64
}

func (s Square) Area() float64 {
	return s.Width * s.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Принцип I - Принцип разделения интерфейса (Interface Segregation Principle)
// Всегда нужно разделять интерфейсы так, чтобы в имплементирующем классе не было лишнего функционала.
type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

type MyPrinter struct{}

func (p MyPrinter) Print() {
	fmt.Println("Printing...")
}

type MyScanner struct{}

func (s MyScanner) Scan() {
	fmt.Println("Scanning...")
}

type MyMultiFunctionDevice struct {
	MyPrinter
	MyScanner
}

// Принцип инверсии зависимостей (Dependency Inversion Principle)
// Суть в том, чтобы основной функционал нашего проекта был зависим от абстракции, а не от конкретной реализации чего либо.
// Так как в будущем, реализация (к примеру - способ оплаты в приложении, используемая бд и тп.) может меняться.
type Storage interface {
	Save(data string)
}

type Database struct{}

func (db Database) Save(data string) {
	fmt.Println("Saving data to the database:", data)
}

type Filesystem struct{}

func (fs Filesystem) Save(data string) {
	fmt.Println("Saving data to the filesystem:", data)
}

type DataManager struct {
	storage Storage
}

func NewDataManager(storage Storage) *DataManager {
	return &DataManager{storage: storage}
}

func (dm *DataManager) SaveData(data string) {
	dm.storage.Save(data)
}

func main() {
	book := BookPrint{Title: "Clean Code", Author: "Robert C. Martin"}
	book.PrintDetails()

	discountPrice := 100.0
	regularDiscount := RegularDiscount{}
	fmt.Printf("Regular Price: $%.2f, Discounted Price: $%.2f\n", discountPrice, regularDiscount.ApplyDiscount(discountPrice))

	square := Square{Width: 5}
	fmt.Printf("Square Area: %.2f\n", square.Area())

	circle := Circle{Radius: 3}
	fmt.Printf("Circle Area: %.2f\n", circle.Area())

	multiFunctionDevice := MyMultiFunctionDevice{}
	multiFunctionDevice.Print()
	multiFunctionDevice.Scan()

	db := Database{}
	fs := Filesystem{}

	dataManagerDB := NewDataManager(db)
	dataManagerFS := NewDataManager(fs)

	dataManagerDB.SaveData("Data to save with Database storage")
	dataManagerFS.SaveData("Data to save with Filesystem storage")
}
