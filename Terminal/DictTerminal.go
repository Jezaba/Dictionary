package main
import (
	"fmt"
	"Dictionary/dict"
	xx "Dictionary/trainer"
	_ "path/filepath"
	"os"
	"sort"
	"strings"
	"bufio"
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

	fmt.Println(xx.TrainerTest())

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

	menu()
	os.Exit(1)


	fmt.Println("TEST 3 <<<<<<<<<<<<<<<")
	//test3()
	fmt.Println(">>>>>>>>>>>>>TEST 3 ENDE")


	RAUSS:
	fmt.Println("==================== ENDE ====================")

}

func training(){
	var x string
	NEXTE:
	x=""

	nRandomDict:=dict.GetRandomNumber(0,len(dict.Dictionaries))
	nRandomVoc:=dict.GetRandomNumber(0,len(dict.Dictionaries[nRandomDict].Vocables))
	nRandomLang:=dict.GetRandomNumber(0,2)
	fmt.Printf("%v ",dict.Dictionaries[nRandomDict].Vocables[nRandomVoc].Languages[nRandomLang])
	NOMOL:
	x=""
	readString ("",&x)
	//fmt.Println("Len x:", len(x), x)
	fmt.Printf("x=%v   len=%v\n", x,  len(x))

	//x="esc"
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
	case ":e": //EDIT
		editVocable()
		goto NOMOL
	case ":s":
		fmt.Println("=====================SUCHE START")
		SUCHAGAIN:
		x=""
		readString (":s >",&x)
		if strings.HasPrefix(x,"!"){
			fmt.Println("=====================SUCHE ENDE")
			goto SWITCH
		}
		such(&x)

		goto SUCHAGAIN
	case "esc",":q",":Q",":!":
		goto RAUSS
	}

	RAUSS:
}
const inputdelimiter = '\n'
func readString(consolesetext string,input *string)  {
	if consolesetext !="" {
		fmt.Print(consolesetext)
	}
	reader := bufio.NewReader(os.Stdin)
	var err error
	*input, err = reader.ReadString(inputdelimiter)
	if err != nil {
		fmt.Println("FEHLER:", err)
		return
	}
	// convert CRLF to LF
	*input = strings.Replace(*input, "\r\n", "", -1)
	*input = strings.Replace(*input, "\n", "", -1)
}

func such(searchstring *string){
	D:= dict.SearchStringInDicts(*searchstring)
	if len(D)>0{
		//fmt.Printf("Länge %v\n", D[0])
		for i:=0;i<len(D);i++{
			fmt.Printf("Dictname: %v  AnzahlVocs(%v)\n",D[i].Name,len(D[i].Vocables))
			for ii:=0 ;ii<len(D[i].Vocables);ii++{
				fmt.Printf("  d[%v].v[%v] %v = %v  \n" ,i,ii,D[i].Vocables[ii].Languages[0],D[i].Vocables[ii].Languages[1])
			}
		}
	}
}


func editVocable(){
	fmt.Println("BeginEdit")
	var x string
	x=""
	old:=dict.Dictionaries[0].Vocables[0]
	fmt.Printf("EditVocable(dict[0].voc[0].lang[1]%v\n>",dict.Dictionaries[0].Vocables[0].Languages[1])
	fmt.Scan(&x)
	dict.EditVocable(0,0,1,x)
	fmt.Printf("aus Old '%v' wurde '%v'\nEndEdit\n",old,dict.Dictionaries[0].Vocables[0].Languages[1])
}

func menu(){
	ANFANG:
	fmt.Print(`Which example do you want to run?
1) Vokabeltrainer
2) bufio.Reader.ReadString(...)
3) bufio.Reader.ReadByte(...)
4) ENDE()
Please enter 1..5 and press ENTER: `)

	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
		return
	}

	switch result {

	case '1':
		training()
		goto ANFANG
		break

	case '2':
		//runReadString()
		goto ANFANG
		break

	case '3':
		editVocable()
		goto ANFANG
		break

	case '4':
		goto ENDE
		break

	default:
		return
	}
	ENDE:
}
