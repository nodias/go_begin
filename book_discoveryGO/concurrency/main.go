package main

func main(){
	defer recoverPanic()
	var urls = []string{
		"http://image.com/img01.jpg",
		//"http://image.com/img02.jpg",
		//"http://image.com/img03.jpg",
	}

	for _, url := range urls{
		func(url string){
			Must(download(url))
		}(url)
	}

	////filenames, err := filepath.Glob("*.jpg")
	////if err != nil {
	////	panic(err)
	////}
	////err = writeZip("images.zip", filenames)
	////if err != nil {
	////	log.Fatal(err)
	//}
}

