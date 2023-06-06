FROM httpd:2.4

COPY ./index.html /temp/

ARG SERVER_ID

RUN sed "s/SERVER_ID/$SERVER_ID/" /temp/index.html > /usr/local/apache2/htdocs/index.html
