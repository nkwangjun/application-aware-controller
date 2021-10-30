# Copy the controller into a thin image
From alpine:3.12.0

ADD application-aware-controller /usr/local/bin

WORKDIR /root/

COPY entrypoint.sh .
RUN chmod +x /root/docker-entrypoint.sh

ENTRYPOINT  ["/root/entrypoint.sh"]