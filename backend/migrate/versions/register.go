package versions

import (
	"github.com/1340691923/eve-plugin-sdk-go/build"
)

var (
	migrationsFns []MigrationsFn
)

func Register(migrationsFn MigrationsFn) {
	migrationsFns = append(migrationsFns, migrationsFn)
}

func GetRegisterMigrationsFn() []MigrationsFn {
	return migrationsFns
}

type MigrationsFn func() *build.Migration
