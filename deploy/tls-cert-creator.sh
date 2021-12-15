#!/bin/sh

openssl req -x509 \
  -nodes \
  -newkey rsa:4096 \
  -days 3650 \
  -keyout ca.key \
  -out ca.crt \
  -subj "/CN=*.default.svc" \
  -addext "subjectAltName=DNS:kube-admission.default.svc,DNS:*.default.svc"

