package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func StartClient() {

	response, err := http.Get("http://localhost:4000")
	checkNilErr(err)
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	checkNilErr(err)

	fmt.Println(response)

	file, err := os.Create("out.json")
	defer file.Close()
	checkNilErr(err)
	_, err = io.WriteString(file, string(content))
	checkNilErr(err)
	fmt.Println("Output written to file")
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
