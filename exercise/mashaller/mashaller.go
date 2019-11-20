package mashaller

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Telecom int

const (
	Default Telecom = iota
	SKT
	KTF
	LGU
)

func (t Telecom) String() string {
	switch t {
	case Default:
		return "Default"
	case SKT:
		return "SKT"
	case KTF:
		return "KTF"
	case LGU:
		return "LGU"
	default:
		return ""
	}
}

type TelecomVehicle int

const (
	VehicleSKT TelecomVehicle = iota
	VehicleKTF
	VehicleLGU
)

func (t Telecom) MarshalJSON() ([]byte, error) {
	data := t.String()
	if data == "" {
		return nil, errors.New("Telecom.MarshalJSON: unknown value")
	}
	return []byte(fmt.Sprintf(`"%s"`, t.String())), nil
}

func (t *Telecom) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "\"Default\"":
		*t = Default
	case "\"SKT\"":
		*t = SKT
	case "\"KTF\"":
		*t = KTF
	case "\"LGU\"":
		*t = LGU
	default:
		return errors.New("Telecom.UnmarshalJSON: unknown value")
	}
	return nil
}

type TelecomCd struct {
	Telecom        Telecom        `json:"telecom,omitempty"`
	TelecomVehicle TelecomVehicle `json:"telecomVehicle,omitempty"`
}

func (t TelecomCd) MarshalJSON() ([]byte, error) {
	if t.Telecom == Default {
		switch t.TelecomVehicle {
		case 0:
			t.Telecom = SKT
			break
		case 1:
			t.Telecom = KTF
			break
		case 2:
			t.Telecom = LGU
			break
		default:
			t.Telecom = Default
			break
		}
	}
	return t.Telecom.MarshalJSON()
}

func (t *TelecomCd) UnmarshalJSON(b []byte) error {
	var v interface{}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}

	switch data := v.(type) {
	case string:
		err = t.Telecom.UnmarshalJSON(b)
		if err != nil {
			return err
		}
		t.TelecomVehicle = TelecomVehicle(int(t.Telecom) - 1)
		return nil
	case map[string]interface{}:
		telecom, ok := data["telecom"]
		if ok {
			err = t.Telecom.UnmarshalJSON([]byte(fmt.Sprintf(`"%s"`, telecom.(string))))
			if err != nil {
				return err
			}
			t.TelecomVehicle = TelecomVehicle(int(t.Telecom) - 1)
		} else {
			t.TelecomVehicle = TelecomVehicle(data["telecomVehicle"].(float64))
			t.Telecom = Telecom(t.TelecomVehicle +1)
		}
		return nil
	default:
		return nil
	}
}

type User struct {
	Name       string    `json:"name"`
	Telecom_cd TelecomCd `json:"telecomCd,omitempty"`
}
