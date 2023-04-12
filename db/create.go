package db

import (
	"gorm_study/entity"
)

// Create
// i需要是个指针，否则无法更新内部信息
func Create(user *entity.User) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	// 根据i的内容自动建表
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
	err = db.AutoMigrate(user)
	if err != nil {
		return err
	}
	// 主键会插入到i的ID属性中
	// result.Error 返回错误信息
	// result.RowsAffected 返回插入记录的条数
	// user的BeforeCreate方法会在插入前被执行，可以通过
	// db.Session(&gorm.Session{SkipHooks: true}).Create(user)来跳过对user中定义的钩子回调
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	// 也可以只使用给定的自动去创建记录，如
	// db.Select("Name", "Age", "CreatedAt").Create(user)
	// 等价于 INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
	// 或者排除掉某些字段然后创建记录，如
	// db.Omit("Name", "Age", "CreatedAt").Create(user)
	// 等价于 INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")
	return nil
}

// BatchCreate
// 批量创建记录，创建出来的ID会被回填回对应的user中
func BatchCreate(users []*entity.User) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	// 一条语句全量插入
	result := db.Create(users)
	// 也可以使用CreateInBatches来分批插入
	// result := db.CreateInBatches(users, 100) // 每100个执行一次批量插入
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// CreateByMap
// 根据map中的字段去插入数据，根据map创建时钩子函数不会被调用，ID也不会被回填
func CreateByMap(data map[string]interface{}) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	//db.Model(&User{}).Create(map[string]interface{}{
	//	"Name": "jinzhu", "Age": 18,
	//})
	//
	// // batch insert from `[]map[string]interface{}{}`
	//db.Model(&User{}).Create([]map[string]interface{}{
	//	{"Name": "jinzhu_1", "Age": 18},
	//	{"Name": "jinzhu_2", "Age": 20},
	//})

	result := db.Model(&entity.User{}).Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
