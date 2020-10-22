package main

import (
	"image"
)

var icons map[string]image.Image // 保存所有的图标

func Icon(name string) image.Image {
	if icons == nil { // 延迟加载
		loadIcons()
	}
	return icons[name]
}

// 加载所有图标
func loadIcons() {
	icons = make(map[string]image.Image)
	icons["warning.png"] = loadIcon("warning.png")
	icons["header.png"] = loadIcon("header.png")
	icons["footer.png"] = loadIcon("footer.png")
}

// 加载图标
func loadIcon(fname string) image.Image {
	var icon image.Image
	// ...
	return icon
}

func main() {
	go func() { // 协程A
		icon := Icon("header.png")
		//...
	}()

	go func() { // 协程B
		icon := Icon("footer.png")
		// ...
	}()

	// ...
}
