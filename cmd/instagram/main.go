package main

import (
   "flag"
   "github.com/89z/mech/instagram"
)

func main() {
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // b
   var shortcode string
   flag.StringVar(&shortcode, "b", "", "shortcode")
   // auth
   var auth bool
   flag.BoolVar(&auth, "auth", false, "authentication")
   // i
   var info bool
   flag.BoolVar(&info, "i", false, "info only")
   // p
   var password string
   flag.StringVar(&password, "p", "", "password")
   // u
   var username string
   flag.StringVar(&username, "u", "", "username")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   // o
   var output string
   flag.StringVar(&output, "o", "", "output")
   flag.Parse()
   if verbose {
      instagram.LogLevel = 1
   }
   if username != "" {
      err := saveLogin(username, password)
      if err != nil {
         panic(err)
      }
   } else if shortcode != "" || address != "" {
      if shortcode == "" {
         shortcode = instagram.Shortcode(address)
      }
      if auth {
         err := doItems(shortcode, info, output)
         if err != nil {
            panic(err)
         }
      } else {
         err := doGraph(shortcode, info, output)
         if err != nil {
            panic(err)
         }
      }
   } else {
      flag.Usage()
   }
}
