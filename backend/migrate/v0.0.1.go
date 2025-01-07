package migrate

import (
	"github.com/1340691923/eve-plugin-sdk-go/build"
)

func V0_0_1() *build.Migration {
	return &build.Migration{
		ID: "0.0.1",
		SqliteMigrateSqls: []*build.ExecSql{
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
		MysqlMigrateSqls: []*build.ExecSql{
			{
				Sql: `CREATE TABLE gm_dsl_history
(
    id      int(11) NOT NULL AUTO_INCREMENT,
    uid     int(11) DEFAULT 0,
    method  varchar(50)   DEFAULT '',
    path    varchar(255)  DEFAULT '',
    body    text ,
    created timestamp(0) DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id) USING BTREE
) ENGINE = InnoDB  ;
`,
			},
		},
	}
}
