# Rate Limiter

An effective way to safeguard your server against attacks.

## Rules 
- This rate limiter supports two configuration methods: IP-based and token-based.  
- When a token is included in the request, it must already be configured; otherwise, the request will return an error.  
- Token-specific settings take precedence over IP-based settings.  
- Tokens must be included in the request header in the format: `API_KEY=tokenname`.

## Configuration Instructions
All configurations should be added to the `.env` file located in the root of the project, using the following format:

### IP-Based Configuration
```plaintext
RATE_LIMITER_BY_IP_LIMIT=5
RATE_LIMITER_BY_IP_WINDOW=10
RATE_LIMITER_BY_IP_BLOCK_WINDOW=5
```

### Token Based Configuration
```plaintext
RATE_LIMITER_BY_TOKEN_tokenname_LIMIT=3
RATE_LIMITER_BY_TOKEN_tokenname_WINDOW=5
RATE_LIMITER_BY_TOKEN_tokenname_BLOCK_WINDOW=10
```

_**You can determine the name of the token and configure as many as you want**_

## How to run
```bash
go mod tidy

docker-compose up -d
```

The server will be running on port `:8080`

## Tests
### Automated tests
```bash
go test ./...   
```

### Testing the Rate Limiter by IP
```bash
bash bash/test_by_ip.bash   
```

### Testing the Rate Limiter by Token
```bash
bash bash/test_by_token.bash 
```