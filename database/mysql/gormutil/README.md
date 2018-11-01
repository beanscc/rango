### gormutil

`gormutil` 包封装了 gorm 的一些常用

特点：
- 封装了 gorm.DB 的快速创建，避免重复代码
- 使用 logrus json 结构化格式的日志（可根据使用者个人习惯，替换成其他日志库），代替 gorm 原生的文本日志
- 支持从上下文中获取 logger 对象，以便输出 sql 日志时，携带上下文信息。（通过 SlaveLoggerFromContext/MasterLoggerFromContext 方法）