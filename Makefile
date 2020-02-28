TEST_DIR ?= tmp/test-results

default: build

deps:
	# Dependencies for this makefile, not for the program itself. Those are
	# handled by `go mod`
	go install gotest.tools/gotestsum
	go install golang.org/x/lint/golint

build:
	go build ./...

check:
	golint ./...
	go vet ./...

test:
	go test ./...

testci:
	rm -rf ${TEST_DIR}
	mkdir -p ${TEST_DIR}
	gotestsum --junitfile ${TEST_DIR}/gotestsum-report.xml -- -coverprofile=tmp/coverage.out -race ./...
	go tool cover -html=tmp/coverage.out -o tmp/coverage.html

testclean:
	rm -rf ${TEST_DIR}

clean: testclean
	rm -f prog
