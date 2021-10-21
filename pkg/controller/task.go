package controller

type Task interface {
	ID() string
	Name() string
	SetID(id string)
	Equals(task Task) bool
	Run() (msg string, err error)
}

type TargetRef struct {
	RefName      string
	RefNamespace string
	RefKind      string
	RefGroup     string
	RefVersion   string
}
