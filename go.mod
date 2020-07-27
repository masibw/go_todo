module github.com/masibw/go_todo

go 1.14

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.15
	github.com/masibw/go_todo/pkg/db v0.0.0-20200726122544-9abcfde1d13a
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
	local.packages/utility v0.0.0-00010101000000-000000000000
)

replace local.packages/utility => ./pkg/utility
