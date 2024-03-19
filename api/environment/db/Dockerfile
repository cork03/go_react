FROM mysql:8.3.0

ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=factory
ENV MYSQL_USER=factory
ENV MYSQL_PASSWORD=factory
COPY ./db/my.cnf /etc/mysql/conf.d/my.cnf
# COPY ./environment/db/initdb.d /docker-entrypoint-initdb.d
