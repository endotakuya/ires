FROM ruby:2.4
RUN apt-get update -qq && apt-get install -y \
    build-essential libpq-dev apt-transport-https vim gettext-base \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN gem install bundler

WORKDIR /ires
COPY . /ires

RUN bundle install

