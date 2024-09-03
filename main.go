package main

import (
	"fmt"
	"io"
	"net/http"
)

type ApiResponse struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	fmt.Println("Welcome to the rest client microservice communication.........")

	request, err := http.NewRequest(http.MethodGet, "http://localhost:1008/get", nil)
	if err != nil {
		fmt.Println("Error while creating the http request", err)
		return
	}

	headers := http.Header{}
	headers.Add("Accept", "application/json")
	request.Header = headers

	client := &http.Client{
		//Transport:     nil,
		//CheckRedirect: nil,
		//Jar:           nil,
		//Timeout:       0,
	}

	// Send the request via the client
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Error while request via client %v\n ", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	//read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error while reading the body %v\n", err)
		return
	}

	//unmarshal into the struct
	//var user ApiResponse
	//err = json.Unmarshal(body, &user)
	//if err != nil {
	//	fmt.Printf("Error while unmarshaling the body %v\n", err)
	//	return
	//}

	fmt.Printf("The response of the api call %v\n", string(body))
	//fmt.Printf("ID : %v\n", user.ID)
	//fmt.Printf("Title :  %v\n", user.Title)
	//fmt.Printf("Body : %v\n", user.Body)
}
