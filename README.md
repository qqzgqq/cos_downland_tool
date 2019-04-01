D:\go_workplace\src\cos_download>go run main.go -h

usage: cosdownload -url $BucketURL -sd $SecretID -sk $SecretKEY -fnm $FileNAME<br>
       cosdownload -url $BucketURL -sds $SecretID -sks $SecretKEY -fnm $FileNAME<br>
-url   bucket_url<br>
-sd    secretid<br>
-sk    secretkey<br>
-sds   receive encrypt secretid<br>
-sks   receive encrypt secretkey<br>
-fnm   the file dir and file name from bucket<br>
-glt   get the new file list from cos bucket<br>
-enc   encrypt the string will use as encrypt secretid or secretkey<br>

eg 1:<br>
       cosdownload -url https://********* -sd ****** -sk ****** -glt<br>
       please cat the file name as bucketdirlist.log in this dir , get you will pull the file dir and name<br>
       cosdownload -url https://********* -sd ****** -sk ****** -fnm /APP_BACKUP/test/test.tar.gz<br>
eg 2:<br>
       ssecretid=`cosdownload-v1 -enc secretid`; ssecretkey=`cosdownload-v1 -enc secretkey`<br>
       cosdownload -url https://********* -sds ****** -sks ****** -glt<br>
       please cat the file name as bucketdirlist.log in this dir , get you will pull the file dir and name<br>
       cosdownload -url https://********* -sds $ssecretid -sks $ssecretkey -fnm /APP_BACKUP/test/test.tar.gz

<br>
# **下载**<br>
[点击下载](https://github.com/qqzgqq/cos_downland_tool/releases/tag/V1.0)