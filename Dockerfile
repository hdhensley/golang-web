FROM node:alpine as frontend
WORKDIR /frontend
COPY package.json .
COPY package-lock.json .
RUN npm install
COPY app/resources/styles.css .
RUN ./node_modules/.bin/tailwindcss build -o computed_styles.css

FROM golang:alpine as builder

WORKDIR /app
ENV GOPATH /go

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .
RUN templ generate
COPY --from=frontend /frontend/computed_styles.css public/styles.css
RUN go build -o /go/bin/app app/*.go

FROM alpine as final

COPY --from=builder /go/bin/app /app

EXPOSE 3000

CMD ["/app"]
