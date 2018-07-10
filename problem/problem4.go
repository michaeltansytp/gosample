package problem

import (
	"fmt"
	"net/http"
)

/*
	Problem4 plays around with errors

	URL: http://localhost:9000/problem4

	Expectation: [SUMMARY] shows the error which stops the main flow
	Reality: which error is which?
*/

// Problem4 plays around with errors
func Problem4(w http.ResponseWriter, r *http.Request) {
	var err error

	defer func() {
		fmt.Println("[SUMMARY]")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("I'm happy, you're happy, we're all happy~")
		}
	}()

	if err := NoError(); err != nil {
		fmt.Println("[DEBUG] You should not see this at all")
	}

	if err := BypassableError(); err != nil {
		fmt.Println("[INFO] Meh, bypass this")
	}

	if err := FatalError(); err != nil {
		fmt.Println("[ERROR] Now this is bad...")
		return
	}

	fmt.Println("No error, yippee!")
}

// NoError returns no error (nil)
func NoError() error {
	return nil
}

// BypassableError returns error, but should be bypassed in the flow
func BypassableError() error {
	return fmt.Errorf("don't mind me, just bypass")
}

// FatalError returns error in which the main flow should stop if it happens
func FatalError() error {
	return fmt.Errorf("stop, stop, stop")
}
