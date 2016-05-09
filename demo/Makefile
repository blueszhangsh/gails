target=release

$(target)/demo:
	go build -ldflags "-s" -o $@ main.go


reactjs:
	cd reactjs && npm run build
	mv reactjs/public $(target)

clean:
	-rm -r $(target) 


