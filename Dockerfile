FROM golang:1.22.4 as builder
WORKDIR ./
COPY ./ /app

ARG APP_VERSION
RUN export GO111MODULE="on"
RUN export GOPROXY=https://goproxy.io,direct
ENV CA_PORT=80
WORKDIR /app
RUN go get github.com/gin-gonic/gin/binding@v1.9.1
RUN go mod vendor
RUN make APP_VERSION=${APP_VERSION}

#RUN go mod vendor
#RUN go build -mod=vendor -o easyca -a -ldflags "-extldflags '-static' -s -X 'main.AppName=${APP_NAME}' \
# 				-X 'conf.AppVersion=${APP_VERSION}' \ 
# 				-X 'conf.BuildVersion=${BUILD_VERSION}' \
# 				-X 'conf.BuildTime=${BUILD_TIME}' \
# 				-X 'conf.GitRevision=${GIT_REVISION}' \
# 				-X 'conf.GitBranch=${GIT_BRANCH}' \
# 				-X 'conf.GoVersion=$GO_VERSION'" 



FROM alpine:latest
WORKDIR /

COPY --from=builder /app/cmd/idx /app/
COPY --from=builder /app/conf ./conf
COPY --from=builder /app/certs ./certs

COPY --from=builder /app/static ./static
#cluster port
EXPOSE 80
CMD ["/app/idx"]
