# TJ Fleet System

Sistem ini terdiri dari beberapa layanan:

* **fleet-api**: Backend Golang API
* **fleet-ingestion**: Service untuk menerima data lokasi kendaraan via MQTT
* **geofence-worker**: Worker untuk memproses geofence alerts
* **mqtt-publisher**: Service untuk publish data ke MQTT broker
* **PostgreSQL**: Database
* **RabbitMQ**: Message broker
* **Eclipse Mosquitto**: MQTT broker

---

## Prasyarat

* Docker >= 20.x
* Docker Compose >= 2.x
* (Opsional) Postman untuk testing API

---

## Struktur Folder

```text
tj-fleet-system/
├── README.md
├── database
│   ├── migrations
│   │   └── 20251219025117_vehicle_location.up.sql
│   └── readme.txt
├── docker-compose.yml
├── fleet-api
│   ├── Dockerfile
│   ├── application
│   │   └── vehicle_location
│   ├── config
│   │   └── config.go
│   ├── database
│   │   └── database.go
│   ├── domain
│   │   ├── errordata.go
│   │   └── location.go
│   ├── go.mod
│   ├── go.sum
│   ├── helpers
│   │   ├── errorhandler
│   │   ├── handler
│   │   ├── helper
│   │   └── response
│   ├── main.go
│   └── routing
│       ├── init.go
│       ├── route.go
│       └── vehicle_route.go
├── fleet-ingestion
│   ├── Dockerfile
│   ├── application
│   │   ├── rabbitmq
│   │   └── vehicle_locations
│   ├── config
│   │   └── config.go
│   ├── database
│   │   └── database.go
│   ├── domain
│   │   └── location.go
│   ├── go.mod
│   ├── go.sum
│   ├── helper
│   │   ├── consts
│   │   ├── handler
│   │   ├── surounding
│   │   └── validation
│   ├── main.go
│   └── routing
│       ├── init.go
│       └── route.go
├── geofence-worker
│   ├── Dockerfile
│   ├── config
│   │   └── config.go
│   ├── go.mod
│   ├── go.sum
│   ├── helper
│   │   └── surounding
│   └── main.go
├── go.mod
├── go.work
├── go.work.sum
├── mqtt
│   └── config
│       └── mosquitto.conf
├── mqtt-publisher
│   ├── Dockerfile
│   ├── config
│   │   └── config.go
│   ├── domain
│   │   └── mqtt.go
│   ├── go.mod
│   ├── go.sum
│   └── main.go
└── postman
```

---

## Cara Menjalankan

1. **Build dan start semua service**:

```bash
docker compose up --build
```

2. **Hentikan semua service**:

```bash
docker compose down
```

3. **Akses layanan**:

* PostgreSQL: `localhost:5432` (user: `postgres`, password: `postgres`)
* RabbitMQ Management: `http://localhost:15673` (user: `guest`, password: `guest`)
* MQTT Broker: `tcp://localhost:1884`

> Note: Service akan otomatis mengeksekusi migrasi database dari folder `database/migrations` saat container database pertama kali dijalankan.

## Tips & Catatan

* Semua service bisa dijalankan lokal menggunakan Docker Compose tanpa install dependencies manual.
* Migrasi database otomatis dari folder `database/migrations` hanya dijalankan saat container DB pertama kali dibuat.

---

## Testing API (Opsional)

* Gunakan Postman untuk test endpoint API `fleet-api` di `http://localhost:8080`.
