package dict

import (
	"math/rand"
	"strings"
	"os"
	"fmt"
	"bufio"
	"time"
	"io/ioutil"
	"path/filepath"
	"sort"
)

var Dictionaries []Dictionary
type Dictionary struct {
	Name string
	LanguageSeparator string
	WordSeparator     string
	NumberOfLanguages int
	Vocables          []Vocable
}
type Vocable struct {
	Languages  []string// alles aufgeteilt in die einzelnen Sprachen
}

type ByLength []string
func (s ByLength) Len() int {
return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

type ByVocable []Vocable //https://gobyexample.com/sorting-by-functions
func (s ByVocable) Len() int {
	return len(s)
}
func (s ByVocable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByVocable) Less(i, j int) bool {
	return  strings.ToLower(s[i].Languages[0]) <  strings.ToLower( s[j].Languages[0])
}

type ByLanguage []string
func (s ByLanguage) Len() int {
	return len(s)
}
func (s ByLanguage) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLanguage) Less(i, j int) bool {
	return true// len(s[i]) < len(s[j])
}

/*func GetVocable(idDict int, index int) (Vocable,error){
	var voc Vocable
	var err error
	if len(Dictionaries)<idDict{
		err= fmt.Errorf("(%v)übersteigt die Anzahl der vorhandenen Dictionaries(%v)",idDict,len(Dictionaries))
		goto RAUSS
	}
	if len(Dictionaries[idDict].Vocables)<index{
		err= fmt.Errorf("(%v)übersteigt die Anzahl der vorhandenen Vokabeln(%v)", index,len(Dictionaries[idDict].Vocables))
		goto RAUSS
	}
	voc = Dictionaries[idDict].Vocables[index]
	RAUSS:
	return voc,err
}*/
/*func (v *Vocable)GetWordsPerLanguage(nLanguage int, dicti *Dictionary)[]string{
	var s []string
	s=strings.Split(v.Languages[nLanguage], dicti.WordSeparator)
	return s
}*/

func (w *Vocable)String(sep string)string{
	return strings.Trim(strings.Join(w.Languages, sep)," == ")
	//return strings.Join(w.Languages, Seppl)
}
func splitString3(voc *Vocable, s string) bool {
	//s=strings.Replace(s, "\n", "\n", -1)
	if strings.Contains(s,"\\n"){
		s=strings.Replace(s, "\\n", "\n",-1)
	}
	if strings.Contains(s,"\t"){
		//fmt.Println(strings.Replace(s,"\t","....",-1))
		s=strings.Replace(s,"\t"," ",-1)
	}
	lbChanged := false
	voc.Languages=make([]string, NumberOfLanguages)
	sp := strings.Split(s, VocableDevider)
	//fmt.Printf("Do 3 split1: %q  --  split2: %q\n", sp[0] , sp[1])
	anz:=len(sp)
	//fmt.Printf("Do0: anz:=len(split) %v, %v , 0: %v , 1: %v \n",anz,s, split[0], split[1])
	if anz<=1{// Kein einziger Separator
		if strings.Trim(s," ")==""{//Leerer String
			lbChanged = false
		}else {
			voc.Languages[0] = strings.Trim(s," ")
			lbChanged = true
		}
	}else {

		if anz > len(voc.Languages) {
			anz = len(voc.Languages)
		}
		for i:=0; i<anz;i++{
			//fmt.Printf("i%v:, len(voc.Languages[0])=%v, Von '%v' zu '%v'.\n",i,len(voc.Languages[i]), voc.Languages[0], strings.Trim(split[1]," "))
			/*if i==1{
				fmt.Printf("voc.Languages[1]:%v\n",voc.Languages[1])
			}*/
			/*if i==0{
				fmt.Printf("voc.Languages[0]:%v\n",voc.Languages[0])
			}*/
			//fmt.Printf("i%v:, voc.Languages[i]:%v\n",i,voc.Languages[i])
			voc.Languages[i]= strings.Trim(sp[i]," ")
			lbChanged = true
		}
	}
	return lbChanged
}
func GetVocsFromStrings3(s []string) []Vocable {
	lbChanged := false
	var words []Vocable
	for _, r := range s {
		w := Vocable{}
		lbChanged = splitString3( &w,r)
		if lbChanged {
			words = append(words, w)
		}
	}
	return words
}
func GetVocsFromFiles3(dictfile string) ([]Vocable,error) {
	//read Lines reads a whole file into memory and returns a slice of its lines.
	var vocs []Vocable
	file, err := os.Open(dictfile)
	if err != nil {
		//fmt.Println("ERROR GetVocsFromFiles3: ", err)
		//return nil,err
		return nil, fmt.Errorf("ERROR in GetVocsFromFiles3(): %v", err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	vocs = GetVocsFromStrings3(lines)
	return vocs, err
}

func GetRandomNumber(min_GroesserGleich, maxKleiner int)int{
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(maxKleiner - min_GroesserGleich) + min_GroesserGleich
}
func GetFilesFromDirectory(directory string)([]string,error){
	var s []string
	//files2, _ := ioutil.ReadDir("./")
	files2, err := ioutil.ReadDir(directory)
	goto NewVersion
	for _, f := range files2 {
		if strings.HasPrefix(f.Name(), strings.ToLower("dict")) && strings.HasSuffix(f.Name(),strings.ToLower(".txt")) {
			s=append(s,f.Name())
		}
	}

	NewVersion:
	for _, file:= range files2{
		if file.Mode().IsRegular(){
			if  filepath.Ext(file.Name()) == ".dict" {
				s=append(s,file.Name())
			}
		}
	}

	return  s,err
}

func SortDictionariesByVocables(){
	for _,r:= range Dictionaries {
		sort.Sort(ByVocable(r.Vocables))
		//fmt.Println(strings.ToUpper(Dictionaries[i].Name),r)
		/*for ii, rr := range r.Vocables {
			counter++
			sort.Sort(ByLanguage(rr.Languages))
			fmt.Println(i+1, ii+1, counter, rr.String(" = "))
		}*/
	}
}

func FindVocable(){}
func AddVocable(){}
func EditVocable(){

}
func DeleteVocable(){}


func SearchStringInDicts(s string)(retourString string, anz int){
	//fmt.Println("--------------Search Start")
	s= strings.ToLower(s)
	var lvoc string
	var lRetStr string
	//for i,r:=range AllDictFiles{
	var d Dictionary
	nfund:=0
	nfundGesamt:=0
	ii:=0
	//i:=0
	for _ ,d = range Dictionaries{
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
			if lRetStr == "" {
				//lvocRet = fmt.Sprintf("%v (%v Einträge)", strings.ToUpper(d.Name) ,nfund)
				lRetStr = fmt.Sprintf("%v (%v)\n%v",strings.ToUpper(d.Name), nfund, lvoc)
			} else {
				lvoc =  fmt.Sprintf("\n\n%v (%v) \n%v", strings.ToUpper(d.Name), nfund , lvoc)
				lRetStr = lRetStr + lvoc
			}
		}
	}
	//fmt.Println("--------------Search END")
	//RAUSS:
	return lRetStr, nfundGesamt
}
func SearchStringInDicts2(s string)(dictGefunden []Dictionary){
	//fmt.Println("--------------Search Start")
	s= strings.ToLower(s)
	//var lvoc string
	//var lRetStr string

	var dNew Dictionary
	var dictsGef []Dictionary
	var vocables []Vocable
	nfund:=0
	nfundGesamt:=0
	ii:=0
	var d Dictionary
	for _ ,d = range Dictionaries{
		//fmt.Printf("DICTIONARY: %v %v:\n",i,d.Name)
		nfund=0
		//dNew=nil
		vocables = vocables[:0]
		for ii, _ = range d.Vocables {
			//fmt.Printf("  ii:%v\n",ii)
			if strings.Contains(strings.ToLower(d.Vocables[ii].Languages[0]),s)||strings.Contains(strings.ToLower(d.Vocables[ii].Languages[1]),s){
				nfund++
				nfundGesamt++
				vocables=append(vocables,d.Vocables[ii])
			}
		}
		if nfund > 0 {
			//if lRetStr == "" {
			//	lRetStr = fmt.Sprintf("%v (%v)\n%v",strings.ToUpper(d.Name), nfund, lvoc)
				dNew =  Dictionary{}
				dNew.LanguageSeparator = d.LanguageSeparator
				dNew.NumberOfLanguages = d.NumberOfLanguages
				dNew.WordSeparator = d.WordSeparator
				dNew.Name= d.Name
				dNew.Vocables=vocables
				//fmt.Printf("		DOO dNew.Name = %v\n", dNew.Name)
				//fmt.Printf("		dNew.Vocables = %v\n", dNew.Vocables)
				dictsGef = append(dictsGef,dNew)
			//} else {
			//	lvoc =  fmt.Sprintf("\n\n%v (%v) \n%v", strings.ToUpper(d.Name), nfund , lvoc)
			//	lRetStr = lRetStr + lvoc
			//}
		}
	}
	return  dictsGef
}
var Seppl string
var VocableDevider string
var NumberOfLanguages int
func init(){
	fmt.Println("\n\n====2=====START===========")
	fmt.Printf("pkg: dict - init()\n" +
		"    Pathseperator: %v\n" +
		"    VocableDevider: %v\n" +
		"    NumberOfLanguages: %v\n",Seppl,VocableDevider,NumberOfLanguages)
	Seppl=string(filepath.Separator)
	VocableDevider = "="
	NumberOfLanguages = 5
	fmt.Println(".........................")
}



