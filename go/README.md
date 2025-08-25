# GitHub Webhook Handler

A secure Go HTTP server that handles GitHub webhook events with HMAC-SHA256 signature validation.

## Features

- **Secure webhook validation** using HMAC-SHA256 signatures
- **Event handling** for push, pull request, and issues events
- **Extensible design** for adding new event handlers
- **Health check endpoint** for monitoring
- **Environment-based configuration**

## Setup

1. Set your webhook secret:
   ```bash
   export GITHUB_WEBHOOK_SECRET="your_secret_here"
   ```

2. Optionally set the port (defaults to 8080):
   ```bash
   export PORT="8080"
   ```

3. Run the server:
   ```bash
   go run main.go
   ```

## GitHub Webhook Configuration

1. In your GitHub repository, go to Settings → Webhooks
2. Add a new webhook with:
   - **Payload URL**: `http://your-server:8080/webhook`
   - **Content type**: `application/json`
   - **Secret**: Same value as your `GITHUB_WEBHOOK_SECRET`
   - **Events**: Select the events you want to handle

## Endpoints

- `POST /webhook` - GitHub webhook endpoint
- `GET /health` - Health check endpoint

## Security

The handler validates all incoming webhooks using HMAC-SHA256 signatures to ensure they come from GitHub and haven't been tampered with.