package database

// import GORM dan driver untuk PostgreSQL
import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "log"
)

var DB *gorm.DB  //untuk koneksi ke database

func initializeDB() *gorm.DB {
    //Untuk database, saya menggunakan PostgreSQL lokal
    dsn := "host=localhost user=admin dbname=rakamin sslmode=disable password=rahasia"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Gagal menghubungkan ke database: %v", err)
    }
    
    
    DB = db
    log.Println("Koneksi database berhasil dibuat")
    return DB
}

func getDB() *gorm.DB {
    return DB
}
