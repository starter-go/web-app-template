package test4wat
import (
    pa44b9f27e "github.com/starter-go/web-app-template/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type pa44b9f27e.DemoUnit in package:github.com/starter-go/web-app-template/src/test/golang/unit
//
// id:com-a44b9f27e475440e-unit-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pa44b9f27e4_unit_DemoUnit struct {
}

func (inst* pa44b9f27e4_unit_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a44b9f27e475440e-unit-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa44b9f27e4_unit_DemoUnit) new() any {
    return &pa44b9f27e.DemoUnit{}
}

func (inst* pa44b9f27e4_unit_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa44b9f27e.DemoUnit)
	nop(ie, com)

	


    return nil
}


