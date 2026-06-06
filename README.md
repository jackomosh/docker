# 📟 Containerized ASCII Art Web Server

An isolated multi-stage microservice transforming standard user text streams into geometric terminal character outputs.

## 👥 Authors & Core Contributors

**Jacob Omondi** *Lead Full Stack Developer & Technical Specialist* Lake Victoria Region, Kisumu County, Kenya  

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/ogondajack)
[![Gmail](https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:jackomosh6gmail.com)
[![WhatsApp](https://img.shields.io/badge/WhatsApp-25D366?style=for-the-badge&logo=whatsapp&logoColor=white)](https://wa.me/254703489975)

---

## 🛠️ Commands Reference

### Executing Local Tests
```bash
go test -v ./...

# Build the updated container image
docker build -t ascii-art-web-service:1.0.0 .

# Run the container
docker run -d -p 8080:8080 --name ascii-container ascii-art-web-service:1.0.0