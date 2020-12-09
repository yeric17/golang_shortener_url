FROM golang:1.15.2-alpine3.12
WORKDIR /app/url-shortener
ENV GOPATH=/app
COPY . /app/url-shortener
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/rs/cors
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/joho/godotenv
RUN go build -o main .
CMD ["/app/url-shortener/main"]