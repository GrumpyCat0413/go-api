# go-api

go开发后台api服务学习
### 一
- 使用GIN 启动一个带有健康检查的hello world 级别的 API后台服务

### 二 

- 使用配置管理工具viper解析配置文件并读取配置、及配置文件的监控与热更新
- 使用pflag.parse() 来解析命令行 参数


### 三
- 使用lexkong/log 日志库 进行对日志的配置和管理，还能够按天或者容量来进行转存压缩日志


### 四
- Gorm 操作mysql 

### 五
- 自定义error code 信息

### 六
- 封装 SendResponse
- 创建 user的 CreateRequest 与 CreateResponse 尽量每一个接口都要创建与之对应的struct

### 使用第三方库
- gin  
- gopsutil ：电脑内存、cpu、硬盘等信息
- pflag ：命令行参数 解析
- viper ：配置文件读取及配置文件的热更新
- lexkong/log 日志库 可以进行自动转存
- Gorm mysql 数据库ORM框架