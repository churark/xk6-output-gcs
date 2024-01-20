.PHONY: build
build:
	xk6 build --with xk6-output-gcs=.

.PHONY: run-example
run-example: build
run-example:
	./k6 run _example/test.js --out gcs --quiet --no-summary --iterations 2
