<?php

function test($x)
{
	if ($x <= 2) {
		return 1;
	} else {
		return test($x-1) + test($x-2);
	}
}



for ($i = 1; $i <= 10; $i++) {
	echo test($i)."\n";
}