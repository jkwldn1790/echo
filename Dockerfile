FROM cgr.dev/chainguard/go AS build
COPY . /echo
RUN cd /echo && go env -w GO111MODULE=auto && go build -o /echo

FROM cgr.dev/chainguard/wolfi-base:latest
COPY --from=build . .
COPY --from=build echo/templates /templates
CMD [ "/echo/echo" ]