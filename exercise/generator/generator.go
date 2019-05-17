package main

type antlineId int
type teratecId int

func main() {
	antlineId := genAntlineId()
	teratecId := genTeratecId()

	createAnlineAccount(antlineId())
	createTeratecAccount(teratecId())
	// createTeratecAccount(antlineId()) //Error발생
}

func createAnlineAccount(account antlineId) {
	println(account)
}

func createTeratecAccount(account teratecId) {
	println(account)
}

func genAntlineId() func() antlineId {
	var num int
	return func() antlineId {
		num++
		return antlineId(num)
	}
}

func genTeratecId() func() teratecId {
	var num int
	return func() teratecId {
		num++
		return teratecId(num)
	}
}
