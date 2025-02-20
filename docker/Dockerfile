# Step 1: Use Alpine to install rsyslog and required libraries
FROM alpine:latest AS builder
RUN apk update && apk add --no-cache rsyslog util-linux tini

# Step 2: Use BusyBox as the final base image
FROM busybox:latest

# Copy tini for proper process management
COPY --from=builder /sbin/tini /sbin/tini
RUN chmod +x /sbin/tini

# Copy rsyslog binary and configuration
COPY --from=builder /usr/sbin/rsyslogd /usr/sbin/rsyslogd
COPY --from=builder /etc/rsyslog.conf /etc/rsyslog.conf

# Copy required shared libraries
COPY --from=builder /lib /lib
COPY --from=builder /usr/lib /usr/lib

# Ensure necessary directories exist
RUN mkdir -p /var/run /var/log /var/lib/rsyslog && touch /var/log/syslog

# Modify rsyslog.conf:
# - Comment out the default rule that logs to kern.log
# - Redirect kernel logs to /var/log/syslog
RUN sed -i 's|^kern\.*.*kern.log|# &|' /etc/rsyslog.conf && \
    echo 'kern.* /var/log/syslog' >> /etc/rsyslog.conf && \
    sed -i 's/^module(load="imklog")/#module(load="imklog")/' /etc/rsyslog.conf

# Copy the Golang logging daemon binary
COPY target/logger-daemon /usr/local/bin/logger-daemon
RUN chmod +x /usr/local/bin/logger-daemon

COPY target/client-app /usr/local/bin/app1
RUN chmod +x /usr/local/bin/app1
COPY target/client-app /usr/local/bin/app2
RUN chmod +x /usr/local/bin/app2
COPY target/client-app /usr/local/bin/app3 
RUN chmod +x /usr/local/bin/app3

# Set tini as the init system and run both services
ENTRYPOINT ["/sbin/tini", "--"] 
# Start rsyslogd and the Golang logger daemon 
CMD ["/bin/sh", "-c", "/usr/sbin/rsyslogd -n & exec /usr/local/bin/logger-daemon"]
