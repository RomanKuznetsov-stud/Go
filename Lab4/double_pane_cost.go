package main

import (
	"fmt"
	"strconv"

	"github.com/andlabs/ui"
)

func calcWindowGo(width, height float64, materialIdx, chamberIdx int, hasSill bool) float64 {
	area := width * height
	var pricePerCm2 float64

	if materialIdx == 0 {
		if chamberIdx == 0 {
			pricePerCm2 = 2.5
		} else {
			pricePerCm2 = 3.0
		}
	} else if materialIdx == 1 {
		if chamberIdx == 0 {
			pricePerCm2 = 0.5
		} else {
			pricePerCm2 = 1.0
		}
	} else if materialIdx == 2 {
		if chamberIdx == 0 {
			pricePerCm2 = 1.5
		} else {
			pricePerCm2 = 2.0
		}
	}

	total := area * pricePerCm2
	if hasSill {
		total += 350.0
	}
	return total
}

func initWindowGUI() {
	window := ui.NewWindow("Калькулятор склопакета", 400, 250, false)
	window.SetMargined(true)

	widthEntry := ui.NewEntry()
	heightEntry := ui.NewEntry()

	materialCombo := ui.NewCombobox()
	materialCombo.Append("Дерево")
	materialCombo.Append("Метал")
	materialCombo.Append("Металопластик")
	materialCombo.SetSelected(0)

	chamberCombo := ui.NewCombobox()
	chamberCombo.Append("Однокамерний")
	chamberCombo.Append("Двокамерний")
	chamberCombo.SetSelected(0)

	sillCheck := ui.NewCheckbox("Підвіконня")

	calcBtn := ui.NewButton("Розрахувати")
	resultLabel := ui.NewLabel("0.00 грн")

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(ui.NewLabel("Розмір вікна: Ширина (см) та Висота (см)"), false)
	vbox.Append(widthEntry, false)
	vbox.Append(heightEntry, false)

	vbox.Append(ui.NewLabel("Матеріал та Склопакет"), false)
	vbox.Append(materialCombo, false)
	vbox.Append(chamberCombo, false)
	vbox.Append(sillCheck, false)

	vbox.Append(calcBtn, false)
	vbox.Append(resultLabel, false)

	window.SetChild(vbox)

	calcBtn.OnClicked(func(*ui.Button) {
		w, _ := strconv.ParseFloat(widthEntry.Text(), 64)
		h, _ := strconv.ParseFloat(heightEntry.Text(), 64)
		matIdx := materialCombo.Selected()
		chamIdx := chamberCombo.Selected()
		hasSill := sillCheck.Checked()

		total := calcWindowGo(w, h, matIdx, chamIdx, hasSill)
		resultLabel.SetText(fmt.Sprintf("%.2f грн", total))
	})

	// Ховаємо вікно при закритті замість виходу з програми
	window.OnClosing(func(*ui.Window) bool {
		window.Hide()
		return true
	})
	window.Show()
}
