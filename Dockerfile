FROM scratch

WORKDIR /app
COPY hello-pods /app/hello-pods
COPY index.html.tpl /app/index.html.tpl

CMD ["/app/hello-pods"]
