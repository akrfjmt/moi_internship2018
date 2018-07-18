package util

func Perror(err error) {
	if err != nil {
		panic(err)
	}
}
