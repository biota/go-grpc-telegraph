CONFIGS = service device/bootstrap device

all:	build

clean:	clean-tls-config

build:	generate-tls-config

lint:
	@echo "  - Linting shell scripts ..."
	shellcheck generate-tls-config
	@echo "  - shellcheck return code = $$?"
	@echo "  - Passed lint checks."

test:	tests
tests:	lint


#
#  Generate TLS configuration.
#
generate-tls-config:	clean-tls-config
	@echo "  - Generating test TLS configuration ..."

	#
	# For more nbits in the key, use:
	#    NEWKEY_ARG="rsa:4096" ./generate-tls-config
	#
	./generate-tls-config

	@echo "  - Generated test TLS configuration."

clean-tls-config:
	@echo "  - Cleaning out test TLS configuration ..."
	$(foreach cfg, $(CONFIGS), ($(RM) tls/$(cfg)/bundle/* tls/$(cfg)/*.*);)
	@echo "  - Cleaned test TLS configuration."


.PHONY:	clean build lint test
.PHONY:	generate-tls-config clean-tls-config
