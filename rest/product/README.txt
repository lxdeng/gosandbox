Need to create the table first. One way is to run the "go test -v" that creates the table.

list products:
curl http://localhost:8080/products

insert a product:
curl -H "Content-Type: application/json" -X POST -d '{"name":"desk","price":50.99}' http://localhost:8080/product
curl -H "Content-Type: application/json" -X POST -d '{"name":"chair","price":20.99}' http://localhost:8080/product

delete a product:
curl -X DELETE http://localhost:8080/product/<id>

get a product:
curl http://localhost:8080/product/<id>

update a product:
curl -H "Content-Type: application/json" -X PUT -d '{"id": 3, "name":"desk","price":55.99}'  http://localhost:8080/product/3
