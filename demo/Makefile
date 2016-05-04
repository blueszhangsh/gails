target=release

$(target)/demo:
	go build -ldflags "-s" -o $@ main.go

clean:
	-rm -r $(target)
