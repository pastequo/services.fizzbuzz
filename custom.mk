NAME=fizzbuzz

.PHONY: run generate

run: run.fizzbuzz
generate: generate.mocks

#####################
# Mock generation   #
#####################

.PHONY:  generate.mocks

#help generate.mocks: Generate/update mocks for testing
generate.mocks:
	$(TOOLS_DIR)/mockgen -source=$(CURDIR)/internal/repo/interfaces.go -destination=$(CURDIR)/internal/repo/mock/mocker.go
	$(TOOLS_DIR)/mockgen -source=$(CURDIR)/internal/usecase/interfaces.go -destination=$(CURDIR)/internal/usecase/mock/mocker.go


#####################
# Custom Run        #
#####################

LIMIT=3
W1_W=fizz
W1_N=2
W2_W=buzz
W2_N=3

.PHONY: run.fizzbuzz run.stats run.metrics

#help run.fizzbuzz: call fizzbuzz route
run.fizzbuzz:
	curl -i localhost:8080/algo/fizzbuzz \
	-X POST \
	-H "Content-Type: application/json" \
	-d '{"limit": $(LIMIT), "word1": {"word": "$(W1_W)", "multiple": $(W1_N)}, "word2": {"word": "$(W2_W)", "multiple": $(W2_N)}}'

#help run.stats: call fizzbuzz stat route
run.stats:
	curl -i localhost:8080/algo/fizzbuzz/stats \
	-X GET \
	-H "Content-Type: application/json"

#help run.metrics: get prometheus metrics
run.metrics:
	curl -i localhost:7777/metrics
