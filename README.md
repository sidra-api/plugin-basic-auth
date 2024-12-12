# Plugin Basic Auth

## Deskripsi
Plugin Basic Auth digunakan untuk memvalidasi autentikasi dasar (Basic Authentication) pada Sidra Api. Plugin ini memastikan hanya pengguna yang memiliki kredensial yang valid dapat mengakses layanan backend.

---

## Cara Kerja
1. **Validasi Authorization Header**
   - Plugin memeriksa header `Authorization` untuk memastikan memiliki nilai:
     ```
     Basic <base64(username:password)>
     ```
   - Kredensial default yang digunakan adalah:
     - Username: `foo`
     - Password: `bar`

2. **Respon**
   - Jika kredensial valid:
     - Status: `200 OK`
     - Body: "Hello, World!"
   - Jika kredensial tidak valid:
     - Status: `401 Unauthorized`
     - Body: "Unauthorized"

---

## Konfigurasi
- **Kredensial**
  - Kredensial dapat dikonfigurasi langsung pada file `main.go`:
    ```go
    username := "foo"
    password := "bar"
    ```

---

## Cara Menjalankan
1. Pastikan Anda sudah menginstal **Sidra Api**.
2. Tambahkan plugin ini ke direktori `plugins/basic-auth/main.go` pada Sidra Api.
3. Kompilasi dan jalankan Sidra Api.
4. Plugin akan otomatis terhubung melalui UNIX socket pada path `/tmp/basic-auth.sock`.

---

## Pengujian

### Endpoint
- **URL**: Endpoint mana saja yang dikonfigurasi untuk melewati plugin Basic Auth.

### Langkah Pengujian
1. Kirim request dengan header `Authorization` yang valid:
   ```
   Authorization: Basic <base64(username:password)>
   ```
2. Respons yang diharapkan:
   - Jika kredensial valid:
     - Status: `200 OK`
     - Body: "Hello, World!"
   - Jika kredensial tidak valid:
     - Status: `401 Unauthorized`
     - Body: "Unauthorized"

---

## Catatan Penting
- **Keamanan**: Jangan menyimpan kredensial dalam bentuk hard-coded di produksi. Gunakan metode yang lebih aman seperti variabel lingkungan.
- **Header Authorization**: Pastikan klien Anda mengirimkan header `Authorization` dengan format yang benar.

---

## Lisensi
MIT License
