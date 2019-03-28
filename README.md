D:\go_workplace\src\cos_download>go run main.go -h

usage: cosdownload -url $BucketURL -sd $SecretID -sk $SecretKEY -fnm $FileNAME
       cosdownload -url $BucketURL -sds $SecretID -sks $SecretKEY -fnm $FileNAME
-url   bucket_url
-sd    secretid
-sk    secretkey
-sds   receive encrypt secretid
-sks   receive encrypt secretkey
-fnm   the file dir and file name from bucket
-glt   get the new file list from cos bucket
-enc   encrypt the string will use as encrypt secretid or secretkey

eg 1:
       cosdownload -url https://********* -sd ****** -sk ****** -glt
       please cat the file name as bucketdirlist.log in this dir , get you will pull the file dir and name
       cosdownload -url https://********* -sd ****** -sk ****** -fnm /APP_BACKUP/test/test.tar.gz
eg 2:
       ssecretid=`cosdownload-v1 -enc secretid`; ssecretkey=`cosdownload-v1 -enc secretkey`
       cosdownload -url https://********* -sds ****** -sks ****** -glt
       please cat the file name as bucketdirlist.log in this dir , get you will pull the file dir and name
       cosdownload -url https://********* -sds $ssecretid -sks $ssecretkey -fnm /APP_BACKUP/test/test.tar.gz