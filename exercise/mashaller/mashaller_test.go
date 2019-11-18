package mashaller

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_status_MarshalJSON(t *testing.T) {
	user := User{
		Name:    "nodias",
		Telecom_cd: TelecomCd{
			Telecom:     0,
			TelecomTemp: 0,
		},
	}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func Test_status_UnmarshalJSON(t *testing.T) {
	userJson := []byte(`{"name":"nodias", "telecomTemp":1}`)
	var user User
	err := json.Unmarshal(userJson, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
}
