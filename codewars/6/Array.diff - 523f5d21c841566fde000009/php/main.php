<?php

function arrayDiff($a, $b) {
    $map = [];
    foreach ($b as $item) {
        $map[$item] = true;
    }

    $result = [];
    foreach ($a as $item) {
        if (!isset($map[$item])) $result[] = $item;
    }

    return $result;
}