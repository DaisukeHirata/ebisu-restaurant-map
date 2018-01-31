all:
		cd $(CURDIR)/functions/create && GOOS=linux go build -o ../../bin/create
		cd $(CURDIR)/functions/get && GOOS=linux go build -o ../../bin/get