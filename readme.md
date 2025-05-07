#  Todo API

Gin (Go) framework kullanılarak geliştirilmiş bir RESTful yapılacaklar listesi (ToDo) API’sidir. Kullanıcılar JWT tabanlı kimlik doğrulama ile sisteme giriş yapabilir, kendilerine özel yapılacaklar listeleri oluşturabilir ve bu listelerin altında görevlerini (todo item) yönetebilirler.

## Özellikler

- Kullanıcı girişi (JWT ile kimlik doğrulama)
- ToDo List oluşturma, listeleme, güncelleme ve silme
- Her listeye özel ToDo öğesi ekleme, görüntüleme, güncelleme, silme
- Admin ve  kullanıcı rolü 
- Katmanlı mimari (controller, service, model, middleware)
- Gin ile yazılmış sade ve  REST API yapısı

---

##  Proje Yapısı

```bash
todo-case/
│
├── controllers/        # HTTP isteklerini karşılayan kontrolcüler
├── services/           # İş mantığını barındıran servis katmanı
├── models/             # Veri modelleri (User, ToDoList, ToDo)
├── routes/             # Endpoint tanımlamaları
├── utils/              # JWT ve yardımcı fonksiyonlar
├── main.go             # Uygulamanın giriş noktası
└── go.mod              # Go modül yapılandırması
```

## Başlangıç
### Gereksinimler
 - Go 1.23.8
 - Postman veya benzeri bir api test aracı
### Kurulum
- repoyu klonlayın 
```bash
git clone https://github.com/BrkCanbul/todo-case.git
cd todo-case
```
- bağımlılıkları yükleyin
```bash
go mod tidy
```
-  Sunucuyu başlatın
```bash
go run main.go
```

## Endpointler

### POST /login
 * Kullanıcı adı ve şifresi ile giriş yapılır
 * Sunucu kullanıcıya uygun Bearer token döner 
 * Bu aşamadan sonra yapılacak tüm istekler headere
 ```json 
 {"token": "Bearer <<token>>"}
 ```
 şeklinde yapılmalıdır
### GET /todos
 * kullanıcıya ait todo listelerini ve bilgilerini döner

### POST /todos
 * kullanıcının liste oluşturmasını sağlar

### PUT /todos/{id}
 * Kullanıcının listeyi düzenlemesini sağlar
### DELETE /todos/{id}
 * Kullanıcının liste silmesini sağlar
### GET /todos/elems 
 * Kullanıcının bütün todo adımlarını çekmesini sağlar
###  GET /todos/elems/{id}
 * Kullanıcının parametre olarak verilen idli listedeki todo adımlarını çekmesini sağlar
### POST /todos/elems/
 * Kullanıcının todo adımı eklemesini sağlar
 
### PUT /todos/elems/{id}
 * Kullanıcının verilen iddeki todo adımının düzenlenmesini sağlar
 * Bu yapıldığında  elemanın güncelleme tarihi güncellenir
 * liste için tamamlanma yüzdesi kontrolü aşaması gerçekleşir 
 * Eğer yapıldı olarak işaretlenirse bu elemanın bağlı olduğu listenin tamamlanma yüzdesi güncellenir
### DELETE /todos/elems/{id}
 * kullanıcının verilen idli todo adımını silmesini sağlar
