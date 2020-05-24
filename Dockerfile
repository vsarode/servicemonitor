FROM alpine:edge AS build
RUN apk update
RUN apk upgrade
RUN apk add --update go=1.13.11-r0 gcc=9.3.0-r2 g++=9.3.0-r2
RUN apk add build-base
WORKDIR /app
ENV GOPATH /app
ADD . /app/src/servicemonitor
WORKDIR src/servicemonitor
RUN CGO_ENABLED=1 GOOS=linux go install -a .

# Build the React application
FROM node:alpine AS node_builder
COPY --from=build /app/src/servicemonitor/web ./
RUN npm install
RUN npm run build
# Final stage build, this will be the container
# that we will deploy to production
FROM alpine:edge
RUN apk update
RUN apk upgrade
RUN apk add gcc=9.3.0-r2 g++=9.3.0-r2
RUN apk add build-base
COPY --from=build /app/bin/servicemonitor .
COPY --from=node_builder /build ./web/build
RUN chmod +x servicemonitor
EXPOSE 8080
CMD ./servicemonitor