package adapter

// 适配的目标接口
type Target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest() string
}

func NewAdaptee() Adaptee {
	return &adapteeImp{}
}

type adapteeImp struct{}

func (a *adapteeImp) SpecificRequest() string {
	return "specific request"
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{adaptee}
}

type adapter struct {
	Adaptee
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}
