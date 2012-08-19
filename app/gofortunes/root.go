package gofortunes

import (
	"net/http"
	"fmt"
)

const rootHTML = `
<html>
 <body>
   <h1> Welcome to GoFortunes </h1>
   <ul>
	<li> <a href="addForm"> Add a new fortune </a> </li>
	<li> <a href="getForm"> Get fortune </a> </li>
	<li> <a href="restoreForm"> Restore from a fortune DB </a></li>
  </ul>
 </body>
</html>
`

func Root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rootHTML)
}
