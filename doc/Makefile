SCHEMES=\
interfaces.png\
mswfAPI.png\
files.png\

%.png: %.dot
	dot -Tpng $< -o $@

all: $(SCHEMES)
	echo done

clean:
	rm -f $(SCHEMES)

