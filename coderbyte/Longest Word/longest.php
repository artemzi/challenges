<?php

function LongestWord($sen)
{

    $sen = preg_replace("/[^a-z\d ]/i", '', $sen);
    $result = '';
    foreach (explode(' ', $sen) as $k => $v) {
        if (strlen($result) < strlen($v)) {
            $result = $v;
        }

    }

    return $result;

}

// keep this function call here
echo LongestWord(fgets(fopen('php://stdin', 'r')));
