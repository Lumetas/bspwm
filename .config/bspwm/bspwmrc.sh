#! /bin/bash

#АВТОСТАРТ
sxhkd &
numlockx &
xsetroot -cursor_name left_ptr &
compton -f --vsync drm &
feh --bg-scale bg4.jpg &
tint2 & 
setxkbmap -option 'ctrl:nocaps' &
setxkbmap "us,ru" ",winkeys" "grp:alt_shift_toggle" &


#ВОРКСПЕЙСЫ
bspc monitor -d             


#РАМКИ И ГАПСЫ
bspc config border_width         1
bspc config window_gap          10

bspc config normal_border_color "#23252e"
bspc config active_border_color "#23252e"
bspc config focused_border_color "#f9f8fe"

#ЦВЕТ ПРОСТРАНСТВА ПРИ ВЫБОРЕ ПОЛОЖЕНИЯ НОВОГО ОКНАЫ
bspc config presel_feedback_color "fd2e59"

#ПРИ ПОМОЩИ ЭТОЙ ХРЕНИ МОЖНО СДЕЛАТЬ ГАПСЫ ТОЛЬКО МЕЖДУ ОКНАМИ
#bspc config window_gap $gap;
#bspc config top_padding -$(($gap+2))
#bspc config left_padding -$(($gap+2))
#bspc config right_padding -$(($gap+4))
#bspc config bottom_padding $(($PANEL_HEIGHT-$gap))



#КОФИЦИЕНТ РАЗДЕЛЕНИЯ, СКОЛЬКО ПРОЦЕНТОВ ОСТАНЕТСЯ ОТ РОДИТЕЛЬСКОГО ОКНА ПРИ ОТКРЫТИИ НОВОГО
bspc config split_ratio          0.5
#ФОКУС ОКНА НАВЕДЕНИЕМ КУРСОРА
bspc config focus_follows_pointer true
bspc config pointer_modifier	super

#РЕЖИМ МОНОКЛЯ КОГДА ОТКРЫТО ОДНО ОКНО
bspc config single_monocle	     false

#ОТКЛЮЧЕНИЕ ГАПСОВ РАМОК И ОТСТУПОВ КОГДА ОТКРЫТО ОДНО ОКНО
bspc config borderless_monocle   false
bspc config gapless_monocle      false
bspc config paddingless_monocle	 true


#СХЕМЫ СОГЛАСНО КОТОРЫМ БУДУТ ОТКРЫВАТСЯ ОКНА
#bspc config automatic_scheme	spiral


#ПРАВИЛА ОКОН
bspc rule -a dmenu rectangle=900x22+610+90
