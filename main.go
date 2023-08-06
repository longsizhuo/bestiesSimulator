package main

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

//go:embed Noto_Sans_SC/HarmonyOS_Sans_SC_Regular.ttf
var resourceNotoSansSCRegularOtf []byte

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Font(_ fyne.TextStyle) fyne.Resource {
	return &fyne.StaticResource{
		StaticName:    "NotoSansSC-Regular.otf",
		StaticContent: resourceNotoSansSCRegularOtf,
	}
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&myTheme{})
	myWindow := myApp.NewWindow("闺闺模拟器")
	myWindow.Resize(fyne.NewSize(400, 400))
	outputLabel := widget.NewLabel("")

	submitButton := widget.NewButton("输入闺闺名字", func() {
		simulator(outputLabel)
	})

	content := container.NewVBox(
		submitButton,
		outputLabel,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()

}

func simulator(output *widget.Label) {
	for i := 10; i > 0; i-- {
		text := fmt.Sprintf("爱你%%%d\n", i*10)
		output.SetText(output.Text + text)
	}
}
