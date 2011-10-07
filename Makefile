include $(GOROOT)/src/Make.inc

TARG=buffer

GOFILES=buffer.go

#CGO_OFILES=buffer_c.o

CLEANFILES+=buffer

include $(GOROOT)/src/Make.pkg

