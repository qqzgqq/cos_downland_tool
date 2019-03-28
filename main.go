package main

import (
	"cos_download/cosdownload"
	"cos_download/restore"
	"cos_download/upset"
	"flag"
	"fmt"
	"os"
)

// receive is usege scanln user input info
var receive, receive2 string

// Helpdoc if input -h or --h  then print helpdoc
var Helpdoc *bool

// Helpdoc2 if input -help or --help then print helpdoc
var Helpdoc2 *bool

// GetList get file list from cos bucket
var GetList *bool

// BucketURL use receive bucket_url
var BucketURL *string

// SecretID use receive secretid
var SecretID *string

// SecretKEY use receive secretkey
var SecretKEY *string

// SSecretID use encrypt secretid
var SSecretID *string

// SSecretKEY use encrypt secrekey
var SSecretKEY *string

// ENCryptstring use encrypt string
var ENCryptstring *string

// BFILEName the file name in bucket
var BFILEName *string

//init fast start
func init() {
	BucketURL = flag.String("url", receive, "bucket_url")
	Helpdoc = flag.Bool("h", false, "-h for helpdoc")
	Helpdoc2 = flag.Bool("help", false, "--help for helpdoc")
	GetList = flag.Bool("glt", false, "get file list from cos bucket")
	SecretID = flag.String("sd", receive, "secretid")
	SecretKEY = flag.String("sk", receive, "secretkey")
	SSecretID = flag.String("sds", receive, "encrypt secretid")
	SSecretKEY = flag.String("sks", receive, "encrypt secretkey")
	BFILEName = flag.String("fnm", receive, "the file name in bucketthe file name in bucket")
	ENCryptstring = flag.String("enc", receive, "encrypt string")

}

// helpinfo cosdownload-v1 helpinfo
func helpinfo(s, s2 *bool) {
	var receive2 = flag.Arg(0)
	if *s || *s2 {
		fmt.Println("usage: cosdownload -url $BucketURL -sd $SecretID -sk $SecretKEY -fnm $FileNAME")
		fmt.Println("       cosdownload -url $BucketURL -sds $SecretID -sks $SecretKEY -fnm $FileNAME")
		fmt.Printf("%5s  %s\n", "-url ", "bucket_url")
		fmt.Printf("%5s  %s\n", "-sd  ", "secretid")
		fmt.Printf("%5s  %s\n", "-sk  ", "secretkey")
		fmt.Printf("%5s  %s\n", "-sds ", "receive encrypt secretid")
		fmt.Printf("%5s  %s\n", "-sks ", "receive encrypt secretkey")
		fmt.Printf("%5s  %s\n", "-fnm ", "the file dir and file name from bucket")
		fmt.Printf("%5s  %s\n", "-glt ", "get the new file list from cos bucket")
		fmt.Printf("%5s  %s\n", "-enc ", "encrypt the string will use as encrypt secretid or secretkey ")
		fmt.Printf("%5s\n", "eg 1:")
		fmt.Println("       cosdownload -url https://********* -sd ****** -sk ****** -glt")
		fmt.Println("       please cat the file name as bucketdirlist.log in this dir , get you will pull the file dir and name")
		fmt.Println("       cosdownload -url https://********* -sd ****** -sk ****** -fnm /APP_BACKUP/test/test.tar.gz")
		fmt.Printf("%5s\n", "eg 2:")
		fmt.Println("       ssecretid=`cosdownload-v1 -enc secretid`; ssecretkey=`cosdownload-v1 -enc secretkey` ")
		fmt.Println("       cosdownload -url https://********* -sds ****** -sks ****** -glt")
		fmt.Println("       please cat the file name as bucketdirlist.log in this dir , get you will pull the file dir and name")
		fmt.Println("       cosdownload -url https://********* -sds $ssecretid -sks $ssecretkey -fnm /APP_BACKUP/test/test.tar.gz")
		os.Exit(0)
	}
	if receive2 != "" {
		fmt.Println("parameter error,please check `cosdownload -h`")
		os.Exit(0)
	}
}

// CHEckouT *string to string
func CHEckouT(s *string) string {
	return *s
}

func main() {
	flag.Parse()
	helpinfo(Helpdoc, Helpdoc2)

	ENCryptstringS := CHEckouT(ENCryptstring)
	if ENCryptstringS != "" {
		fmt.Println("++++++++++++++encrypt string+++++++++:", upset.StringUPSET(ENCryptstringS))
		os.Exit(0)
	}
	BucketURLS := CHEckouT(BucketURL)
	SecretIDS := CHEckouT(SecretID)
	SSecretIDS := CHEckouT(SSecretID)
	SecretKEYS := CHEckouT(SecretKEY)
	SSecretKEYS := CHEckouT(SSecretKEY)
	BFILENameS := CHEckouT(BFILEName)
	if BucketURLS == "" {
		fmt.Println("BucketURL is not find,please check `cosdownload -h`")
		os.Exit(0)
	}
	if SecretIDS == "" && SSecretIDS == "" {
		fmt.Println("SecretID is not find,please check `cosdownload -h`")
		os.Exit(0)
	}
	if SecretKEYS == "" && SSecretKEYS == "" {
		fmt.Println("SecretKEY is not find,please check `cosdownload -h`")
		os.Exit(0)
	}
	if SSecretIDS == "" && SSecretKEYS == "" && SecretIDS != "" && SecretKEYS != "" {
		// fmt.Println("yuanlaijiami:", SecretIDS, SecretKEYS)
		if *GetList {
			cosdownload.CosGetList(BucketURLS, SecretIDS, SecretKEYS)
			fmt.Println("please cat the file name as bucketdirlist.log in this dir")
			os.Exit(0)
		}
		if BFILENameS == "" {
			fmt.Println("the filename will be download is not find,please check `cosdownload -h`")
			os.Exit(0)
		}
		cosdownload.CosDownLoad(BucketURLS, SecretIDS, SecretKEYS, BFILENameS)
		os.Exit(0)
	} else if SecretIDS == "" && SecretKEYS == "" && SSecretIDS != "" && SSecretKEYS != "" {
		// fmt.Println("jiami:", SSecretIDS, SSecretKEYS)
		SSecretIDS := restore.StringRESTORE(SSecretIDS)
		SSecretKEYS := restore.StringRESTORE(SSecretKEYS)
		if *GetList {
			cosdownload.CosGetList(BucketURLS, SSecretIDS, SSecretKEYS)
			fmt.Println("please cat the file name as bucketdirlist.log in this dir")
			os.Exit(0)
		}
		if BFILENameS == "" {
			fmt.Println("the filename will be download is not find,please check `cosdownload -h`")
			fmt.Println("please cat the file name as bucketdirlist.log in this dir")
			os.Exit(0)
		}
		cosdownload.CosDownLoad(BucketURLS, SSecretIDS, SSecretKEYS, BFILENameS)
		os.Exit(0)
	} else {
		fmt.Println("cosdownload can't use secretid and encrypt secretid as the same time!")
		fmt.Println("cosdownload can't use secretkey and encrypt secretkey as the same time!")
		os.Exit(0)
	}
}
