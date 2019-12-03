#!/bin/bash

set -o errexit

main() {
  echo "Connect DATABASE"
  create_user

  echo "CREATING DATABASE"
  create_databases
}

create_user() {
  mongo --port 27017 <<EOF
     use admin;
     db.createUser(
        {
          user: "mongo",
          pwd: "mongo",
           roles: [ 
              { role: "userAdminAnyDatabase", db: "admin" }, 
              { role: "dbAdminAnyDatabase", db: "admin" }, 
              { role: "readWriteAnyDatabase", db: "admin" } 
            ]
        }
     );
EOF
}

create_databases() {
  mongo --port 27017 -u amic -p amic --authenticationDatabase admin <<EOF
      use app_example;
      db.createUser(
        {
          user: "mongo",
          pwd: "mongo",
           roles: [ 
              { role: "readWrite", db: "app_example" } 
            ]
        }
      );
EOF
}

main "$@"