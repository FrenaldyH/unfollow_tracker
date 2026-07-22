// internal/gui/app.go
package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"unfollow_tracker/internal/parser"
	"unfollow_tracker/internal/service"
	"unfollow_tracker/pkg/logger"
)

// defaultPath adalah path default folder data export Instagram,
// sama dengan konstanta PATH yang sebelumnya ada di main.go.
const defaultPath = "[INSTAGRAM JSON FOLDER NAME]"

// Run menjalankan aplikasi GUI Unfollow Tracker.
func Run() {
	a := app.New()
	w := a.NewWindow("Unfollow Tracker")
	w.Resize(fyne.NewSize(500, 600))

	// Input path folder data
	pathEntry := widget.NewEntry()
	pathEntry.SetText(defaultPath)

	// Data hasil, disimpan di sini supaya bisa diakses oleh list widget
	var resultData []string

	// List widget untuk menampilkan hasil
	resultList := widget.NewList(
		func() int { return len(resultData) },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(resultData[i])
		},
	)

	statusLabel := widget.NewLabel("Siap.")

	// Tombol untuk pilih folder lewat dialog (opsional, biar tidak ketik manual)
	pickFolderButton := widget.NewButton("Pilih Folder", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil || uri == nil {
				return
			}
			pathEntry.SetText(uri.Path())
		}, w)
	})

	// Tombol utama: jalankan proses parsing + analisis
	checkButton := widget.NewButton("Cek Unfollowed", func() {
		path := pathEntry.Text
		statusLabel.SetText("Memproses...")

		following, err := parser.FollowingParser(path)
		if err != nil {
			logger.Log.Error("gagal parsing following", "error", err, "path", path)
			dialog.ShowError(fmt.Errorf("gagal membaca following: %w", err), w)
			statusLabel.SetText("Gagal.")
			return
		}

		followers, err := parser.FollowersParser(path)
		if err != nil {
			logger.Log.Error("gagal parsing followers", "error", err, "path", path)
			dialog.ShowError(fmt.Errorf("gagal membaca followers: %w", err), w)
			statusLabel.SetText("Gagal.")
			return
		}

		resultData = service.UnfollowedDisplay(followers, following)
		resultList.Refresh()

		statusLabel.SetText(fmt.Sprintf("Ditemukan %d akun yang tidak follow balik.", len(resultData)))
		logger.Log.Info("proses selesai", "total", len(resultData))
	})

	topBar := container.NewBorder(nil, nil, nil, pickFolderButton, pathEntry)

	content := container.NewBorder(
		container.NewVBox(topBar, checkButton, statusLabel), // atas
		nil, nil, nil,
		container.NewVScroll(resultList), // tengah, mengisi sisa ruang, scrollable
	)

	w.SetContent(content)
	w.ShowAndRun()
}
