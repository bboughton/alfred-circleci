package cli

type Handler interface {
	Run([]string) int
}
