package mashaller

import (
	"encoding/json"
	"fmt"
)

func Example_User_MarshalJSON() {
	user := User{
		Name:    "nodias",
		Telecom_cd: TelecomCd{
			Telecom:        3,
			TelecomVehicle: 2,
		},
	}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	//Output:
	//
}

func Example_User_UnmarshalJSON() {
	userJson := []byte(`{"name":"nodias", "telecomCd":"KTF"}`)
	userJson2 := []byte(`{"name":"nodias", "telecomCd":{"telecom":"KTF"}}`)
	userJson3 := []byte(`{"name":"nodias", "telecomCd":{"telecomVehicle":1}}`)
	userJson4 := []byte(`{"name":"nodias", "telecomCd":{"telecom":"KTF","telecomVehicle":1}}`)
	userJson5 := []byte(`{"name":"nodias"}`)
	var user User
	err := json.Unmarshal(userJson, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

	var user2 User
	err = json.Unmarshal(userJson2, &user2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user2)

	var user3 User
	err = json.Unmarshal(userJson3, &user3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user3)

	var user4 User
	err = json.Unmarshal(userJson4, &user4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user4)


	var user5 User
	err = json.Unmarshal(userJson5, &user5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user5)
	//Output:
	//
}
