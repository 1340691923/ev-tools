package versions

import (
	"github.com/1340691923/eve-plugin-sdk-go/build"
)

func init() {
	Register(V0_0_1)
}

func V0_0_1() *build.Migration {
	return &build.Migration{
		ID: "0.0.1",
		MigrateSqls: []*build.ExecSql{
			{
				Sql: `create table gm_dsl_history
(
    id      INTEGER   not null primary key,
    uid     INTEGER   default 0,
    method  TEXT,
    path    TEXT,
    body    TEXT,
    created timestamp default CURRENT_TIMESTAMP not null
);
`,
			},
		},
	}
}
