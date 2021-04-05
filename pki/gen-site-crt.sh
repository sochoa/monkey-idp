#!/bin/bash -xe

TMPDIR="tmp"
SITE="${1:-monkey-idp.local}"
SITE_KEY="$TMPDIR/$SITE.key"
SITE_CSR="$TMPDIR/$SITE.csr"
SITE_CRT="./static/$SITE.crt"
CA_CRT="./static/ca.crt"
CA_KEY="$(ls $TMPDIR/root-ca*.key)"
PASSPHRASE_FILE="$(ls $TMPDIR/root-ca-passphrase-*.txt)"
SITE_CRT_EXTENSIONS="$TMPDIR/$SITE.openssl-conf"

# Generate a private key for the sites
openssl genrsa -out $SITE_KEY 2048

# Generate a CSR based on the site's private key
openssl req                 \
  -new                      \
  -key $SITE_KEY            \
  -subj "/CN=$SITE/O=$SITE" \
  -out $SITE_CSR

# Create extensions file
cat ./pki/site-crt.template | sed "s/{{SITE}}/$SITE/g" > "$SITE_CRT_EXTENSIONS"

# Finally, generate the site cert
openssl x509 -req                        \
  -passin         file:$PASSPHRASE_FILE  \
  -in             $SITE_CSR              \
  -CA             $CA_CRT                \
  -CAkey          $CA_KEY                \
  -CAcreateserial                        \
  -out            $SITE_CRT              \
  -days 365                              \
  -sha256                                \
  -extfile        "$SITE_CRT_EXTENSIONS"
