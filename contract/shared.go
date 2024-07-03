package contract

type ShutdownReady interface {
	SetReady(ready bool) bool
}
