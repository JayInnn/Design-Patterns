package decorator

type InputSteam interface {
	read()
	close()
}
