package migrate

import (
	"ev-tools/backend/migrate/versions"
	"github.com/1340691923/eve-plugin-sdk-go/build"
)

func GetMigrates() *build.Gormigrate {

	m := new(build.Gormigrate)

	for _, fn := range versions.GetRegisterMigrationsFn() {
		m.Migrations = append(m.Migrations, fn())
	}

	return m
}
