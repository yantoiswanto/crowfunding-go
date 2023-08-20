mysql:
	 docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql:8
createdb:
	    docker exec -it mysql mysql -u root --password=password  -e "CREATE DATABASE db_crowfunding;"
dropdb:
	  docker exec -it mysql mysql -u root -p -e "DROP DATABASE db_crowfunding;"
migrateup:
		migrate -path db/migration -database "mysql://root:password@tcp(localhost:3306)/db_crowfunding?multiStatements=true" -verbose up
migratedown:
			migrate -path db/migration -database "mysql://root:password@tcp(localhost:3306)/db_crowfunding?multiStatements=true" -verbose down

.PHONY: mysql createdb dropdb migrateup migratedown