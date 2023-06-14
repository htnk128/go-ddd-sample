FROM mysql:8.0.33

EXPOSE 3306
ENV MYSQL_ROOT_USER root
ENV MYSQL_ROOT_PASSWORD password
ENV MYSQL_USER user
ENV MYSQL_PASSWORD password
ENV TZ 'Asia/Tokyo'

# Set MySQL character and max allowed packet
RUN { \
    echo '[client]'; \
    echo 'default-character-set=utf8mb4'; \
    echo '[mysqld]'; \
    echo 'max_allowed_packet=32MB'; \
    echo 'character-set-server=utf8mb4'; \
    echo 'collation-server=utf8mb4_bin'; \
} > /etc/mysql/conf.d/charset.cnf
