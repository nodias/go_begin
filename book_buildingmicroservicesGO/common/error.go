package common

func MustByte(b []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return b
}
