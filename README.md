# GO-KAFKA-PLAYGROUND
The project is example for using kafka with microservice 

### Package that used in project 
*  Webservice
    *  [echo](https://github.com/labstack/echo)
*  Kafka
    *  [sarama](https://github.com/Shopify/sarama)  
*  Postgres ORM
    *  [go-pg](https://github.com/go-pg/pg)
*  Mongo ORM
    *  [mgo](github.com/globalsign/mgo)
*  Migration
    *  [golang-migrate](https://github.com/golang-migrate/migrate)

### Installing 

1.  **Create .env and .kafka.env file** <br>
`.env` is using config for application <br> 
`.kafka.env` is using config host ip for build kafka on docker <br> 
so You can create .env and .kafka.env By Command
```
cp .env.example .env
touch .kafka.env
```
Explaning .env file 
```
PORT=3000                   //APP PORT
PSQL_DATABASE_URL=postgres://postgres:postgres@psql_db:5432/app_example?sslmode=disable  //DATABASE POSTGRES CONNECTION STRING
MONGO_DATABASE_URL=mongodb://mongoadmin:mongoadmin@mongo_db:27017/app_example //DATABASE Mongo CONNECTION STRING

// kafka 
PRODUCER_URL=10.10.8.16:9092,10.10.8.16:9093  // ${HOST_IP:PORT}
CONSUMER_URL=10.10.8.16:9094,10.10.8.16:9095  // ${HOST_IP:PORT}
```
2.  **Find Host IP to kafka docker** <br>
Kafka image on docker must use the host IP and write it to file `.kafka.env` that have variable `MACHINE_HOST=${HOST_IP}`
    * Window
        - The docker on Window already have host ip from docker-machine you can find host ip from docker-machine
    * Mac OS 
        - The docker on Mac OS and Window are different. Mac OS using the native docker cli that doesn't have docker-machine to build docker host ip to image but we can find host ip from ethernet card `en0`
        - We write the command to get host ip and write file .kafa.env by this command 
        ```
            make kafka.env
        ```
3.  **Running app**  <br>
Now we can running application with docker compose by this command
```
   docker-compose up
```
4. **Migrate Database** <br>
In this project we use postgres and mongodb. the postgres image will create postgres database and uuid-ossp extension only so we must manual migrate by this command 
    -  Downloading migration to app 
    ```
        make install-migration
    ```     
    -  Migrate table to app with parameter `${POSTGRES_CONNECTIONSTRING}`
    ```
        make app.migration.up db_url="${POSTGRES_CONNECTIONSTRING}"
    ```

5. **Playing**  <br>
    Open Browser and route to `http://127.0.0.1/users` This route will create user and insert to postgres and invoke kafka message to create in mongodb

