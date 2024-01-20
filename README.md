## xk6-output-gcs

`xk6-output-gcs` is a k6 output extension that uploads test results to Google Cloud Storage.

## Install

```bash
# Install xk6
go install go.k6.io/xk6/cmd/xk6@latest

# Build the k6 binary
xk6 build --with github.com/churark/xk6-output-gcs=.
```

You will have a k6 binary in the current directory.

## Usage

TODO

## Configuration

Configurations are set by environment variables.

| Key            | Type   | Default |          | 
| -------------- | ------ | ------- | -------- | 
| GCS_PROJECT_ID | string |         | required | 
| GCS_BUCKET     | string |         | required | 

## Contribution

Contributions of any kind welcome!

## LICENSE

[MIT LICENSE](https://github.com/churark/xk6-output-gcs/blob/main/LICENSE)
