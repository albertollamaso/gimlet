FROM --platform=$BUILDPLATFORM alpine:3.16

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

ENV DATABASE_DRIVER=sqlite3
ENV DATABASE_CONFIG=/var/lib/gimlet-dashboard/gimlet-dashboard.sqlite
ENV XDG_CACHE_HOME /var/lib/gimlet-dashboard

RUN addgroup -S gimlet-dashboard && adduser -S gimlet-dashboard -G gimlet-dashboard

ADD docker/dashboard/known_hosts /etc/ssh/ssh_known_hosts

RUN mkdir /var/lib/gimlet-dashboard
WORKDIR /gimlet-dashboard

ARG TARGETPLATFORM
ARG BUILDPLATFORM

COPY --chown=gimlet-dashboard:gimlet-dashboard bin/${TARGETPLATFORM}/gimlet-dashboard gimlet-dashboard
COPY --chown=gimlet-dashboard:gimlet-dashboard web/dashboard/build ./web/build/

USER gimlet-dashboard

RUN git config --global user.name "Gimlet Dashboard"
RUN git config --global user.email "gimlet-dashboard@gimlet.io"

EXPOSE 9000
CMD ["/gimlet-dashboard/gimlet-dashboard"]
