package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm with 'Enter'.")

	var prices []string
	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (cmd CMDManager) WriteResult(job interface{}) error {
	fmt.Println(job)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
