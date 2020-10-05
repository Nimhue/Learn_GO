package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true

  var eludeGuards int
  eludeGuards = rand.Intn(100)
  fmt.Println(eludeGuards)

  if (eludeGuards <= 50) {
    fmt.Println("Psst, we got past the guards. Let's continue with the heist, but stop talking Karen, and focus on the job!")
  } else {
    fmt.Println("It's your fault, Karen! I'm done with this, just do this thing without me!")
  }

  var openedVault int
  openedVault = rand.Intn(100)
  fmt.Println(openedVault)

  if ((isHeistOn) && (openedVault >= 70)) { 
     fmt.Println("Go, go, go, go! Fire in the hall!")
   } else if isHeistOn {
     isHeistOn = false   
     fmt.Println("We cannot open the vault. Abort, abort!")   
   }

   var leftSafely int
   leftSafely = rand.Intn(5)

   if (isHeistOn) {
      switch leftSafely {
      case 0:
      isHeistOn = false
      fmt.Println("Pff, the cops are here. I ain't gonna have grilled cheese for a while :(")
      case 1:
      isHeistOn = false
      fmt.Println("Oh, darn it. The security code for the vault has changed and we don't know how to crack it.")
      case 2:
      isHeistOn = false
      fmt.Println("Shit, the guard has my gun. Run, you fools!")
      case 3:
      isHeistOn = false
      fmt.Println("Ok, so, I know I was supposed to open this vault, but I've forgot my tools at home. Oups?")
      case 4:
      isHeistOn = false
      fmt.Println("You know what? Eff you guys, this bank is too smal for me, I'm outta here. Bang. Bang. Bang.")
      default:
      isHeistOn = false
      fmt.Println("I got it, start the car. Go, go, go!")
     } 
  }

  if (isHeistOn) {
    var amtStolen int
    amtStolen = 10000 + rand.Intn(1000000)
    fmt.Println(amtStolen)
  }
  fmt.Println("Status of heist:", isHeistOn)
}


