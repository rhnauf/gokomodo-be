# Gokomodo BE Assessment

the api request & response is exported as a *Postman* collection in JSON format

---
create env files & modify to your own configurations

```cp .env.example .env```

run docker for postgres instance

```docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:15-alpine```

run migrations, using your own preferred tools (example using goose)

```cd ./sql/schema && goose postgres postgres://root:root@localhost:5432/dbname up```

to seed the ```users``` & ```buyers``` table, uncomment ```seedBuyersTable``` & ```seedSellersTable```  function on ```main.go``` files

to run

```make run``` or ```go run ./cmd/app/```

alternatively if you prefer to build binary and run

```make build-run``` or ```go build -C cmd/app -o ../../gokomodo-be.exe && ./gokomodo-be.exe```