## docker-compose搭建基础设施

> 包括mysql、mongodb、redis，主要用于测试和玩耍

```yaml
version: "3"

services:
  mongo:
    image: mongo
    container_name: mongo
    ports:
      - "27017:27017"
    volumes:
      - /etc/localtime:/etc/localtime:ro
      - ./mongodb:/data/db
    environment:
      MONGO_INITDB_ROOT_USERNAME: root      # 用户名
      MONGO_INITDB_ROOT_PASSWORD: 123456    # 密码

  mysql:
    image: mysql
    container_name: mysql
    ports:
      - "3306:3306"
    volumes:
      - "./mysql:/var/lib/mysql"
    environment:
      MYSQL_ROOT_PASSWORD: 123456

  redis:
    image: redis
    container_name: redis
    ports:
      - "6379:6379"
    volumes:
      - /etc/localtime:/etc/localtime:ro
      - ./redis:/data
```

### 注意点（持续更新

#### 设定密码

虽然是测试用，最好还是设定一下密码。

#### redis配置文件

`redis.conf`

```
requirepass 123456
#daemonize yes
bind 0.0.0.0
appendonly yes
```

