package main 
import "fmt"
func main() {
	var number1 , number2 float64

	var operator string 

	fmt.Print("Enter your first number:-")
	fmt.Scanln(&number1)


	fmt.Print("Enter your second number:-")
	fmt.Scanln(&number2)

	fmt.Print("Enter Operator(- + / * ):-")
	fmt.Scanln(&operator)


switch operator {
case "+" :
	fmt.Printf("%f %s %f = %f",number1,operator,number2,number1 + number2 )

case "-" : 
fmt.Printf("%f %s %f = %f",number1 ,operator ,number2 ,number1 - number2)	

case "*" : 
fmt.Printf("%f %s %f = %f",number1 ,operator ,number2 ,number1 * number2)	

case "/" : 
if number2 == 0.0 {
	fmt.Println("Divided By Zero Situation")
}else{
	fmt.Printf("%f %s %f = %f ",number1 ,operator,number2 ,number1 / number2)
}

default : 
fmt.Println("Invalid Operator")
	
}
// if i want to display values in 3 decimal then use %.3f
}