OUTPUT_FOLDER=./bin
CMD_FOLDER=./cmd

build_cash bc:
	@CGO_ENABLED=0 go build -o "${OUTPUT_FOLDER}/cash" "${CMD_FOLDER}/cash"

clear:
	@rm -rf "${OUTPUT_FOLDER}"

.PHONY: build_cash clear
