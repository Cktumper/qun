package bootstrap

type Bootstrap interface {
	Start() error
	End() error
}
