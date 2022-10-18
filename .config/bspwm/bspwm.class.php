<?php
class Bspwm {


    
    public function autostart($arr){
        $i ='';
        foreach ($arr as &$val) {
            $i = "$i $val & ";
}
    shell_exec($i);}

    public function config($a, $b){
        shell_exec("bspc config $a $b");
    }

    public function rule($rule){
        shell_exec('bspc rule '.$rule);
    }
    public function bspc($a){
        shell_exec("bspc $a");
    }
    public function custom($a){
        shell_exec($a);
    }

}

$bspwm = new Bspwm;
shell_exec('bspc monitor -d             ');
