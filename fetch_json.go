package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

type Result struct {
	TotalPages int `json:"total_pages"`
}


func findAndFillIntValues(v string, out *[]interface{}) {
	pattern := "[0-9]+"
	re := regexp.MustCompile(pattern)
	match := re.FindString(v)

	if match != "" {
		intValue, err := strconv.Atoi(match)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			*out = append(*out, intValue)
		}
	}
}

func main() {
	url := "https://jsonmock.hackerrank.com/api/weather/search?"
	var name string
	fmt.Printf("\nEnter name : ")
	fmt.Scanln(&name)

	url = url + "name=" + name

	// Make a GET request to the endpoint
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}

	result := Result{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	var finalOut []interface{}
	for i := 1; i <= result.TotalPages; i++ {
		new_url := url + fmt.Sprintf("&page=%v", i)
		fmt.Print("GET ")
		fmt.Println(new_url)

		// Make a GET request to the endpoint
		res, err := http.Get(new_url)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer res.Body.Close()
		// Read the response body
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		//json tags are not needed here because unmarshal will automatically identify Data field of struct as data key of json and so on
		var data struct {
			Data []struct {
				Name    string
				Weather string
				Status  []string
			}
		}

		jsonErr := json.Unmarshal(body, &data)
		if jsonErr != nil {
			fmt.Println(jsonErr)
		}

		for _, v := range data.Data {
			var out []interface{}
			out = append(out, v.Name)
			findAndFillIntValues(v.Weather,&out)
			for _, val := range v.Status {
				findAndFillIntValues(val,&out)
			}
			finalOut = append(finalOut, out)
		}
	}
	fmt.Printf("\n\nOutput : ")
	fmt.Println(finalOut)
}

