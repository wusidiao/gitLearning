package main

import (
	"encoding/json"
	"fmt"
)
type Student struct {
	Grade int `json:"grade"`
	Class int `json:"class"`
	Teacher string `json:"teacher"`
}

func structToJson(student Student) []byte {
	jsonStudent,_:= json.Marshal(student)
	return jsonStudent
}
func jsonToMap (data []byte) map[string]interface{} {
	tmp := make(map[string]interface{})
	_ = json.Unmarshal(data,&tmp)
	return tmp
}
func sb() (a int) {
	a = 1
	defer func() {
		a = 2
	}()
	return a
}
func main()  {
	student := Student {Grade: 1,Class: 2,Teacher: " xhy"}
	jsonStudent := structToJson(student)
	fmt.Println(jsonStudent)
	fmt.Println(string(jsonStudent))
	mapStudent := jsonToMap(jsonStudent)

	fmt.Println(mapStudent)
	fmt.Println(sb())
}
