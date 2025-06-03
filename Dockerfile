ARG GOLANG_VERSION=1.24
FROM golang:${GOLANG_VERSION}-alpine as builder

WORKDIR /app

COPY . . 

RUN go build -o terradrift main.go


FROM alpine:latest as runner
ARG TERRAGRUNT_VERSION=0.72.2
ARG TERRAFORM_VERSION=1.7.4

WORKDIR /app

RUN apk add --no-cache \
    curl \
    unzip \
    git

## Terraform
RUN curl -LO https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
    mv terraform /usr/local/bin/ && \
    rm terraform_${TERRAFORM_VERSION}_linux_amd64.zip

## Terragrunt
RUN curl -LO https://github.com/gruntwork-io/terragrunt/releases/download/v${TERRAGRUNT_VERSION}/terragrunt_linux_amd64 && \
    mv terragrunt_linux_amd64 /usr/local/bin/terragrunt && \
    chmod +x /usr/local/bin/terragrunt

COPY --from=builder /app/terradrift .

ENTRYPOINT ["/app/terradrift"] 