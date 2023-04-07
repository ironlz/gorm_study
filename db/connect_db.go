package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GetDB
// 获取一个数据库对象
func GetDB() (db *gorm.DB, err error) {
	// 首先构建一个类似于java jdbcurl的数据库地址，格式为：<user>:<pass>@tcp(<ip>:<port>)/<db>?<param>=<value>
	// param支持的参数由使用的数据库驱动决定，mysql的可以参考https://github.com/go-sql-driver/mysql_driver#parameters
	dsn := "root:redis@123@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDialector := mysql.Open(dsn)
	// 也可以使用mysql.Config来进行一些高级配置https://github.com/go-gorm/mysql_driver
	//cfg, _ := mysql_driver.ParseDSN(dsn)
	//config := mysql.Config{
	//	DSN:                           dsn,
	//	DSNConfig:                     cfg,
	//	SkipInitializeWithVersion:     false, // 根据当前mysql版本自动配置
	//	DefaultStringSize:             0,     // string类型的字段在数据库中的默认长度
	//	DefaultDatetimePrecision:      nil,   // 默认时间精度
	//	DisableWithReturning:          false,
	//	DisableDatetimePrecision:      false,
	//	DontSupportRenameIndex:        false,
	//	DontSupportRenameColumn:       false,
	//	DontSupportForShareClause:     false,
	//	DontSupportNullAsDefaultValue: false,
	//	DontSupportRenameColumnUnique: false,
	//}
	//mysqlDialector = mysql.New(config)
	return gorm.Open(mysqlDialector)
}
