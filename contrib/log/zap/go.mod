module github.com/sql668/go-toolkit/contrib/log/zap

go 1.18

require (
	github.com/sql668/go-toolkit v0.0.0-20240527125930-705d3cf9dfa4
	go.uber.org/zap v1.27.0
)

require go.uber.org/multierr v1.11.0 // indirect

replace (
	github.com/sql668/go-toolkit => ../../../
)