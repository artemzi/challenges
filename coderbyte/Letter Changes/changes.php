<?php

function LetterChanges($str)
{

    $alphabet = "abcdefghijklmnopqrstuvwxyz";
    $result = '';
    $loc = 0;

    for ($i = 0; $i < strlen($str); $i++) {
        $loc = strpos($alphabet, $str[$i]);
        if ($loc === 25) {
            $result .= "a";
        } elseif ($loc === FALSE) {
            $result .= $str[$i];
        } else {
            $result .= $alphabet[$loc + 1];
        }
    }

    return str_replace(["a", "e", "i", "o", "u"], ["A", "E", "I", "O", "U"], $result);

}

// keep this function call here
echo LetterChanges("hello world");
