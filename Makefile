.PHONY: present
present:
	go run golang.org/x/tools/cmd/present -notes
# For code highlightling use Chrome + https://github.com/josephbuchma/Go-Present-code-highlighter

.PHONY: test
test:
	go test -count=1 ./src

.PHONY: ginkgo
ginkgo:
	go run github.com/onsi/ginkgo/ginkgo -r -noisyPendings=false
