package main

type QueryCommand struct {
	Title string
	Arg   string
}

type QueryCommands []QueryCommand

func (c QueryCommands) Len() int {
	return len(c)
}

func (c QueryCommands) Index(i int) string {
	return c[i].Title
}

func (c *QueryCommands) Remove(i int) {
	*c = append([]QueryCommand(*c)[:i], []QueryCommand(*c)[i+1:]...)
}

func (c *QueryCommands) Add(cmd QueryCommand) {
	*c = append(*c, cmd)
}
