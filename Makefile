# roscoe

TABWIDTH=4

fmt:
	gofmt -w -tabs=false -tabwidth=$(TABWIDTH) *.go
