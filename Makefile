get-protos-submodule:
	@echo "initializing protos submodule..." && \
	git submodule add --progress --force https://github.com/qreeket/protos.git protos

update-protos-submodule:
	@echo "initializing protos submodule..." && \
	git submodule update --init --recursive

gen-protos:
	@echo "generating protos for microservices..." && \
	chmod +x gen-protos.sh && ./gen-protos.sh

.PHONY: get-protos-submodule update-protos-submodule