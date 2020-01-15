FROM golang:1.13.4-alpine
WORKDIR /app
COPY . .
RUN go mod download\
&& go build -o main .
CMD [ "/app/main" ]