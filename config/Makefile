
all:	lint

clean:
	(cd test && $(MAKE) clean)

lint:
	(cd test && $(MAKE) lint)
	@echo "  - Linting env config files ..."
	@shellcheck -e SC2034 *.env || exit 78
	@echo "  - Passed lint checks."

test:	tests
tests:
	(cd test && $(MAKE) tests)

config:
	(cd test && $(MAKE))


.PHONY:	clean lint test tests config
