
PREFIX=github.com/petelliott/shooty-game
PLUGINS=plugins/core.so

.PHONY: all clean

all: config.json $(PLUGINS)

clean:
	rm -f config.json
	rm -rf plugins

config.json: ../config
	d2json $< > $@

plugins/%.so: ../plugins/% | plugins
	go build -buildmode=plugin -o $@ $(PREFIX)/plugins/$*

plugins:
	mkdir -p plugins
