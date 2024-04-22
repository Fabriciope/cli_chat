# create server image
docker build -f build/server.Dockerfile -t tcp_chat-server:prod .

# create client image
docker build -f build/client.Dockerfile -t tcp_chat-client:prod .
