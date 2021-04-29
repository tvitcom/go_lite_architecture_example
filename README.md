The lite architecture

This structure the result of my impression reading the "Organising Database Access in Go" articles https://www.alexedwards.net/blog/organising-database-access

You can use it example for yourself, imported database example data/bookstore_start.sql into mysql database bookstore server and edit interanl/config.go: STORAGE_DSN 
then run as:

```bash
./run.sh
```

Web server should have available at url http://localhost:3000/books 