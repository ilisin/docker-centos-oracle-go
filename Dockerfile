FROM 192.168.10.40:5000/linux-oracle

MAINTAINER lisin gao <ilisin.gao@gmail.com>


COPY pkg-config /opt/pkg-config
COPY go /opt/go

RUN yum install subversion -y
RUN yum install gcc -y
RUN yum install git -y
RUN yum install zip -y
RUN yum install tar -y


RUN mkdir -p /workplace/go/src

ENV GOROOT /opt/go
ENV GOARCH amd64
ENV GOOS linux
ENV GOPATH /workplace/go
ENV CGO_CFLAGS -I/usr/include/oracle/11.2/client64
ENV CGO_LDFLAGS -L/usr/lib/oracle/11.2/client64/lib
ENV LD_LIBRARY_PATH /usr/lib/oracle/11.2/client64/lib:/lib:/usr/lib:/usr/local/lib
ENV PKG_CONFIG_PATH /opt/pkg-config
ENV PATH $PATH:/opt/go/bin

VOLUME /opt/share
EXPOSE 10000 10001 10002 10003 10004 10005

CMD ["/bin/sh"]
