#Backend build
FROM golang as backbuild
WORKDIR build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build  -o /backend main.go


#Frontend build
FROM node as frontbuild
ENV NODE_OPTIONS=--openssl-legacy-provider
WORKDIR build
COPY /client .
RUN npm install
RUN npm run build

#Production
FROM scratch
WORKDIR /app
COPY --from=backbuild /backend /app
COPY --from=frontbuild /build/dist /app/static
EXPOSE 8080
CMD ["/app/backend" ]