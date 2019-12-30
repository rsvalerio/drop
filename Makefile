VERSION:=0.1
TAG:=v$(VERSION)

COVEROUT = cover.out
GOFMTCHECK = test -z `gofmt -l -s -w *.go | tee /dev/stderr`
GOTEST = go test -v
COVER = $(GOTEST) -coverprofile=$(COVEROUT) -covermode=atomic -race
GOPATH?=$(HOME)/go

all: fmt test

.PHONY: build
build: fmt

.PHONY: push
push:
	@echo "Starting push..."
	@echo "Nothing to push..."
	@echo "Finishing push..."

.PHONY: fmt
fmt:
	@echo "Checking format..."
	@$(GOFMTCHECK)

.PHONY: test
test:
	@echo "Running tests..."
	@$(COVER)

# Docker targets
.PHONY: docker
docker:
	rm coredns || true
	cp /home/rodrigo/projects/coredns/coredns .
	docker build -t coredns-drop . --no-cache

.PHONY: run
run:
	docker run \
		-p 25353:53/udp \
		-p 15353:15353 \
		-v $(shell pwd)/Corefile:/etc/coredns/Corefile \
		--rm --name coredns-drop \
		coredns-drop

# Use the 'release' target to start a release
.PHONY: release
release: commit push
	@echo Released $(VERSION)

.PHONY: commit
commit:
	@echo Committing release $(VERSION)
	git commit -am"Release $(VERSION)"
	git tag $(TAG)
