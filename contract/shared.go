package contract

//go:generate mockery --name ShutdownReady
type ShutdownReady interface {
	SetReady(ready bool) bool
}
