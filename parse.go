package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Customers struct {
	Customers []Customer `json:"users"`
}

type Customer struct {
	Name     string   `json:"name"`
	Company  string   `json:"company"`
	Services Services `json:"services"`
	SSO      string   `json:"login"`
}

type Services struct {
	Google   string `json:"Google"`
	Facebook string `json:"Facebook"`
	Amazon   string `json:"Amazon"`
}

func main() {
	File, err := os.Open("data1.json") // using os package to open json file
	if err != nil {                    // checks if there is an error when opening json file
		fmt.Println(err)
	}

	fmt.Println("data1.json has been opened without any issues")

	defer File.Close()

	// parsing contents of json file as byte array
	byteVal, _ := ioutil.ReadAll(File)

	var customers Customers // create instance using Customers

	json.Unmarshal(byteVal, &customers)

	//going through array to output data to screen

	for i := 0; i < len(customers.Customers); i++ {
		fmt.Println("Customer Name: " + customers.Customers[i].Name)
		fmt.Println("Customer Company: " + customers.Customers[i].Company)
		fmt.Println("Customer SSO: " + customers.Customers[i].SSO)
		fmt.Println("Google link: " + customers.Customers[i].Services.Google)
		fmt.Println("Facebook link: " + customers.Customers[i].Services.Facebook)
		fmt.Println("Amazon link: " + customers.Customers[i].Services.Amazon)
	}
}
