package main

import (
   "fmt"
   "github.com/89z/format/hls"
   "github.com/89z/mech/twitter"
   "net/http"
   "os"
)

func doSpace(id string, info bool) error {
   guest, err := twitter.NewGuest()
   if err != nil {
      return err
   }
   space, err := guest.AudioSpace(id)
   if err != nil {
      return err
   }
   if info {
      fmt.Println(space)
   } else {
      source, err := guest.Source(space)
      if err != nil {
         return err
      }
      fmt.Println("GET", source.Location)
      res, err := http.Get(source.Location)
      if err != nil {
         return err
      }
      defer res.Body.Close()
      seg, err := hls.NewSegment(res.Request.URL, res.Body)
      if err != nil {
         return err
      }
      ext, err := seg.Ext()
      if err != nil {
         return err
      }
      file, err := os.Create(space.Base() + ext)
      if err != nil {
         return err
      }
      defer file.Close()
      for i, info := range seg.Info {
         fmt.Println(i, len(seg.Info)-1)
         res, err := http.Get(info.URI)
         if err != nil {
            return err
         }
         if _, err := file.ReadFrom(res.Body); err != nil {
            return err
         }
         if err := res.Body.Close(); err != nil {
            return err
         }
      }
   }
   return nil
}
