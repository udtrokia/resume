#!/bin/bash
# path
if [[ "$(uname)" == 'Darwin' ]]; then
    DIR="$(dirname "$(realpath "$0")")"
else
    DIR="$(dirname "$(readlink -f "$0")")"
fi

# interaction
echo 'plz input the question number you want to test: "1) 2) 3) 4) 5)"'
read q;

# main
case ${q} in
    1)
	echo '~~~~~~~~'
	node ${DIR}/q1.js
	;;
    2)
	echo '~~~~~~~~'
	node ${DIR}/q2.js
	;;
    3)
	echo '~~~~~~~~'
	cat ${DIR}/q3.js
	;;
    4)
	echo '~~~~~~~~'
	open ${DIR}/q4.html
	;;
    5)
	echo '~~~~~~~~'
	open ${DIR}/q5.html
	;;
    *)
	echo 'question does not exist.'
	;;
esac
