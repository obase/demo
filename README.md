# package demo

obase工具箱的救命项目.

# 工具介绍
- conf: 兼容推栏828的yaml配置conf.yml, 完成代码外的配置设置. 与Java/PHP/Python等不同, go是编译型, 代码编译打包后就是一个二进制可执行文件, 无法交"程序，资源，配置"混合.
- kit: 常用辅助工具,例如http,ipv4等
- center: 服务发现/健康检查的封装, 目前是基于consul实现. 并提供服务发现相关的http, grpc等调用客户端工具.
- cmd: 命令行工具调用骨架.
- log: 自实现的行式日志工具,参考glog核心,实现日志轮转(大小/日期). 由于舍弃业务基本不会用到的V级别, 在性能方法弱胜于glog.
- mysql: 主要是提供半成品ORM骨架, 结合proto代码工具实现类似Mybatis/Hibernate等Java的ORM功能. 目前内部使用kingshard, 完全不支持事务与预编译, 而且只能使用mymysql才能适配, 如果使用新版本go-sql-driver, 无法与运营哪套kingshard兼容.
- mongodb: 封装官网驱动mongo-go-driver, 提供常用简易API. 同时保证原驱动底层客户端的扩展性. 说白了, 就是golang的mongo驱动API过于烦琐, 自己根据业务常用的情景封装了一批API, 从而提高性能.
- clickhouse: 集成clickhouse的sql接口支持, 与mysql的API接口一致! 如果不依赖事务的话, 切换mysql与clickhouse的驱动就可以实现存储引擎的切换.
- redis.v2: 基于redigo提供简化版API, 并实现redis集群的驱动, 单例与集群之间切换修改配置"cluster: true"即可. 同时支持pika
- pbapi: 基于gin的扩展:
    1. 支持多渠道访问: 从proto的service自动生成grpc, http, websocket访问接口.
    2. 支持配置式插件: 在conf.yml为http接口动态添加access, cache, proxy等常用插件. 
    3. 支持自动注册/解注册: 服务启动时自动注册到consul, 关闭时自动解注册.

# 示例项目

## 命令行工具

## 微服务(http, websocket, grpc)

