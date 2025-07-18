CRUD ke-6

POST Customer

curl -X POST http://localhost:8080/customers \
  -H "Content-Type: application/json" \
  -d '{"name":"Ahmad","email":"ahmad@example.com","phone":"08123456789"}'