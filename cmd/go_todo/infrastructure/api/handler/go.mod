replace local.packages/model => ../../../model


replace local.packages/repository => ../../../usecase/repository

module github.com/masibw/go_todo/cmd/go_todo/infrastructure/api/handler

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/masibw/go_todo/pkg/utility v0.0.0-20200727031924-c5ea13bdec3e
	local.packages/model v0.0.0-00010101000000-000000000000
	local.packages/repository v0.0.0-00010101000000-000000000000
)
