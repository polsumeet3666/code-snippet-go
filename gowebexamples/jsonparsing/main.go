package main

import (
	"encoding/json"
	"fmt"
)

/*node
> var user = {fname:'sumeet',lname:'pol' }
undefined
> JSON.stringify(user)
'{"fname":"sumeet","lname":"pol"}'
>
*/

// User struct type
type User struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

func main() {

	jsonString := `{"fname":"sumeet","lname":"pol"}`
	var user1 User
	json.Unmarshal([]byte(jsonString), &user1)
	fmt.Printf("\n %s - %s\n", user1.Fname, user1.Lname)

	userbytes, err := json.Marshal(user1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(userbytes))

}
