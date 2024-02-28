package main

import "fmt"

func divideByZero(a, b float64)(res float64,err error){
	defer func(){
		//Executing a call to recover inside a deferred function (but not any function called by it) stops the panicking sequence by restoring normal execution and retrieves the error value passed to the call of panic.
		r := recover();
		if r != nil {
			err = errors.New("Recovered from pani ")
		}
	}()

	fmt.Println("Recovered Successfully ")

	if (b == 0){
		panic("division by zero")
	}
	return a/b, nil
}


func main(){
	fmt.Printf("\n\nThis program divides two numbers\n")
	for {
		var a,b float64
		var c string
		fmt.Printf("\nEnter c to continue or q to quit : ")
		fmt.Scanln(&c)
		if c == "q" {
			fmt.Println("Quiting...")
			break
		}

		fmt.Printf("\nEnter first no : ")
		fmt.Scanln(&a)
		fmt.Printf("Enter second no : ")
		fmt.Scanln(&b)
		
		res, err := divideByZero(a,b)
		if err != nil {
			fmt.Printf("Error : %v\n", err.Error())
		} else{
			fmt.Printf("%f / %f = %f\n", a, b, res)
		}
	}
}
