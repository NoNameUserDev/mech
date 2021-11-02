package instagram

import (
   "fmt"
   "os"
   "testing"
)

const (
   like = "CUrAS88Pr1G"
   video = "CUWBw4TM6Np"
)

func TestLike(t *testing.T) {
   Verbose(true)
   m, err := GraphQL(like, nil)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", m)
}

func TestData(t *testing.T) {
   f, err := os.Open(appData)
   if err != nil {
      t.Fatal(err)
   }
   defer f.Close()
   var auth Login
   if err := auth.Decode(f); err != nil {
      t.Fatal(err)
   }
   Verbose(true)
   m, err := NewQuery(like).Data(&auth)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", m)
}

func TestItem(t *testing.T) {
   f, err := os.Open(appData)
   if err != nil {
      t.Fatal(err)
   }
   defer f.Close()
   var auth Login
   if err := auth.Decode(f); err != nil {
      t.Fatal(err)
   }
   Verbose(true)
   i, err := auth.Item(video)
   if err != nil {
      t.Fatal(err)
   }
   fmt.Printf("%+v\n", i)
}

func TestWrite(t *testing.T) {
   Verbose(true)
   l, err := NewLogin("srpen6", password)
   if err != nil {
      t.Fatal(err)
   }
   if err := l.Encode(os.Stdout); err != nil {
      t.Fatal(err)
   }
}
