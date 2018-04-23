package initialize

import (
	"github.com/yanghai23/GoLib/atfile"
	"github.com/yanghai23/GoLib/aterr"
	"github.com/yanghai23/GoLib/atdb"
	"database/sql"
	"utils"
)

var Db *sql.DB
var BaseConfig *utils.Config

func init() {
	config, err := atdb.ReadDbConfig(atfile.GetCurrentDirectory() + "/logDBConfig.json")
	aterr.CheckErr(err)
	Db, err = atdb.InitMysql(*config)
	aterr.CheckErr(err)

	BaseConfig, err = utils.ReadBaseConfig(atfile.GetCurrentDirectory() + "/config.json")
	aterr.CheckErr(err)

}
