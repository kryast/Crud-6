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


Product

POST
curl -X POST http://localhost:8080/products \
-H "Content-Type: application/json" \
-d '{
  "name": "Monitor LED 24 Inch",
  "category": "Elektronik",
  "price": 1500000
}'

GET
curl http://localhost:8080/products

curl http://localhost:8080/products/1

PUT
curl -X PUT http://localhost:8080/products/1 \
-H "Content-Type: application/json" \
-d '{
  "name": "Monitor LED 27 Inch",
  "category": "Elektronik",
  "price": 1800000
}'

DELETE
curl -X DELETE http://localhost:8080/products/1


Order

POST
curl -X POST http://localhost:8080/orders \
-H "Content-Type: application/json" \
-d '{
  "customer_id": 1,
  "status": "pending"
}'

GET
curl http://localhost:8080/orders

curl http://localhost:8080/orders/1

PUT
curl -X PUT http://localhost:8080/orders/1 \
-H "Content-Type: application/json" \
-d '{
  "customer_id": 1,
  "status": "paid"
}'