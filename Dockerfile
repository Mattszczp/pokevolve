FROM node:18 as node
WORKDIR /app
COPY . /app
RUN npm i
RUN npx tailwindcss -i ./static/styles.css -o ./static/styles.min.css --minify

FROM golang:1.18 as build
WORKDIR /app
COPY . /app
COPY --from=node /app/static/styles.min.css ./static/styles.min.css
RUN go build -a -v -o pokevolve .

FROM scratch
WORKDIR /app
COPY --from=build /app/pokevolve ./
CMD ["/app/pokevolve"]
