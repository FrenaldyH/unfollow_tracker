// Package logger menyediakan logger terstruktur (structured logging)
// berbasis slog yang menulis output ke stdout dan ke file harian
// di dalam folder "logs/".
package logger

import (
	"io"
	"log"
	"log/slog"
	"os"
	"time"
)

// Log adalah instance slog.Logger global yang digunakan di seluruh
// aplikasi untuk mencatat log. Harus diinisialisasi terlebih dahulu
// dengan memanggil Init() sebelum digunakan.
var Log *slog.Logger

// logFile menyimpan referensi ke file log yang sedang aktif,
// digunakan oleh Close() untuk menutup file dengan rapi saat
// aplikasi berhenti.
var logFile *os.File

// Init menginisialisasi logger global Log.
//
// Fungsi ini melakukan beberapa hal:
//   - Membuat folder "logs/" jika belum ada.
//   - Membuat (atau membuka) file log dengan nama berdasarkan
//     tanggal saat ini, contoh: "logs/2026-07-23.log".
//   - Mengatur output logger ke stdout dan file log secara bersamaan
//     menggunakan io.MultiWriter.
//   - Menggunakan format JSON (slog.NewJSONHandler) dengan level
//     minimum Debug dan menyertakan informasi source code
//     (file & baris) di setiap entri log.
//
// Init akan memanggil log.Fatalf (yang menghentikan aplikasi dengan
// os.Exit) jika folder atau file log gagal dibuat.
//
// Setelah memanggil Init, sebaiknya panggil Close melalui defer
// di main() agar file log ditutup dengan rapi saat aplikasi berhenti.
//
// Contoh pemakaian:
//
//	func main() {
//	    logger.Init()
//	    defer logger.Close()
//	    logger.Log.Info("aplikasi dimulai")
//	}
func Init() {
	// Pastikan folder "logs/" tersedia untuk menyimpan file log.
	err := os.MkdirAll("logs", 0o755)
	if err != nil {
		log.Fatalf("Failed to create logs directory: %v", err)
	}

	// Nama file log berdasarkan tanggal hari ini, contoh: logs/2026-07-23.log
	filename := "logs/" + time.Now().Format("2006-01-02") + ".log"

	// Buka file log dalam mode append, buat baru jika belum ada.
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	logFile = file

	// Tulis log ke stdout (untuk development/monitoring langsung)
	// dan ke file (untuk penyimpanan permanen) secara bersamaan.
	multi := io.MultiWriter(os.Stdout, file)

	// Inisialisasi logger dengan format JSON, level minimum Debug,
	// dan menyertakan lokasi kode (file & baris) di setiap log.
	Log = slog.New(slog.NewJSONHandler(multi, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}))
}

// Close menutup file log yang sedang aktif.
//
// Fungsi ini sebaiknya dipanggil melalui defer segera setelah
// Init() dipanggil di main(), untuk memastikan semua data log
// ter-flush ke disk dengan rapi saat aplikasi berhenti.
//
// Close aman dipanggil meski Init() belum pernah dipanggil
// (logFile bernilai nil), dalam hal ini Close tidak melakukan apa-apa.
func Close() error {
	if logFile == nil {
		return nil
	}
	return logFile.Close()
}
