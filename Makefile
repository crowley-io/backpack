VERSION=0.1.2
NAME=backpack
ORGANIZATION=crowley-io
PACKAGE=github.com/${ORGANIZATION}/${NAME}

GITHUB_USER=${ORGANIZATION}
GITHUB_REPO=${NAME}

ARTIFACTS = \
	crowley-${NAME}_linux-amd64

UPLOAD_CMD = github-release upload --user ${GITHUB_USER} --repo ${GITHUB_REPO} --tag "v${VERSION}" \
	--name ${FILE} --file ${FILE}

LDFLAGS="-X ${PACKAGE}/engine.Version=v${VERSION}"

DOCKER_TEST_IMAGE=backpack-test
DOCKER_TEST_DOCKERFILE=testdata/Dockerfile

all: ${NAME}

test:
	docker build -t $(DOCKER_TEST_IMAGE) -f $(DOCKER_TEST_DOCKERFILE) .
	docker run --rm -it -v $(CURDIR):/go/src/$(PACKAGE) $(DOCKER_TEST_IMAGE)

style:
	gofmt -w .

lint:
	gometalinter --vendor --disable=gotype --enable=lll --enable=gofmt \
		--dupl-threshold=80 --deadline=10s --line-length=120 --tests ./...

${NAME}:
	go build -ldflags ${LDFLAGS} -o ${NAME}

clean:
	rm -rf ${NAME}

install: ${NAME}
	install -o root -g root -m 0755 ${NAME} /usr/local/bin/crowley-${NAME}

release: artifacts
	git tag "v${VERSION}" && git push --tags
	github-release release --user ${GITHUB_USER} --repo ${GITHUB_REPO} --tag "v${VERSION}" \
		--name ${VERSION} --pre-release
	$(foreach FILE,$(ARTIFACTS),$(UPLOAD_CMD);)

artifacts:
	gox -osarch="linux/amd64" -ldflags ${LDFLAGS} -output="crowley-${NAME}_{{.OS}}-{{.Arch}}"

.PHONY: clean ${NAME} install artifacts test style lint release
