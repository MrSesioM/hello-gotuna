FROM debian:11-slim
ENV PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/go/bin

RUN apt-get update
RUN apt-get install -y git
RUN apt-get install -y wget
RUN wget https://go.dev/dl/go1.20.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz
RUN rm go1.20.linux-amd64.tar.gz
RUN cd opt/
RUN git clone https://github.com/gotuna/gotuna.git

WORKDIR /opt/gotuna
ENV GOROOT=/usr/local/go
RUN go build examples/fullapp/cmd/main.go
EXPOSE 8888
