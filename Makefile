.PHONY: release release.snapshot

setup:
	which goreleaser > /dev/null 2>&1 || go get -u github.com/goreleaser/goreleaser

version.list:
	@ echo 'latest releases:'
	@ git tag -l | sort -nr | head -5

version.new: version.list
ifndef NEW_VERSION
	$(eval NEW_VERSION = $(shell \
	  read -p $$'\e[33mPlease enter the version name for the new release\e[0m: ' new_version; \
	  echo $${new_version} \
	 ))
	git tag -a $(NEW_VERSION) -m ""
	git push origin $(NEW_VERSION)
endif

release: version.new
	goreleaser --rm-dist

release.snapshot:
	goreleaser --snapshot --rm-dist
