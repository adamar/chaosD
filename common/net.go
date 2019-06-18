
package common

import (
    //"bytes"
    //"encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)


func GetJson(url string) ([]byte, error){

    fmt.Println("Starting the application...")
    response, err := http.Get(url)
    if err != nil {
	    return nil, err
        //fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        //fmt.Println(string(data))
	return data, nil
    }
}
