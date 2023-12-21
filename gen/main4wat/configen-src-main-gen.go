package main4wat
import (
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
    p7b81a070f "github.com/starter-go/web-app-template/app/data/dao"
    p61f784186 "github.com/starter-go/web-app-template/app/data/database"
    pf0504f32e "github.com/starter-go/web-app-template/app/data/dxo"
    p38900dfec "github.com/starter-go/web-app-template/app/implements/example"
    pa05b85197 "github.com/starter-go/web-app-template/app/web/controllers/apiv1"
     "github.com/starter-go/application"
)

// type p61f784186.ThisGroup in package:github.com/starter-go/web-app-template/app/data/database
//
// id:com-61f784186728c903-database-ThisGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:alias-f0504f32ef23a57c77c365ee536fcbe1-DatabaseAgent
// scope:singleton
//
type p61f7841867_database_ThisGroup struct {
}

func (inst* p61f7841867_database_ThisGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-61f784186728c903-database-ThisGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = "alias-f0504f32ef23a57c77c365ee536fcbe1-DatabaseAgent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p61f7841867_database_ThisGroup) new() any {
    return &p61f784186.ThisGroup{}
}

func (inst* p61f7841867_database_ThisGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p61f784186.ThisGroup)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Alias = inst.getAlias(ie)
    com.URI = inst.getURI(ie)
    com.Prefix = inst.getPrefix(ie)
    com.Source = inst.getSource(ie)
    com.SourceManager = inst.getSourceManager(ie)


    return nil
}


func (inst*p61f7841867_database_ThisGroup) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.default.enabled}")
}


func (inst*p61f7841867_database_ThisGroup) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.alias}")
}


func (inst*p61f7841867_database_ThisGroup) getURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.uri}")
}


func (inst*p61f7841867_database_ThisGroup) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.table-name-prefix}")
}


func (inst*p61f7841867_database_ThisGroup) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.datasource}")
}


func (inst*p61f7841867_database_ThisGroup) getSourceManager(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}



// type p38900dfec.DaoImpl in package:github.com/starter-go/web-app-template/app/implements/example
//
// id:com-38900dfec1721eba-example-DaoImpl
// class:
// alias:alias-7b81a070f2b3b382b12754d3526915a5-ExampleDAO
// scope:singleton
//
type p38900dfec1_example_DaoImpl struct {
}

func (inst* p38900dfec1_example_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-38900dfec1721eba-example-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-7b81a070f2b3b382b12754d3526915a5-ExampleDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p38900dfec1_example_DaoImpl) new() any {
    return &p38900dfec.DaoImpl{}
}

func (inst* p38900dfec1_example_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p38900dfec.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)


    return nil
}


func (inst*p38900dfec1_example_DaoImpl) getAgent(ie application.InjectionExt)pf0504f32e.DatabaseAgent{
    return ie.GetComponent("#alias-f0504f32ef23a57c77c365ee536fcbe1-DatabaseAgent").(pf0504f32e.DatabaseAgent)
}



// type pa05b85197.ExampleController in package:github.com/starter-go/web-app-template/app/web/controllers/apiv1
//
// id:com-a05b851976363380-apiv1-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pa05b851976_apiv1_ExampleController struct {
}

func (inst* pa05b851976_apiv1_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a05b851976363380-apiv1-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa05b851976_apiv1_ExampleController) new() any {
    return &pa05b85197.ExampleController{}
}

func (inst* pa05b851976_apiv1_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa05b85197.ExampleController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*pa05b851976_apiv1_ExampleController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*pa05b851976_apiv1_ExampleController) getDao(ie application.InjectionExt)p7b81a070f.ExampleDAO{
    return ie.GetComponent("#alias-7b81a070f2b3b382b12754d3526915a5-ExampleDAO").(p7b81a070f.ExampleDAO)
}


