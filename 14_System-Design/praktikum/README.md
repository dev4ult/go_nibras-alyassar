# Latihan - System Design

### Jawaban

1. Gambarkan desain ERD dari sistem pembelian tiket kereta api!
   ![](https://github.com/dev4ult/go_nibras-alyassar/blob/main/14_System-Design/screenshots/ERD.png)

2. Gambarkan use case diagram dari sistem pembelian tiket kereta api!
   ![](https://github.com/dev4ult/go_nibras-alyassar/blob/main/14_System-Design/screenshots/USE%20CASE.png)

3. Terdapat sebuah query pada SQL yaitu SELECT \* FROM users; Dengan tujuan yang sama, tuliskan dalam bentuk perintah:

- Redis : `KEYS *`

- Neo4j :

```
MATCH (u:User)
RETURN u
```

- Cassandra : `SELECT * FROM users;`
