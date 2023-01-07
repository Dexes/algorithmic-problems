<?php

function XO($s) {
    $x = 0;
    $o = 0;
    foreach (str_split($s) as $ch) {
        if ($ch == 'x' || $ch == 'X') {
            $x++;
            continue;
        }

        if ($ch == 'o' || $ch == 'O') {
            $o++;
        }
    }

    return $x == $o;
}