<?php

function FirstFactorial($num)
{
    $result = 1;
    foreach (range(1, $num) as $k => $v) {
        $result *= $v;
    }

    return $result;
}

// keep this function call here
echo FirstFactorial(fgets(fopen('php://stdin', 'r')));
