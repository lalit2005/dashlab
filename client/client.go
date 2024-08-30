package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func StartClient() {

	byteContent, err := os.ReadFile("input.txt")
	checkNilErr(err)

	fileContent := string(byteContent)
	questions := strings.Split(strings.TrimSpace(fileContent), "\n")

	var llmResults []string

	for _, question := range questions {
		result := fetchFromServer(question)
		llmResults = append(llmResults, result)
	}

	finalContentToBeWritten := "[\n" + strings.Join(llmResults, ",\n") + "\n]"

	file, err := os.Create("out.json")
	defer file.Close()
	checkNilErr(err)
	_, err = io.WriteString(file, string(finalContentToBeWritten))
	checkNilErr(err)
	fmt.Println("Output written to file")
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func fetchFromServer(prompt string) string {
	response, err := http.Get("http://localhost:4000/?prompt=" + url.QueryEscape(prompt))
	checkNilErr(err)
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	checkNilErr(err)

	return string(content)
}
