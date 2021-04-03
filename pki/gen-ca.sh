#!/bin/bash -xe
NOW=$(date +"%Y%m%s%H%M%S")
TMPDIR="tmp"
PASSPHRASE_FILE="$TMPDIR/root-ca-passphrase-$NOW.txt"
ROOT_CA_KEY="$TMPDIR/root-ca-${NOW}.key"
ROOT_CERT="$TMPDIR/root-${NOW}.crt"
openssl rand -base64 64 | tee $PASSPHRASE_FILE &>/dev/null
chmod 0400 $PASSPHRASE_FILE

# Generate the root certificate authority key
openssl genrsa -aes256 -passout file:$PASSPHRASE_FILE -out $ROOT_CA_KEY 4096

# Generate a root cert
openssl req \
  -x509 \
  -passin file:$PASSPHRASE_FILE \
  -new \
  -nodes \
  -key $ROOT_CA_KEY \
  -sha256 \
  -days 365 \
  -subj '/CN=ca.local/O=Local Certifiate Authority/C=US' \
  -out static/$ROOT_CERT

(cd static; ln -sf $ROOT_CERT ca.crt)
