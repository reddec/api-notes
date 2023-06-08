FROM scratch
# 8080 for API
EXPOSE 8080/tcp
VOLUME /data
ENV API_NOTES_BIND=":8080" \
    API_NOTES_PUBLIC_URL="http://127.0.0.1:8081" \
    API_NOTES_DIR="/data"
ENTRYPOINT ["/api-notes"]
ADD api-notes /