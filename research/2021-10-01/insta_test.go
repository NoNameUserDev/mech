package goinsta

import (
   "os"
   "testing"
)

func TestInsta(t *testing.T) {  
   pass, ok := os.LookupEnv("PASS")
   if ! ok {
      t.Fatal("PASS")
   }
   insta := New("srpen6", pass)
   err := insta.Login()
   if err != nil {
      panic(err)
   }
   insta.Export("ig.json")
}
