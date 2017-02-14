package controller

type Scope string

const (
	GlobalScope  Scope = "global"
	LimitedScope Scope = "limited"
)

type Config struct {
	ListenAddr string
	Scope      Scope
}
