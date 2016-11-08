package main
import (
	"fmt"
	"Dictionary/dict"
	xx "Dictionary/trainer"
	_ "path/filepath"
	"os"
	"sort"
	"strings"

)
var AllDictFiles []string
//var AllDictionaries []dict.Dictionary

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
	var errore error
	AllDictFiles,errore=dict.GetFilesFromDirectory(CurDir + dict.Seppl + TESTsrc )
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
	fmt.Println("             TEST 3")
	test3()
	fmt.Println("             TEST 3 ENDE")

	os.Exit(1)
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
	SWITCH:
	switch x {
	default:
		//fmt.Println("------------------")
		goto NEXTE
	/*case "ä","ü","ß","ö": // sobald des an Umlaut od. scharfes ß isch, gibts (IN DER EXE!!) a Dauerschleife
		x=""*/

	//case ".":
	case "-":
		if nRandomLang==1 {
			//nRandomLang=0
			//fmt.Printf("                       %v",dict.Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[nRandomLang])
			fmt.Printf("%v = %v\n",dict.Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[nRandomLang],dict.Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[0])
		}else {
			//nRandomLang=1
			//fmt.Printf("                       %v",dict.Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[nRandomLang])
			fmt.Printf("%v = %v\n",dict.Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[nRandomLang],dict.Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[1])
		}
		//goto NOMOL
		fmt.Printf("------------------\n")
		goto NEXTE
	/*case "XX-":
		fmt.Printf("%v \n__________________\n\n",dict.Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[0] + " = " + dict.Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[1])
		goto NEXTE
	*/
	case ",":
	case ":d":
		fmt.Printf("RandomDict: %v    RandomVoc: %v     RandomLang: %v\n",dict.Dictionaries[nRandomDict].Name, nRandomVoc, nRandomLang)
		goto NOMOL
	case ":s": //Suche Vocable
		fmt.Println("=====================SUCHE START")
		SUCHAGAIN:
		fmt.Printf("%v",":s >")
		x=""
		fmt.Scan(&x)
		if strings.HasPrefix(x,"!"){
			fmt.Println("=====================SUCHE ENDE")
			goto SWITCH
		}
		s,anz:=such(x)
		if anz==0{
			fmt.Printf("%v\n", "not found")
		}else {
			fmt.Printf("%v\n", s)
		}
		goto SUCHAGAIN
	case "esc",":q",":Q":
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

func test3(){
	//var x  dict.Dictionaries
 	 a,b,x:= dict.SearchStringInDicts2("abb")
	fmt.Printf("  %v\n  %v\n  %v\n",a,b,x)
}

func such(s string)(gef string, anz int){
	//fmt.Println("--------------Search Start")
	s= strings.ToLower(s)
	var lvoc string
	var lvocRet string
	//for i,r:=range AllDictFiles{
	var d dict.Dictionary
	nfund:=0
	nfundGesamt:=0
	ii:=0
	//i:=0
	for _ ,d = range dict.Dictionaries{
		//fmt.Printf("DICTIONARY: %v %v:\n",i,d.Name)
		nfund=0
		lvoc=""
		for ii, _ = range d.Vocables {
			//fmt.Printf("  ii:%v\n",ii)
			if strings.Contains(strings.ToLower(d.Vocables[ii].Languages[0]),s)||strings.Contains(strings.ToLower(d.Vocables[ii].Languages[1]),s){
				nfund++
				nfundGesamt++
				if lvoc =="" {
					//fmt.Printf("    Fund: %v\n", d.Vocables[ii])
					lvoc = "    " + d.Vocables[ii].Languages[0] + " = " + d.Vocables[ii].Languages[1]
				}else{
					lvoc = lvoc + "\n" + "    " + d.Vocables[ii].Languages[0] + " = " + d.Vocables[ii].Languages[1]
				}
			}
		}
		if nfund > 0 {
			if lvocRet == "" {
				//lvocRet = fmt.Sprintf("%v (%v Einträge)", strings.ToUpper(d.Name) ,nfund)
				lvocRet = fmt.Sprintf("%v (%v)\n%v",strings.ToUpper(d.Name), nfund, lvoc)
			} else {
				lvoc =  fmt.Sprintf("\n\n%v (%v) \n%v", strings.ToUpper(d.Name), nfund , lvoc)
				lvocRet = lvocRet + lvoc
			}
		}
	}
	//fmt.Println("--------------Search END")
//RAUSS:
	return lvocRet, nfundGesamt
}
func Susi(i int)int{
	return  i+777
}
