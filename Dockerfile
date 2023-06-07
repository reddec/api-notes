FROM scratch
# 8080 for API
EXPOSE 8080/tcp
VOLUME /data
ENV BIND=":8080" \
    PUBLIC_URL="http://127.0.0.1:8081" \
    DIR="/data"
ENTRYPOINT ["/api-notes"]
ADD api-notes /