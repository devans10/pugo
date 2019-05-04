TEST?=./pure1 ./flasharray
GOFMT_FILES?=$$(find . -name '*.go' |grep -v pkg)

default: build

build: fmtcheck
	go install

test: fmtcheck
	go test $(TEST) -coverprofile=coverage.txt -covermode=atomic

testacc: fmtcheck
	PURE_ACC=1 go test $(TEST) -v -timeout 120m -parallel=4 -cover

vet:
	@echo "go vet ."
	@go vet $(TEST) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"


.PHONY: build test testacc vet fmt fmtcheck errcheck

