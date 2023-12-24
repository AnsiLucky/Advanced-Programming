package main

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"net/http"
//)
//
//func sendRequest(url string, payload map[string]interface{}) error {
//	// map to JSON
//	jsonPayload, err := json.Marshal(payload)
//	if err != nil {
//		return err
//	}
//
//	// http post request
//	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
//	if err != nil {
//		return err
//	}
//	req.Header.Set("Content-Type", "application/json")
//
//	// client for sending request
//	client := &http.Client{}
//	_, err = client.Do(req)
//	if err != nil {
//		return err
//	}
//}
//
//func parseResponse(response string) (map[string]interface{}, error) {
//	// JSON to map
//	var data map[string]interface{}
//	err := json.Unmarshal([]byte(response), &data)
//	if err != nil {
//		return nil, err
//	}
//	return data, nil
//}
//
//func handleRequest(w http.ResponseWriter, r *http.Request) {
//	url := "https://example.com/api"
//	data := map[string]interface{}{
//		"key1": "value1",
//		"key2": 123,
//		"key3": true,
//	}
//
//	// Step 1: Send request
//	err := sendRequest(url, data)
//	if err != nil {
//		http.Error(w, "Error sending request", http.StatusInternalServerError)
//		return
//	}
//
//	// Step 2: Parse JSON response
//	parsedData, err := parseResponse()
//	if err != nil {
//		http.Error(w, "Error parsing response", http.StatusInternalServerError)
//		return
//	}
//
//	// Output the results
//	//fmt.Fprintln(w, "Response:", response)
//	fmt.Fprintln(w, "Parsed Data:", parsedData)
//}
//
//func main() {
//	http.HandleFunc("/", handleRequest)
//
//	fmt.Println("Server listening on port 8080")
//	if err := http.ListenAndServe(":8080", nil); err != nil {
//		fmt.Println("Server Error:", err)
//	}
//
//	//if err := http.ListenAndServe(":8088", nil); err != nil {
//	//	fmt.Println("Server Error:", err)
//	//}
//}
