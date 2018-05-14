<?php

function SimpleAdding($num)
{

    $result = 0;
    foreach (range(1, $num) as $k => $v) {
        $result += $v;
    }

    return $result;

}

// keep this function call here
echo SimpleAdding(fgets(fopen('php://stdin', 'r')));
