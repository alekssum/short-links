FROM "mysql:5.7"

ENV MYSQL_DATABASE=shortlinks
ENV MYSQL_ROOT_PASSWORD=qwerty

ADD init_schema.sql /docker-entrypoint-initdb.d

EXPOSE 3306
