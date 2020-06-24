//在工厂方法模式中，工厂父类负责定义创建产品对象的公共接口，而工厂子类则负责生成具体的产品对象，
//这样做的目的是将产品类的实例化操作延迟到工厂子类中完成，即通过工厂子类来确定究竟应该实例化哪一个具体产品类。
package factorymethod

//抽象产品
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

//抽象工厂
type OperatorFactory interface {
	Create() Operator
}
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// 具体工厂
type PlusOperatorFactor struct{}

// 具体产品
type PlusOperator struct {
	*OperatorBase
}

func (p *PlusOperatorFactor) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

func (p *PlusOperator) Result() int {
	return p.a + p.b
}
