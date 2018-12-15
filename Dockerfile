FROM alpine

COPY dist/web_blog /bin/

EXPOSE 8005

ENTRYPOINT [ "/bin/web_blog" ]
