package main

import "fmt"
// TASK 1
// Creating student struct
type Student struct {
	name string
	age int
	grade float32
}

// function for creating the Student object
func NewStudent(name string, age int, grade float32) Student { 
	return Student{
		name: name,
		age: age,
		grade: grade,
	}
}
//function for printing student details
func ShowStudentDetails(student *Student) {
	fmt.Println(student.age, student.grade, student.name)
}


//TASK 2
// creating Book struct
type book struct {
	title string
	author string
	pages int
}

//function for creating the Book object
func NewBook(title string, author string, pages int) book {
	return book{
		title: title,
		author: author,
		pages: pages,
	}
}


func main() {
	// initializing book inctance
	var book1 book
	book1 = NewBook("Harry Potter", "J.K. Rowling", 500)
	// printing books details
	fmt.Println(book1.title, book1.author, book1.pages)
}

//TASK 3
// defining Employee struct
type Employee struct {
	Name string
	Salary float64
	Position string
}

// function for creating the Employee object
func NewEmployee(name string, salary float64, position string) Employee {
	return Employee {
		Name: name,
		Salary: salary,
		Position: position,
	}
}
//Method that increases the employee’s salary by the given percentage
func (e Employee) RaiseSalary(percent float64) {
	e.Salary = e.Salary * ( 1.0 + percent/100)
}

//TASK 4
//Defining car structure
type Car struct {
	Brand string 
	Model string
	Year int
}
// function for creating the Car object
func NewCar(brand string, model string, year int) Car {
	return Car{
		Brand: brand,
		Model: model,
		Year: year,
	}
}
// Method for defining cars Age
func (c Car) Age() int {
	return 2024 - c.Year
}

//TASK 5
//defining rectangle structure
type Rectangle struct {
	Width int
    Height int

}

// Method for creating rectangle object
func NewRectangle(Width int, Height int) Rectangle {
	return Rectangle{
		Width: Width,
		Height: Height,
	}
}
// Method for defining reacatngle area
func (r Rectangle)Area() int {
	return r.Width * r.Height
}
// Method for defining Perimeter area
func (r Rectangle)Perimeter() int {
	return (r.Width + r.Height)*2
}

//TASK 6
//Defining BankAccount structure
type BankAccount struct {
	Owner string
	Balance float64
}
//Method for definign deposit
func (b *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		b.Balance += amount
	}
}
//Method for definign Withdraw
func (b *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("amount negative")
	}
	if amount > b.Balance {
		fmt.Println("balance not enough")
	}
	b.Balance -= amount
}

//TASK 7
//Defining Person structure
type Person struct{
	Name string
	Age int
}

//Method for creating Person object
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age: age,
	}
}
//Method grow older
func (p *Person)GrowOlder() {
	p.Age += 1
}

//TASK 8 
//Defining Laptop structure
type Laptop struct {
	Brand string
	Model string
	Price float64
}

//Method for creating Laptop object
func NewLaptop(Brand string, Model string, Price float64) Laptop {
	return Laptop{
		Brand: Brand,
		Model: Model,
		Price: Price,
	}
}

func (l *Laptop)ApplyDiscount(percent float64) {
	l.Price = l.Price * ( 1.0 - percent/100)
}

//TASK 9
type Circle struct {
	Radius int
}

func NewCircle(Radius int) Circle {
	return Circle{
		Radius: Radius,
	}
}

func (c *Circle)Circumference() float64 {
	return 3.14 * float64(c.Radius) * 2
}

func (c *Circle)Area() float64 {
	return 3.14 * float64(c.Radius) * float64(c.Radius)
}

//TASK 10
type Movie struct {
	Title string 
	Director string
	Rating float64
}

func NewMovie(Title string, Director string, Rating float64) Movie {
	return Movie{
		Title: Title,
		Director: Director,
		Rating: Rating,
	}
}

func (m *Movie) UpdateRating(newRating float64) {
	m.Rating = newRating
}
