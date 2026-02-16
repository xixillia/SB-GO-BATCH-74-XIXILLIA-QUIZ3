# Quiz 3 Golang

REST API menggunakan gin framework dan di-deploy ke railway app untuk kelola data kategori dan buku.
 
##  Endpoint API

### Categories
| Method | Endpoint | Deskripsi |
| :--- | :--- | :--- |
| `GET` | `/api/categories` | List semua kategori |
| `GET` | `/api/categories/:id` | Detail kategori berdasarkan ID |
| `POST` | `/api/categories` | Tambah kategori baru |
| `PUT` | `/api/categories/:id` | Update data kategori |
| `DELETE` | `/api/categories/:id` | Hapus kategori |
| `GET` | `/api/categories/:id/books` | List semua buku berdasarkan kategori tertentu |

### 📚 Books
| Method | Endpoint | Deskripsi |
| :--- | :--- | :--- |
| `GET` | `/api/books` | List semua buku |
| `GET` | `/api/books/:id` | Detail buku berdasarkan ID |
| `POST` | `/api/books` | Tambah buku baru |
| `PUT` | `/api/books/:id` | Update data buku |
| `DELETE` | `/api/books/:id` | Hapus buku |
