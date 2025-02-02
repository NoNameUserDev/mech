package main

import (
   "flag"
   "fmt"
   "github.com/89z/mech/youtube"
   "os"
)

func main() {
   var choose choice
   // a
   var address string
   flag.StringVar(&address, "a", "", "address")
   // b
   var videoID string
   flag.StringVar(&videoID, "b", "", "video ID")
   // c
   var construct bool
   flag.BoolVar(&construct, "c", false, "OAuth construct request")
   // e
   var embed bool
   flag.BoolVar(&embed, "e", false, "use embedded player")
   // f
   choose.itags = make(map[string]bool)
   flag.Func("f", "formats", func(itag string) error {
      choose.itags[itag] = true
      return nil
   })
   // i
   flag.BoolVar(&choose.info, "i", false, "information")
   // r
   var refresh bool
   flag.BoolVar(&refresh, "r", false, "OAuth token refresh")
   // v
   var verbose bool
   flag.BoolVar(&verbose, "v", false, "verbose")
   // x
   var exchange bool
   flag.BoolVar(&exchange, "x", false, "OAuth token exchange")
   flag.Parse()
   if verbose {
      youtube.LogLevel = 1
   }
   if exchange {
      oauth, err := youtube.NewOAuth()
      if err != nil {
         panic(err)
      }
      fmt.Println(oauth)
      fmt.Scanln()
      exc, err := oauth.Exchange()
      if err != nil {
         panic(err)
      }
      cache, err := os.UserCacheDir()
      if err != nil {
         panic(err)
      }
      if err := exc.Create(cache + "/mech/youtube.json"); err != nil {
         panic(err)
      }
   } else if refresh {
      cache, err := os.UserCacheDir()
      if err != nil {
         panic(err)
      }
      exc, err := youtube.OpenExchange(cache + "/mech/youtube.json")
      if err != nil {
         panic(err)
      }
      if err := exc.Refresh(); err != nil {
         panic(err)
      }
      if err := exc.Create(cache + "/mech/youtube.json"); err != nil {
         panic(err)
      }
   } else if videoID != "" || address != "" {
      if videoID == "" {
         var err error
         videoID, err = youtube.VideoID(address)
         if err != nil {
            panic(err)
         }
      }
      play, err := player(construct, embed, videoID)
      if err != nil {
         panic(err)
      }
      if err := choose.adaptiveFormats(play); err != nil {
         panic(err)
      }
   } else {
      flag.Usage()
   }
}
