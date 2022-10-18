package main

import (
	"fmt"
	"os/exec"
)

func main() {



wallpaper := "~/bg4.jpg"



autostart(
"sxhkd",
"numlockx",
"xsetroot -cursor_name left_ptr",
"compton -f --vsync drm",
"tint2",
"feh --bg-scale " + wallpaper,
"flameshot")


border_width := `1`
window_gap := "10"

normal_border_color := `"#23252e"`
active_border_color := `"#23252e"`
focused_border_color := `"#f9f8fe"`

//ЦВЕТ ПРОСТРАНСТВА ПРИ ВЫБОРЕ ПОЛОЖЕНИЯ НОВОГО ОКНА
presel_feedback_color := `"fd2e59"`




/*  ПРИ ПОМОЩИ ЭТОЙ ХРЕНИ МОЖНО СДЕЛАТЬ ГАПСЫ ТОЛЬКО МЕЖДУ ОКНАМИ

    bspc config window_gap gap
    bspc config top_padding -((gap+2))
    bspc config left_padding -((gap+2))
    bspc config right_padding -((gap+4))
    bspc config bottom_padding ((PANEL_HEIGHT-gap))      */



//КОФИЦИЕНТ РАЗДЕЛЕНИЯ, СКОЛЬКО ПРОЦЕНТОВ ОСТАНЕТСЯ ОТ РОДИТЕЛЬСКОГО ОКНА ПРИ ОТКРЫТИИ НОВОГО
split_ratio := `0.5`
//ФОКУС ОКНА НАВЕДЕНИЕМ КУРСОРА
focus_follows_pointer := `true`
pointer_modifier := `super`

//РЕЖИМ МОНОКЛЯ КОГДА ОТКРЫТО ОДНО ОКНО
single_monocle := `false`

//ОТКЛЮЧЕНИЕ ГАПСОВ РАМОК И ОТСТУПОВ КОГДА ОТКРЫТО ОДНО ОКНО
borderless_monocle := `false`
gapless_monocle := `false`
paddingless_monocle := `true`

//СХЕМЫ СОГЛАСНО КОТОРЫМ БУДУТ ОТКРЫВАТСЯ ОКНА
//automatic_scheme := `spiral`


//ПРАВИЛА ОКОН
dmenu_rule := `bspc rule -a dmenu rectangle:=900x22+610+90`


























	
	
	
	
	
	
	
	
	
	






config(`border_width`, border_width)
config(`window_gap`, window_gap)
config(`normal_border_color`, normal_border_color)
config(`active_border_color`, active_border_color)
config(`focused_border_color`, focused_border_color)
config(`presel_feedback_color`, presel_feedback_color)
config(`split_ratio`, split_ratio)
config(`focus_follows_pointer`, focus_follows_pointer)
config(`pointer_modifier`, pointer_modifier)
config(`single_monocle`, single_monocle)
config(`borderless_monocle`, borderless_monocle)
config(`gapless_monocle`, gapless_monocle)
config(`paddingless_monocle`, paddingless_monocle)
shell(dmenu_rule)
shell("bspc monitor -d  1 2 3 4 5 6 7 8 9 10")
}












































func config(a string, b string) {
	cmd := exec.Command("bspc", "config", a, b) 
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
}


func autostart(args ...string) {
   total := ""
    for _, v := range args {

total += v + " & "


    }


	cmd := exec.Command("bash", "-c", total)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
}


func shell(a string){

	cmd := exec.Command("bash", "-c", a)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}

}

