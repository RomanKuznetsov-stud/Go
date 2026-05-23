package main

/*
double calcTourC(int days, int tickets, int countryIdx, int seasonIdx, int hasGuide, int isLuxury) {
    double pricePerDay = 0.0;

    if (countryIdx == 0) {
        pricePerDay = (seasonIdx == 0) ? 100.0 : 150.0;
    } else if (countryIdx == 1) {
        pricePerDay = (seasonIdx == 0) ? 160.0 : 200.0;
    } else if (countryIdx == 2) {
        pricePerDay = (seasonIdx == 0) ? 120.0 : 180.0;
    }

    double totalBase = pricePerDay * days * tickets;

    double markup = 0.0;
    if (isLuxury) {
        markup = totalBase * 0.20;
    }

    double guideCost = 0.0;
    if (hasGuide) {
        guideCost = 50.0 * days;
    }

    return totalBase + markup + guideCost;
}
*/
import "C"
import (
	"fmt"
	"strconv"

	"github.com/andlabs/ui"
)

func calcTourWrapper(days, tickets int, countryIdx, seasonIdx int, hasGuide, isLuxury bool) float64 {
	cGuide := C.int(0)
	if hasGuide {
		cGuide = 1
	}

	cLuxury := C.int(0)
	if isLuxury {
		cLuxury = 1
	}

	res := C.calcTourC(C.int(days), C.int(tickets), C.int(countryIdx), C.int(seasonIdx), cGuide, cLuxury)
	return float64(res)
}

func initTourGUI() {
	window := ui.NewWindow("Калькулятор туру (CGO)", 400, 350, false)
	window.SetMargined(true)

	daysEntry := ui.NewEntry()
	ticketsEntry := ui.NewEntry()

	countryCombo := ui.NewCombobox()
	countryCombo.Append("Болгарія")
	countryCombo.Append("Німеччина")
	countryCombo.Append("Польща")
	countryCombo.SetSelected(0)

	seasonCombo := ui.NewCombobox()
	seasonCombo.Append("Літо")
	seasonCombo.Append("Зима")
	seasonCombo.SetSelected(0)

	guideCheck := ui.NewCheckbox("Індивідуальний гід ($50/день)")
	luxuryCheck := ui.NewCheckbox("Номер Люкс (+20%)")

	calcBtn := ui.NewButton("Розрахувати тур")
	resultLabel := ui.NewLabel("$0.00")

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	vbox.Append(ui.NewLabel("Кількість днів та кількість путівок:"), false)
	vbox.Append(daysEntry, false)
	vbox.Append(ticketsEntry, false)

	vbox.Append(ui.NewLabel("Країна та Сезон:"), false)
	vbox.Append(countryCombo, false)
	vbox.Append(seasonCombo, false)

	vbox.Append(ui.NewLabel("Додаткові послуги:"), false)
	vbox.Append(guideCheck, false)
	vbox.Append(luxuryCheck, false)

	vbox.Append(calcBtn, false)
	vbox.Append(resultLabel, false)

	window.SetChild(vbox)

	calcBtn.OnClicked(func(*ui.Button) {
		days, _ := strconv.Atoi(daysEntry.Text())
		tickets, _ := strconv.Atoi(ticketsEntry.Text())
		cIdx := countryCombo.Selected()
		sIdx := seasonCombo.Selected()
		hasGuide := guideCheck.Checked()
		isLux := luxuryCheck.Checked()

		total := calcTourWrapper(days, tickets, cIdx, sIdx, hasGuide, isLux)
		resultLabel.SetText(fmt.Sprintf("$%.2f", total))
	})

	window.OnClosing(func(*ui.Window) bool {
		window.Hide()
		return true
	})
	window.Show()
}
