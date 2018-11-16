FROM ruby:2.4

# Timezone
ENV TZ Asia/Tokyo

RUN apt-get update -qq \
 && apt-get install -y --no-install-recommends \
        build-essential \
        libpq-dev \
        apt-transport-https \
        vim \
        gettext-base \
        gcc \
        libc6-dev \
        make \
        nodejs \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

# Go
ENV GOLANG_VERSION 1.11.2
ENV GOLANG_DOWNLOAD_URL https://dl.google.com/go/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 1dfe664fa3d8ad714bbd15a36627992effd150ddabd7523931f077b3926d736d

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
 && echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
 && tar -C /usr/local -xzf golang.tar.gz \
 && rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
COPY go-wrapper /usr/local/bin

# pkg管理用
RUN go get -u github.com/golang/dep/cmd/dep

# App
WORKDIR /go/src/github.com/endotakuya/ires
COPY . .

RUN gem install bundler rake-compiler && bundle install

VOLUME [ "/usr/local/bundle" ]
VOLUME [ "/go/src" ]

EXPOSE 3000

