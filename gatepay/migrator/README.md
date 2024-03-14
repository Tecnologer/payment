# Migrator

This is a simple tool that creates the tables and seed the database with some data.

## Usage
```text
  -db-host string
    	Database host (default "localhost")
  -db-name string
    	Database name (default "gatepay")
  -db-pass string
    	Database password (default "S3cret*_2024")
  -db-port int
    	Database port (default 5432)
  -db-user string
    	Database user (default "postgres")
```

## Execute

```shell
   go run main.go -db-host localhost \
                  -db-name gatepay \
                  -db-pass S3cret*_2024 \
                  -db-port 5432 \
                  -db-user postgres 
```

## Update seeder 

If you want to update the seeder, you can modify the `seeder/seeder.go` file and run the migrator again.