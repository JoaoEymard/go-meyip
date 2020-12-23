FROM golang:1.15
# ----------
# TimeZone
# ----------
ENV TZ=America/Fortaleza
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
# ----------
# Env variables
# ----------
# ENV GOPATH /app
# ENV DATADIR /dat
# ----------
# Work dir
# ----------
RUN mkdir -p /app
ADD . /app
WORKDIR /app
# ----------
# Build App
# ----------
RUN go get github.com/gorilla/mux
RUN go build -o audit .
# ----------
# Run App
# ----------
CMD ["./audit"]