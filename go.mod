module github.com/masibw/go_todo

go 1.14

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.15
	github.com/masibw/go_todo/pkg/db v0.0.0-20200726122544-9abcfde1d13a
	github.com/masibw/go_todo/pkg/utility v0.0.0-20200727031924-c5ea13bdec3e // indirect
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
	local.packages/handler v0.0.0-00010101000000-000000000000
	local.packages/model v0.0.0-00010101000000-000000000000
	local.packages/persistance v0.0.0-00010101000000-000000000000
	local.packages/repository v0.0.0-00010101000000-000000000000 // indirect
	local.packages/utility v0.0.0-00010101000000-000000000000 // indirect
)

replace local.packages/utility => ./pkg/utility

replace local.packages/model => ./cmd/go_todo/model

replace local.packages/handler => ./cmd/go_todo/infrastructure/persistance

replace local.packages/persistance => ./cmd/go_todo/infrastructure/api/handler

replace local.packages/repository => ./cmd/go_todo/usecase/repository
