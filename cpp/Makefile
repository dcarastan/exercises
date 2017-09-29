IDIR =../include
CC=g++
CFLAGS=-I$(IDIR)

ODIR=obj
LDIR =../lib

LIBS=-lm

#_DEPS = cpp_test.h
DEPS = $(patsubst %,$(IDIR)/%,$(_DEPS))

_OBJ = cpp_test.o
OBJ = $(patsubst %,$(ODIR)/%,$(_OBJ))

.DEFAULT_GOAL = all

all: cpp_test

$(ODIR)/%.o: %.cpp $(DEPS)
	@echo "***** Compiling $@"
	mkdir -p $(ODIR)
	$(CC) -c -o $@ $< $(CFLAGS)

cpp_test: $(OBJ)
	@echo "***** Linking $@"
	g++ -o $@ $^ $(CFLAGS) $(LIBS)

.PHONY: clean

clean:
	rm -f $(ODIR)/*.o *~ core $(INCDIR)/*~



