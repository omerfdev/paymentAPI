# Ödeme Servisi

Bu basit Go programı, bir ödeme servisi oluşturur ve yeni ödemeler oluşturmak ve ödemeleri almak için HTTP API sağlar.

## Nasıl Çalışır

1. **Kurulum**: Projeyi klonlayın ve Go'nun yüklü olduğundan emin olun.
2. **Bağımlılıklar**: `go get` komutuyla `github.com/gorilla/mux` ve `github.com/rs/cors` paketlerini indirin.
3. **Servisi Başlatın**: `go run main.go` komutuyla servisi başlatın.
4. **Ödemeleri Yönetin**: Tarayıcıdan veya bir API test aracından `/payments` endpoint'ine bir POST isteği yaparak yeni ödemeler oluşturun veya `/payments/{id}` endpoint'ine bir GET isteği yaparak bir ödemenin detaylarını alın.

## API Dökümantasyonu

### `/payments` Endpoint'i

- **Method**: POST
- **Cevap**: Yeni bir ödemenin oluşturulduğuna dair bir yanıt döndürür.

Örnek istek:
POST http://localhost:8080/payments
GET http://localhost:8080/payments/123
