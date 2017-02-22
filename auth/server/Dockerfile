FROM scratch
MAINTAINER Emilie Sicilia <emilie.sicilia@gmail.com>
MAINTAINER Romain Baugue <romain.baugue@elwinar.com>

ADD auth /app

ENV APP_DEBUG false
ENV APP_PASSWORD_COST 12
ENV APP_PORT 80

EXPOSE 80

ENTRYPOINT ["/app"]
