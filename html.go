package pretty

// github.com/tdewolff/minify binding for glub.
// No Configuration required.

import (
	"sync"
   "bytes"
   "fmt"
	"github.com/yosssi/gohtml"
	"github.com/omeid/slurp"
)

func Pretty(c *slurp.C) slurp.Stage {
	return func(in <-chan slurp.File, out chan<- slurp.File) {

		var wg sync.WaitGroup
		defer wg.Wait()

         for file := range in {
            wg.Add(1)
            go func(file slurp.File) {
               defer wg.Done()

               html := new(bytes.Buffer)
               _, err := html.ReadFrom(file.Reader)
               if err != nil {
                  fmt.Println(err.Error())
                  return
               }
               pretty := bytes.NewBufferString(gohtml.Format(html.String()))

               file.Reader = pretty
               file.FileInfo.SetSize(int64(pretty.Len()))
               out <- file
            }(file)
		}
	}
}
