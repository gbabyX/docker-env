FROM php:7.2-fpm

LABEL maintainer="Walker Xiong <exchange@walkerx.cn>"

RUN apt-get update -yqq && \
    apt-get install -y apt-utils && \
    pecl channel-update pecl.php.net

ARG PHP_EXT_MYSQLI=false

RUN if [ ${PHP_EXT_MYSQLI} = true ]; then \
    docker-php-ext-install mysqli \
;fi

ARG PHP_EXT_REDIS=false

RUN if [ ${PHP_EXT_REDIS} = true]; then \
	docker-php-ext-install redis \
;fi

ARG PHP_EXT_PDO_MYSQL=false

RUN if [ ${PHP_EXT_PDO_MYSQL} = true ]; then \
    docker-php-ext-install pdo_mysql \
;fi

RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    rm /var/log/lastlog /var/log/faillog

EXPOSE 9000