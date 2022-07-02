0. `sudo apt install postgresql postgresql-contrib`
1. `su - postgres`
2. `createuser --interactive --pwprompt`
3. `createdb -h localhost -p 5432 -U [username] [dbname]`
4. `psql -f schema.sql -h 127.0.0.1 -p 5432 -U [username] [dbname]`
