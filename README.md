CRUD ke-6

Customer

POST
curl -X POST http://localhost:8080/customers \
  -H "Content-Type: application/json" \
  -d '{"name":"Ahmad","email":"ahmad@example.com","phone":"08123456789"}'

GET
curl http://localhost:8080/customers

curl http://localhost:8080/customers/1

PUT 
curl -X PUT http://localhost:8080/customers/1 \
-H "Content-Type: application/json" \
-d '{
  "name": "Ahmad Updated",
  "email": "ahmad.updated@example.com",
  "phone": "081234567899"
}'

DELETE
curl -X DELETE http://localhost:8080/customers/1