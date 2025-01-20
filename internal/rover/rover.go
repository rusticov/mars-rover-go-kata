package rover

type Rover struct {
}

func (r Rover) ExecuteCommands(commands string) string {
	return "0:0:N"
}

func New() *Rover {
	return &Rover{}
}
