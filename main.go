package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type myTheme struct {
}

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	//TODO implement me
	return theme.LightTheme().Color(name, variant) // 使用默认主题颜色
	//panic("implement me")
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	//TODO implement me
	if style.Monospace {
		return theme.DefaultTheme().Font(style) // 使用默认字体
	}
	customFont, err := fyne.LoadResourceFromPath("../../字体/Noto_Sans_SC")
	if err != nil {
		return theme.DefaultTheme().Font(style) // 加载失败时使用默认字体
	}

	return customFont
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	//TODO implement me
	panic("implement me")
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	//TODO implement me
	panic("implement me")
}

var _ fyne.Theme = (*myTheme)(nil)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.LightTheme())
	myWindow := myApp.NewWindow("闺闺模拟器")
	myWindow.Resize(fyne.NewSize(400, 400))
}

func simulator(output *widget.Label) {
	for i := 10; i > 0; i-- {
		text := fmt.Sprintf("爱你%%%d\n", i*10)
		output.SetText(output.Text + text)
	}
}
