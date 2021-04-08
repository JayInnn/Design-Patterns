package command

type Invoker struct {
	cmd Command
}

func (i *Invoker) Do() {
	i.cmd.Execute()
}
