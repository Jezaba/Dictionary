package main
import (
	"fmt"
	"Dictionary/dict"
	xx "Dictionary/dict/trainer"
	_ "path/filepath"
	"os"
	"sort"
)

func main() {
	//var AlleDictionaryFiles []string
	CurDir, _ := os.Getwd() //filepath.Abs(filepath.Dir(os.Args[0]))

	fmt.Println("Current Directory:", CurDir)
	if len(os.Args) <=1{
		fmt.Println("No OpenArgs!!")
	}else{
		for i:=1;i<len(os.Args);i++{  //0==Name von exe
			fmt.Printf("Argument %v: %v",i,os.Args[i])
			if _, err := os.Stat(os.Args[i]); err == nil {
				fmt.Println("    DOESN'T EXIST!")
			}else{
				fmt.Println("    OK")
			}
		}
	}
	/*p:="H:\\Test\\gotest1\\src\\Dictionary\\dict\\dictionaries"
s,e:=dict.GetFilesFromDirectory(p)
if e !=nil{
	fmt.Println("EORRORE 1: " , e)
}
for _,f:=range s{
	fmt.Println(f)
}*/

	fmt.Println(xx.TrainerTest())

	//return
	TESTsrc:=""
	//TESTsrc:="dict" + seppl
	fmt.Printf("\nPath + TESTsrc = '%v\\%v'\n\n",CurDir,TESTsrc)
	AllDictFiles,errore:=dict.GetFilesFromDirectory(CurDir + dict.Seppl + TESTsrc )
	if errore !=nil{
		fmt.Println("EORRORE 2: " , errore)
	}
	fmt.Printf("%v \n","AllDictFiles:")
	for a,b:= range AllDictFiles{
		fmt.Printf("    %v: %v\n",a+1,b)
	}
	fmt.Println(".........................")

	var vocs []dict.Vocable;
	var err error;

	var counter int
	counter=-1
	if len(AllDictFiles)==0{
		fmt.Println("Kein dict-File vorhanden in '" + CurDir + dict.Seppl + TESTsrc + "'")
		goto RAUSS
	}

	for _,r:=range AllDictFiles{
		counter++
		dictus :=*new(dict.Dictionary)
		dictus.LanguageSeparator="="
		dictus.NumberOfLanguages=5

		dict.Dictionaries = append(dict.Dictionaries, dictus)
		//AlleDictionaryFiles = append(AlleDictionaryFiles, dictus)

		vocs,err =   dict.GetVocsFromFiles3(CurDir + dict.Seppl + TESTsrc + r)
		if err==nil {
			dict.Dictionaries[counter].Name = CurDir + dict.Seppl + TESTsrc + r
			dict.Dictionaries[counter].Vocables = vocs

		}else{
			fmt.Println("ERROR: ", CurDir + dict.Seppl + TESTsrc + r)
		}
	}
	counter=0
	for _,r:= range dict.Dictionaries {
		sort.Sort(dict.ByVocable(r.Vocables))
		//fmt.Println(strings.ToUpper(Dictionaries[i].Name),r)
		/*for ii, rr := range r.Vocables {
			counter++
			sort.Sort(ByLanguage(rr.Languages))
			fmt.Println(i+1, ii+1, counter, rr.String(" = "))
		}*/
	}


	test1()
	/*if test2()=="test1" {
		test1()
	}*/
	RAUSS:
	fmt.Println("==================== ENDE ====================")

}
func test2()string{
	var ret string
	ret="ölkj"
	for{
		fmt.Scan(&ret)
		switch ret {
		case "!q","!Q","esc":
			goto RAUSS
		case "test1":
			goto RAUSS
		}
	}
	RAUSS:
	return ret
}

func test1(){
	var x string
	NEXTE:
	x=""

	nRandomDict:=dict.GetRandomNumber(0,len(dict.Dictionaries))
	nRandomVoc:=dict.GetRandomNumber(0,len(dict.Dictionaries[nRandomDict].Vocables))
	nRandomLang:=dict.GetRandomNumber(0,2)
	fmt.Printf("%v ",dict.Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[nRandomLang])
	NOMOL:
	x=""
	fmt.Scan(&x)
	//fmt.Println("Len x:", len(x), x)
	if x=="esc"{
		goto RAUSS
	}

	switch x {
	default:
		fmt.Println("------------------")
		goto NEXTE
	/*case "ä","ü","ß","ö": // sobald des an Umlaut od. scharfes ß isch, gibts (IN DER EXE!!) a Dauerschleife
		x=""*/

	case ".":
		if nRandomLang==1 {
			nRandomLang=0
			fmt.Printf("                       %v",dict.Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[nRandomLang])
		}else {
			nRandomLang=1
			fmt.Printf("                       %v",dict.Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[nRandomLang])
		}
		goto NOMOL
	case "-":
		fmt.Printf("%v \n__________________\n\n",dict.Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[0] + " = " + dict.Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[1])
		goto NEXTE
	case ",":
	case "d":
		fmt.Printf("RandomDict: %v    RandomVoc: %v     RandomLang: %v\n",dict.Dictionaries[nRandomDict].Name, nRandomVoc, nRandomLang)
		goto NOMOL
	case "esc","!q","!Q":
		goto RAUSS
	}
	/*if x==" "{
		fmt.Print("Do5: ")
		if nRandomLang==1 {
			fmt.Printf(" %v  =  \n",Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[0])
		}else {
			fmt.Printf(" %v  =  \n",Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[1])
		}
		fmt.Scanf("%c",&x)
	}else{
		fmt.Println("Do2")
		goto NEXTE
	}*/
	RAUSS:

}

func Susi(i int)int{
	return  i+777
}