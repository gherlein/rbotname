NAME   = rbotname
FILE   = ./${NAME}

build: dependencies 
	go build 

${FILE}: build

dependencies:
	go get ./...

clean:
	-rm -f ${FILE}
	-rm -f *~

run:
	sudo ${FILE} 



