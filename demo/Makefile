target=release

$(target)/gails:
	go build -ldflags "-s" -o $@ main.go

clean:
	-rm -r $(target)
