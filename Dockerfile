FROM golang:latest

COPY run.sh /app/run.sh

RUN curl -sL https://deb.nodesource.com/setup_12.x | bash - && apt install nodejs \
    && mkdir buffalo && cd buffalo \
    && wget https://github.com/gobuffalo/buffalo/releases/download/v0.14.10/buffalo_0.14.10_linux_amd64.tar.gz \
    && tar -xvzf buffalo_0.14.10_linux_amd64.tar.gz \
    && mv buffalo /usr/local/bin/buffalo \
    && cd .. && rm -rf buffalo \
    && mkdir -p $GOPATH/src/github.com \
    && git config --global user.email "me@example.com" \
    && git config --global user.name "My Name" \
    && chmod +x /app/run.sh \
    && cd $GOPATH/src/github.com \
    && git clone https://github.com/gobuffalo/authrecipe.git \
    && buffalo new myapp

#COPY api/* $GOPATH/src/github.com/myapp/

COPY database.yml $GOPATH/src/github.com/authrecipe/database.yml

EXPOSE 3000

CMD ["/bin/bash", "/app/run.sh"]  
