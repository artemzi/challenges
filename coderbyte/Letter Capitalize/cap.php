<?php

function LetterCapitalize($str)
{

    return ucwords($str);

}

// keep this function call here
echo LetterCapitalize(fgets(fopen('php://stdin', 'r')));
