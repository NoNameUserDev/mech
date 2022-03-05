// source: https://stackoverflow.com/questions/64179218/retry-http-request-roundtrip
package instagram

import (
   "context"
   "log"
   "net/http"
   "time"
)

type Retry struct {
   nums      int
   transport http.RoundTripper
}

func (r *Retry) RoundTrip(req *http.Request) (resp *http.Response, err error) {
   var (
      duration time.Duration
      ctx      context.Context
      cancel   func()
   )
   if deadline, ok := req.Context().Deadline(); ok {
      duration = time.Until(deadline)
   }
   for i := 0; i < r.nums; i++ {
      if duration > 0 {
         ctx, cancel = context.WithTimeout(context.Background(), duration)
         req = req.WithContext(ctx)
      }
      log.Println("Attempt: ", i+1)
      resp, err = r.transport.RoundTrip(req)
      if resp != nil && err == nil {
         return
      }
      log.Println("Retrying...")
   }
   defer cancel()
   return
}
