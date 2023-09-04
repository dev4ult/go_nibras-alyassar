# Concurrent Programming

### Poin poin pembelajaran

1.  **Sequential Program** merupakan program yang berjalan secara beurutan dengan setiap pekerjaan harus diselesaikan terlebih dahulu sebelum memulai pekerjaan lain.
2. **Parralell Program** merupakan program yang berjalan secara *parallel* atau dibagi pekerjaan nya dengan pekerja lainnya (tidak dikerjakan oleh satu pekerja). Pekerja yang dimaksud didalam komputer adalah *Thread* CPU.
3. **Concurrent Program** merupakan program yang berjalan secara bergantian antara pekerjaanya walau belum selesai diantara pekerjaannya untuk mengoptimalkan waktu pekerjaan dengan satu pekerja (*thread* CPU melakukan *multitasking* dengan *switching task*).

---

### Keyword yang digunakan

1. **Goroutine** merupakan fungsi yang berjalan secara *concurrently* atau *independently* bersama dengan fungsi yang lainnya.
2. **GOMAXPROCS** digunakan untuk mengatur berapa banyak *threads* nya sistem operasi.
3. **Channel** merupakan objek komunikasi yang digunakan goroutine untuk berkomunikasi dengan yang lainnya.
