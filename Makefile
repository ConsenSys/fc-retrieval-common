# Copyright (C) 2020 ConsenSys Software Inc.

COV?=80

default: clean utest

utest:
	go test ./...

coverage:
	bash ./scripts/coverage.sh $(COV)

clean:
#nothing to do

useremote:
	cd scripts; bash use-remote-repos.sh

.PHONY: default clean utest coverage useremote
