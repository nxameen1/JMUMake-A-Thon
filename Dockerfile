FROM golang:1.24.2
WORKDIR /app
COPY . .
RUN cd serveathon && go build -o /app/server .
CMD ["/app/server"]
