package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

// Only exported fields will be encoded/decoded in JSON.
// Fields must start with capital letters to be exported.

// filed with tag names
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bloB, _ := json.Marshal(true)
	fmt.Println(string(bloB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "banana", "peach"}
	slcJ, _ := json.Marshal(slcD)
	fmt.Println(string(slcJ))

	mapD := map[string]int{"apple": 2, "banana": 4}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	//You can use tags on struct field declarations
	// to customize the encoded JSON key names.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// We need to provide a variable where the
	// JSON package can put the decoded data
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// In order to use the values in the decoded map,
	// we’ll need to convert them to their appropriate type
	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	//We can also decode JSON into custom data types.
	//This has the advantages of adding additional
	//type-safety to our programs and eliminating the
	//need for type assertions when accessing the decoded data.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// We can also stream JSON encodings directly to os.Writers like
	// os.Stdout or even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
