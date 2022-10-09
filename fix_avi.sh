#!/bin/sh

[ -n "$1" ] || {

echo "Empty file name"

exit 1

}



file=$1

code="${2:-FMP4}"



echo $file $code

echo $code | dd conv=notrunc of="$file" bs=b1 count=${#code} seek=112

#echo $code | dd conv=notrunc of="$file" bs=1 count=${#code} seek=188

echo patched!