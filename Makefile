init:
	git submodule update --init --recursive
	cd ./tests/aws-kms-xksproxy-test-client/ && $(MAKE) -f Makefile docker

test:
	./tests/test.sh
