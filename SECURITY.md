# Security policy

## Reporting vulnerabilities

Do not disclose suspected vulnerabilities through a public issue. Contact the repository owner privately with reproduction steps, affected versions, and any known mitigations.

## Security boundaries

- The control plane is designed for read-only cloud credentials.
- Production infrastructure changes must traverse the repository's pull-request and approval controls.
- Credentials, customer billing exports, and personally identifiable information must never be committed.
- Synthetic fixtures must be used for the public demonstration environment.

Security support is provided for the latest release only while the project remains pre-1.0.
