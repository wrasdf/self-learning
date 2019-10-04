FROM centos:centos7
RUN yum update -y
RUN yum -y install httpd postgresql-server postgresql-contrib

RUN mkdir -p /var/www/selflearn.com/ && \
    mkdir -p /var/www/selflearn.com/dist/

COPY conf/httpd.conf /etc/httpd/conf/
COPY dist/* /var/www/selflearn.com/dist/

ENV PGWEB_VERSION "0.11.3"
RUN yum -y install unzip && \
    curl -L https://github.com/sosedoff/pgweb/releases/download/v${PGWEB_VERSION}/pgweb_linux_amd64.zip > /tmp/pgweb.zip && \
    unzip /tmp/pgweb.zip -d /app && yum -y remove unzip && yum clean all && rm -rf /tmp/pgweb.zip

RUN postgresql-setup initdb && systemctl start postgresql

EXPOSE 80
CMD /app/bin/run.sh
