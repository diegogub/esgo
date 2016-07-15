package esgo

type Auther interface {
	Auth(cmd *Command) error
}
