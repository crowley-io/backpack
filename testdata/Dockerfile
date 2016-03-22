FROM golang:1.6

ENV CROWLEY_PACK_USER 1002
ENV CROWLEY_PACK_GROUP 1002
ENV CROWLEY_PACK_DIRECTORY /usr/local/app
ENV CROWLEY_PACK_OUTPUT foo

WORKDIR /go/src/github.com/crowley-io/backpack

ADD testdata/backpack-test /usr/local/bin/
RUN chown root:root /usr/local/bin/backpack-test && chmod 0744 /usr/local/bin/backpack-test
CMD /usr/local/bin/backpack-test
