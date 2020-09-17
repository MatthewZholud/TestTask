FROM golang:latest

WORKDIR app/

COPY CompanyManager .

CMD ["go","run","main.go"]