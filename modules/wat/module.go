package wat

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	"github.com/starter-go/security/modules/security"
	webapptemplate "github.com/starter-go/web-app-template"
	"github.com/starter-go/web-app-template/gen/main4wat"
	"github.com/starter-go/web-app-template/gen/test4wat"
)

// Module  ...
func Module() application.Module {
	mb := webapptemplate.NewMainModule()
	mb.Components(main4wat.ExportComponents)

	mb.Depend(libgin.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(security.Module())

	mb.Depend(mysql.Module())
	mb.Depend(sqlserver.Module())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := webapptemplate.NewTestModule()
	mb.Components(test4wat.ExportComponents)
	return mb.Create()
}
