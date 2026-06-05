# 📟 Containerized ASCII Art Web Server

An isolated multi-stage microservice transforming standard user text streams into geometric terminal character outputs.

## 🛠️ Commands Reference

### Executing Local Tests
```bash
go test -v ./...

# Run unit tests locally to verify file access matching
go test -v ./...

# Build the updated container image
docker build -t ascii-art-web-service:1.0.0 .

# Run the container
docker run -d -p 8080:8080 --name ascii-container ascii-art-web-service:1.0.0