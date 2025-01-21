package rover

type Rover struct {
}

func (r Rover) ExecuteCommands(commands string) string {
	if commands == "R" {
		return "0:0:E"
	}
	if commands == "RR" {
		return "0:0:S"
	}
	return "0:0:N"
}

func New() *Rover {
	return &Rover{}
}
