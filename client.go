package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func main() {
	mainmenu()

}

func mainmenu() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1. Get Trending")
		fmt.Println("2. Create Post")
		fmt.Println("3. Exit")
		fmt.Println(">> ")
		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			menu1()
			break
		case "2":
			menu2()
			break
		case "3":
			os.Exit(0)
		default:
			fmt.Println("Choice not found")
		}

	}
}

func extractMessage(r *http.Response) (string, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func menu1() {
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Get("http://127.0.0.1:1234/trending")
	if err != nil {
		panic(err)
	}
	text, err := extractMessage(resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
	return
}

func menu2() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Title: ")
	sc.Scan()
	title := sc.Text()

	file, err := os.Create("New_post.txt")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	_, err = file.WriteString("Title: " + title)
	if err != nil {
		panic(err.Error())
	}

	body := &bytes.Buffer{}
	multi := multipart.NewWriter(body)
	defer multi.Close()

	part, err := multi.CreateFormFile("file", "new_post.txt")
	file.Seek(0, 0)
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:1234/post", body)
	if err != nil {
		panic(err.Error())
	}
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	data, err := extractMessage(resp)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(data)
	return
}
