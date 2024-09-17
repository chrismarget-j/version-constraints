check-repo-clean:
	git update-index --refresh && git diff-index --quiet HEAD --

compliance:
	go run github.com/chrismarget-j/go-licenses save   --ignore github.com/Juniper --save_path Third_Party_Code --force ./... || exit 1 ;\
	go run github.com/chrismarget-j/go-licenses report --ignore github.com/Juniper --template .notices.tpl ./... > Third_Party_Code/NOTICES.md || exit 1 ;\

.PHONY: all compliance compliance-check
