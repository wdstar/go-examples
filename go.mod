module github.com/wdstar/go-examples

go 1.13

require (
	github.com/kr/pretty v0.2.0 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/wesovilabs/koazee v0.0.4
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15
)

// for koazee's broken dependencies
replace github.com/golangci/golangci-lint v1.12.3 => github.com/golangci/golangci-lint v1.23.7
