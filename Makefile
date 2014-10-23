# roscoe

# include $(GOROOT)/src/Make.$(GOARCH)

GOFILES=\
	osc.go

# include $(GOROOT)/src/Make.pkg

fmt:
	gofmt -w *.go

osctest: package osc.go
	$(GC) -o main.$(O) -I_obj osc.go
	$(LD) -o osctest -L_obj main.$(O)
