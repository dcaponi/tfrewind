test:
	go get github.com/jpoles1/gopherbadger
	gopherbadger -md="readme.md" -png=false

ship:
	bash ship.sh github.com/dcaponi/tfrewind
