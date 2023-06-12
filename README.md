# Change video codec DivX (XviD) description to compatible for work in new TV
FMP4 https://trac.ffmpeg.org/wiki/Encode/MPEG-4

### Change AVI FourCC via BusyBox work not in all builds
```
file=$1
code="${2:-FMP4}"

echo $code | dd conv=notrunc of="$file" bs=b1 count=${#code} seek=112
echo $code | dd conv=notrunc of="$file" bs=1 count=${#code} seek=188

echo patched!
```
So use this app

### Cross-compilation for router
env GOOS=linux GOARCH=mipsle go build *.go

### Remove any audio except eng
https://www.reddit.com/r/ffmpeg/comments/r3dccd/how_to_use_ffmpeg_to_detect_and_delete_all_non/
ffmpeg -i input.mkv -map 0:V -map 0:a:m:language:eng -c copy output.mkv