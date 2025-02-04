# service-account

this repo is used for assessment for hiring purpose

## Installation

use docker to deploy the app

```bash
docker-compose up --build
```
    
## Usage/Examples

there is 4 endpoints in this program. here is the list of samples usage using curl:
- daftar

```
curl --location 'localhost:8080/daftar' \
--header 'Content-Type: application/json' \
--data '{
  "nama": "JoA",
  "nik": "1234567890123451",
  "no_hp": "081234567120"
}'

```

- tabung
```
curl --location 'localhost:8080/tabung' \
--header 'Content-Type: application/json' \
--data '{
  "no_rekening": "68245773340385",
  "nominal": 1
}'
```

- tarik
```
curl --location 'localhost:8080/tarik' \
--header 'Content-Type: application/json' \
--data '{
  "no_rekening": "68245773340385",
  "nominal": 0.99
}'
```

- saldo
```
curl --location 'localhost:8080/saldo/682457733asad40385'
```
