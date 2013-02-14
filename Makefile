# roscoe

include $(GOROOT)/src/Make.$(GOARCH)

TABWIDTH=4

#TARG=goargcfg.googlecode.com/hg/argcfg
GOFILES=\
	osc.go

include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w -tabs=false -tabwidth=$(TABWIDTH) *.go

osctest: package osc.go
	$(GC) -o main.$(O) -I_obj osc.go
	$(LD) -o osctest -L_obj main.$(O)