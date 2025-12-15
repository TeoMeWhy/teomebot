.PHONY: test
test:
	cd repositories && go test -v .
	cd services && go test -v .
	cd controllers && go test -v .