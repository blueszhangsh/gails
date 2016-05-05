target=release

$(target)/demo:
	go build -ldflags "-s" -o $@ main.go
	npm run build
	mv public $(target)

clean:
	-rm -r $(target) public
