package taboo

func Throw(err error) {
	panic(fromThrow(err))
}
