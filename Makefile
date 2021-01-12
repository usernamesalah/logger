.PHONY: build
build:
	@echo "> Building the server binary ..."
	@rm -rf bin && go build -o bin/logger .

.PHONY: run
run: build
	@echo "> RUN the server binary ..."
	@./bin/logger

.PHONY: push
push: 
	@git add .
	@git commit -m "update"
	@git push origin master
