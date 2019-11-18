package mashaller

import (
	"errors"
	"fmt"
)

type telecom int

func (t telecom) String() string {
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

const (
	Default telecom = iota
	SKT
	KTF
	LGU
)

func (t telecom) MarshalJSON() ([]byte, error) {
	data := t.String()
	if data == "" {
		return nil, errors.New("telecom.MarshalJSON: unknown value")
	}
	return []byte(fmt.Sprintf(`"%s"`, t.String())), nil
}

func (t *telecom) UnmarshalJSON(data []byte) error {
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
		return errors.New("telecom.UnmarshalJSON: unknown value")
	}
	return nil
}

type TelecomCd struct {
	Telecom     telecom `json:"telecom,omitempty"`
	TelecomTemp int     `json:"telecomTemp,omitempty"`
}

func (t TelecomCd) MarshalJSON() ([]byte, error) {
	if t.Telecom == Default {
		switch t.TelecomTemp {
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

func (t TelecomCd) UnmarshalJSON([]byte) error {
return nil
}

type User struct {
	Name       string    `json:"name"`
	Telecom_cd TelecomCd `json:"telecomCd,omitempty"`
}
