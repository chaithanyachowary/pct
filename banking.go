package main
func main(){

	a := []customerdata{
		""chaithanya","Prathipati","3163722261",1200 ;"sai","prathipati","9491125948",2290;
		"Ram","krishna","3163723451",1670;"Karim","khan","9398022746",1256;
		"Vamsi","Yarrarapu","9949955143",9083;"sai ","Tummala","1234567880",2419;
	}



}

type customerdata struct {
	ID int
	name string
	balance string
	close bool
}

func (c *Customer) deposit(amount float64) {
	c.balance += amount
}

func (c *Customer) withdraw(amount float64) {
	c.balance -= amount
}

func (c *Customer) printBalance() {
	fmt.Println("The balance of the customer Account is ",c.balance)
}

func (c *Customer) closeAccount() {
	c.clase := False
}

