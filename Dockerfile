FROM gobuffalo/buffalo:latest as builder

COPY run.sh /app/run.sh

#COPY api $GOPATH/src/github.com/api

RUN chmod +x /app/run.sh && cd $GOPATH/src/github.com/ && buffalo new coke

WORKDIR /app

EXPOSE 3000

CMD ["/bin/bash", "run.sh"]  
