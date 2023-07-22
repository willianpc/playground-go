LIBP = $(LIB_PATH)

has_path:
ifeq ($(LIBP),)
	@echo "A LIB_PATH environment variable is needed."
	@echo "eg '.' for core module, 'instagin', 'instaredis/v2'"
	exit 1
endif

release: updated_version
	@echo "making a release '$(LIBP)'"

release_tag: has_path
	@echo "generating a tag"

updated_version: release_tag
	@echo "updating version.go"

major:
	@echo "this is major version"

# .PHONY: clean