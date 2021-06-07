package test

import (
	"fmt"
	"github.com/alvaro259818/golang-restclient/rest"
	"github.com/alvaro259818/golang-testing/src/api/app"
	"os"
	"testing"
)

func TestMain(m *testing.M)  {
	rest.StartMockupServer()
	fmt.Println("About to start the application...")
	go app.StartApp()
	fmt.Println("Application started, about to start test cases...")
	os.Exit(m.Run())
}
