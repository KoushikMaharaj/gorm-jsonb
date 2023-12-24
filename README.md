# gorm-jsonb

[gorm](https://gorm.io) jsonb datatype

In this package you can store json as well as json array in database
****************************************************************

# install

```
go get github.com/KoushikMaharaj/gorm-jsonb
```

****************************************************************

# Creation of table

##### postgres

```sql
CREATE TABLE public."data"(
    info    jsonb NULL,
    address jsonb NULL
);
```

##### mysql

```sql
CREATE TABLE `data`(
    `info`    json DEFAULT NULL,
    `address` json DEFAULT NULL
);
```

****************************************************************

# example

Please refer [example](https://github.com/KoushikMaharaj/gorm-jsonb/blob/master/exmaple/main.go)
****************************************************************

# Reference

[dariubs](https://github.com/dariubs/gorm-jsonb)

