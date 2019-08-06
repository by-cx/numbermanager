build:
	docker build -t creckx/numbermanager:$(shell cat VERSION) .
	docker tag creckx/numbermanager:$(shell cat VERSION) creckx/numbermanager:latest

push: build
	docker push creckx/numbermanager:$(shell cat VERSION)
	docker push creckx/numbermanager:latest