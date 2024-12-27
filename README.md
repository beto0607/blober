# Blober

Blober, your own blob storage

## Development

### Requires

- MongoDB
- Root folder to save the files with `0700` permissions

### Configuration

Duplicate and rename the `.env.dev` file to `.env`.

#### Default configuration

- DB server: `mongodb://localhost:27017``
- Root folder: `/var/tmp/blober/` (will be automatically generated)
- Port: `8978``
- Host: `api.blober.local`
