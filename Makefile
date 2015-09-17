NAME := todoist
VERSION := 0.0.1

all:
	go install -ldflags "-X main.name=$(NAME) -X main.version=$(VERSION)" .

install:
	go get github.com/jawher/mow.cli
	go get github.com/mitchellh/go-homedir
	go get github.com/twinj/uuid

release: clean
	mkdir -p release
	@for os in linux darwin freebsd; do \
		for arch in 386 amd64; do \
			echo "Building $(NAME) v$(VERSION) for $$os-$$arch."; \
			env GOOS=$$os GOARCH=$$arch go build -o release/$(NAME)-$(VERSION)-$$os-$$arch -ldflags "-X main.name=$(NAME) -X main.version=$(VERSION)"; \
		done \
	done

clean:
	rm -rf release
