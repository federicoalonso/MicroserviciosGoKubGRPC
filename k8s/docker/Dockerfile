FROM golang

WORKDIR /app

COPY authsvc /app/authsvc
COPY apisvc /app/apisvc

EXPOSE 9015
EXPOSE 9016

CMD [ "/app/authsvc", "/app/apisvc" ]