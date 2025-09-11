check-repo-clean:
	git update-index --refresh && git diff-index --quiet HEAD --

compliance:
	@sh -c "$(CURDIR)/.ci/scripts/compliance.sh"

compliance-check: compliance check-repo-clean

.PHONY: all compliance compliance-check
