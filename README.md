# GuacAlert

GuacAlert is a Go application that monitors active connections in an Apache Guacamole server and sends alerts to a Webhook when users connect or disconnect.

## Features

- Monitors active connections in real-time
- Sends alerts when users connect or disconnect
- Handles token generation and expiration
- Docker support for easy deployment

## Environment Variables

The application requires the following environment variables:

- `GUAC_URL`: The URL of your Guacamole server
- `GUAC_USER`: The username for the Guacamole API
- `GUAC_PASS`: The password for the Guacamole API
- `GUAC_DATASOURCE`: The data source for the Guacamole API
- `CHAT_HOOK`: The webhook URL for sending chat notifications

## Docker

The latest image can be found here:

```docker push jandersonai/guacalert:latest```

To run the Docker container, use the following command:

```bash
docker run -e "GUAC_URL=https://example.com/guacamole" -e "GUAC_USER=username" -e "GUAC_PASS=password" -e "GUAC_DATASOURCE=datasource" -e "CHAT_HOOK=webhook_url" guacalert
```

Replace the placeholders with your actual values.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
