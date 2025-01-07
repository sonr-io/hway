
# Use the dendrite-monolith image from matrixdotorg
FROM matrixdotorg/synapse:latest

RUN mkdir -p /data

# Copy App Services 
COPY apps/hooks.sonr.chat/registration.yml /data/svc/hookshot.yml

# Copy the hookshot passkey
COPY apps/sonr.chat/server.crt /data/server.crt
COPY apps/sonr.chat/server.key /data/server.key
COPY apps/sonr.chat/signing.key /data/sonr.chat.signing.key

# Copy the synapse configuration
COPY config/synapse.yaml /data/homeserver.yaml
COPY config/log.config /data/sonr.chat.log.config

# Expose the necessary ports
EXPOSE 8008
EXPOSE 8448

# Set volumes for media, jetstream, and search index
ENTRYPOINT ["./start.py"]
