package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func setupUI() {
	// Створюємо головне вікно-меню
	mainWindow := ui.NewWindow("Лабораторна робота №4", 300, 150, false)
	mainWindow.SetMargined(true)

	btnTask1 := ui.NewButton("Завдання 1: Калькулятор склопакета")
	btnTask2 := ui.NewButton("Завдання 2: Калькулятор туру")

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	vbox.Append(btnTask1, false)
	vbox.Append(btnTask2, false)

	mainWindow.SetChild(vbox)

	// Прив'язуємо кнопки до функцій з інших файлів
	btnTask1.OnClicked(func(*ui.Button) {
		initWindowGUI()
	})

	btnTask2.OnClicked(func(*ui.Button) {
		initTourGUI()
	})

	// Закриття головного вікна завершує всю програму
	mainWindow.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	mainWindow.Show()
}

func main() {
	err := ui.Main(setupUI)
	if err != nil {
		panic(err)
	}
}
