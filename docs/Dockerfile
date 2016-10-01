FROM grafana/docs-base:master

ENV PROJECT=grafana

# to get the git info for this repo

COPY . /src

RUN rm -rf /docs/content/$PROJECT/

COPY . /docs/content/$PROJECT/
