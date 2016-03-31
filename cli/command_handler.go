package cli

type CommandHandler struct {
	defaultHandler Handler
	handlers       map[string]Handler
}

func NewCommandHandler(h Handler) *CommandHandler {
	return &CommandHandler{
		defaultHandler: h,
		handlers:       make(map[string]Handler),
	}
}

func (c *CommandHandler) Handle(name string, handler Handler) {
	c.handlers[name] = handler
}

func (c CommandHandler) Run(args []string) int {
	var name string
	if len(args) > 0 {
		name = args[0]
	}

	handler, ok := c.handlers[name]
	if !ok {
		handler = c.defaultHandler
	}

	return handler.Run(args[1:])
}
