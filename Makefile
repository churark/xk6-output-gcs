.PHONY: build
build:
	xk6 build --with $(shell go list -m)=.

.PHONY: run-example
run-example: build
run-example:
	./k6 run _example/test.js --out gcs --iterations 2
