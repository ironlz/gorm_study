package entity

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// User
// grom支持的模型倾向于约定，内部成员需要实现了Scanner和Valuer接口
// 默认情况下, 使用ID作为主键，使用结构体的蛇形复数作为表名，字段名的蛇形作为列名，并使用CreatedAt，UpdatedAt，DeletedAt来跟踪列的生命周期
// gorm内置了一个Model模型，定义了ID，CreatedAt，UpdatedAt，DeletedAt，可以内嵌到数据结构中，减少重复定义
// 可以使用gorm.DB的AutoMigrate接口实现自动建表，如User会被创建如下结构的表：
// CREATE TABLE `users` (
// `id` bigint unsigned NOT NULL AUTO_INCREMENT,
// `created_at` datetime(3) DEFAULT NULL,
// `updated_at` datetime(3) DEFAULT NULL,
// `deleted_at` datetime(3) DEFAULT NULL,
// `name` longtext,
// `email` longtext,
// `age` tinyint unsigned DEFAULT NULL,
// `birthday` datetime(3) DEFAULT NULL,
// `member_number` longtext,
// `actived_at` datetime(3) DEFAULT NULL,
// PRIMARY KEY (`id`),
// KEY `idx_users_deleted_at` (`deleted_at`)
// ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
type User struct {
	*gorm.Model
	//ID           uint
	//CreatedAt    time.Time
	//UpdatedAt    time.Time
	//DeletedAt    sql.NullTime
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
}

// TableName 默认情况下, gorm使用users作为表名，也可以自定义修改
func (u *User) TableName() string {
	return "users"
}

// 钩子函数
// gorm支持定义钩子函数，在User的某些特定生命周期阶段被回调
// 支持的钩子函数接口在callbacks/interfaces.go中定义
// 支持创建前后，更新前后，删除前后，查询后的相关回调
// 钩子可以被跳过，当db回话声明SkipHooks为true时，会跳过本次操作的相关回调
//

// BeforeCreate 创建前被调用
func (u *User) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("will save to db")
	return nil
}

// AdvanceUser
// gorm也支持自定义配置字段，适用于数据库表已存在的情况
// 通过对字段增加标签，来显示的声明数据库表中的字段和数据库的映射关系
// 也可以对字段的权限做限制，例如声明一个字段是只读、只写、只创建、只更新或被忽略
// 字段支持的标签列表如下
// 标签名	说明
// column	指定 db 列名
// type	列数据类型，推荐使用兼容性好的通用类型，例如：所有数据库都支持 bool、int、uint、float、string、time、bytes 并且可以和其他标签一起使用，例如：not null、size, autoIncrement… 像 varbinary(8) 这样指定数据库数据类型也是支持的。在使用指定数据库数据类型时，它需要是完整的数据库数据类型，如：MEDIUMINT UNSIGNED not NULL AUTO_INSTREMENT
// size	指定列大小，例如：size:256
// primaryKey	指定列为主键
// unique	指定列为唯一
// default	指定列的默认值
// precision	指定列的精度
// scale	指定列大小
// not null	指定列为 NOT NULL
// autoIncrement	指定列为自动增长
// embedded	嵌套字段
// embeddedPrefix	嵌入字段的列名前缀
// autoCreateTime	创建时追踪当前时间，对于 int 字段，它会追踪时间戳秒数，您可以使用 nano/milli 来追踪纳秒、毫秒时间戳，例如：autoCreateTime:nano
// autoUpdateTime	创建 / 更新时追踪当前时间，对于 int 字段，它会追踪时间戳秒数，您可以使用 nano/milli 来追踪纳秒、毫秒时间戳，例如：autoUpdateTime:milli
// index	根据参数创建索引，多个字段使用相同的名称则创建复合索引，查看 索引 获取详情
// uniqueIndex	与 index 相同，但创建的是唯一索引
// check	创建检查约束，例如 check:age > 13，查看 约束 获取详情
// <-	设置字段写入的权限， <-:create 只创建、<-:update 只更新、<-:false 无写入权限、<- 创建和更新权限
// ->	设置字段读的权限，->:false 无读权限
// -	忽略该字段，- 无读写权限

type AdvanceUser struct {
	ID           uint `gorm:"<-:create"` // 只允许读和创建
	CreatedAt    time.Time
	UpdatedAt    time.Time `gorm:"<-:update"` // 只允许读和更新
	DeletedAt    sql.NullTime
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivedAt    sql.NullTime
}
