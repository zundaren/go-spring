module github.com/go-spring/spring-go-redis

go 1.14

require (
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-spring/spring-base v1.1.0-beta.0.20211001035852-bfba805daa15
	github.com/go-spring/spring-core v1.0.6-0.20211001040940-f4fed6e6c943
)

replace github.com/go-spring/spring-core => ../spring-core
