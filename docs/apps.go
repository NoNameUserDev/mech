package main

import (
   "fmt"
   "github.com/89z/format"
   "github.com/89z/googleplay"
   "os"
   "sort"
   "time"
)

var apps = []application{
   {id: "com.amazon.mp3"},
   {id: "com.apple.android.music"},
   {id: "com.aspiro.tidal"},
   {id: "com.bandcamp.android", done: true},
   {id: "com.cbs.app"},
   {id: "com.clearchannel.iheartradio.controller"},
   {id: "com.google.android.youtube", done: true},
   {id: "com.imdb.mobile", done: true},
   {id: "com.instagram.android", done: true},
   {id: "com.mtvn.mtvPrimeAndroid", done: true},
   {id: "com.nbcuni.nbc", done: true},
   {id: "com.pandora.android", done: true},
   {id: "com.pbs.video"},
   {id: "com.qobuz.music"},
   {id: "com.reddit.frontpage"},
   {id: "com.rhapsody"},
   {id: "com.soundcloud.android", done: true},
   {id: "com.spotify.music"},
   {id: "com.ted.android", done: true},
   {id: "com.tumblr", done: true},
   {id: "com.twitter.android", done: true},
   {id: "com.vimeo.android.videoapp", done: true},
   {id: "com.zhiliaoapp.musically", done: true},
   {id: "deezer.android.app"},
}

func main() {
   cache, err := os.UserCacheDir()
   if err != nil {
      panic(err)
   }
   tok, err := googleplay.OpenToken(cache + "/googleplay/token.json")
   if err != nil {
      panic(err)
   }
   auth, err := tok.Auth()
   if err != nil {
      panic(err)
   }
   dev, err := googleplay.OpenDevice(cache + "/googleplay/device.json")
   if err != nil {
      panic(err)
   }
   for i, app := range apps {
      det, err := auth.Details(dev, app.id)
      if err != nil {
         panic(err)
      }
      apps[i].installs = det.NumDownloads
      apps[i].name = det.Title
      time.Sleep(99 * time.Millisecond)
   }
   sort.Slice(apps, func(a, b int) bool {
      return apps[b].installs < apps[a].installs
   })
   for _, app := range apps {
      fmt.Println(format.Number.GetUint64(app.installs), app.done, app.name)
   }
}

type application struct {
   id, name string
   done bool
   installs uint64
}

