package problem

import (
	"fmt"
	"net/http"
	"strconv"
)

/*
	Problem5 plays around with errors (2)

	URL: http://localhost:9000/problem5?name=abcde&length=4
	Expectation: everything works properly
	Reality: everything works properly (?)

	RULES:
	1. name and length must not be empty
	2. len(name) != length => WRONG
	3. len(name) == length => CORRECT
*/

// RequestForm contains data of the request
type RequestForm struct {
	Data Payload
}

// Payload is the detail of Data
type Payload struct {
	Name   string
	Length int
}

// Problem5 plays around with errors (2)
func Problem5(w http.ResponseWriter, r *http.Request) {
	response := ""
	defer func() {
		w.Write([]byte(response))
	}()

	success, err := CheckCheck(r)
	if err != nil {
		response = fmt.Sprintf("ERROR! message: %s", err.Error())
		return
	}

	if success {
		response = fmt.Sprintf("CORRECT! len(name) == length")
	} else {
		response = fmt.Sprintf("WRONG! len(name) != length")
	}
}

// CheckCheck checks everything
func CheckCheck(r *http.Request) (bool, error) {
	form := RequestForm{}

	err := form.GetData(r)
	if err != nil {
		fmt.Println(err.Error()) // log the error
		return false, err
	}

	if !form.Validate() {
		return false, err
	}

	return true, nil
}

// GetData gets data from http request and assigns the value to RequestForm
func (form *RequestForm) GetData(r *http.Request) error {
	var err error

	name := r.FormValue("name")
	lenStr := r.FormValue("length")

	// validate empty input
	if name == "" || lenStr == "" {
		return err
	}

	// convert len to int
	len, err := strconv.Atoi(lenStr)
	if err != nil {
		return err
	}

	// set data to form
	form.Data = Payload{
		Name:   name,
		Length: len,
	}

	return nil
}

// Validate validates if len(name) == length
func (form *RequestForm) Validate() bool {
	return len(form.Data.Name) == form.Data.Length
}
