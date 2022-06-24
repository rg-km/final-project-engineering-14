package helper

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIfErrorWithMessage(message string, err error) {
	if err != nil {
		panic(message + " " + err.Error())
	}
}
