FROM php:7.2-fpm

RUN apt-get update -yqq && \
    apt-get install -y apt-utils && \
    pecl channel-update pecl.php.net


ARG PHP_EXT_MYSQLI=false

RUN if [ ${PHP_EXT_MYSQLI} = true ]; then \
    docker-php-ext-install mysqli \
;fi

RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    rm /var/log/lastlog /var/log/faillog

EXPOSE 9000