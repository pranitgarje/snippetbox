package main 
 
import ( 
    "fmt" 
	"log"
	"html/template"
    "net/http" 
    "strconv" 
) 
 
func (app *application) home(w http.ResponseWriter, r *http.Request) { 
    if r.URL.Path != "/" { 
        http.NotFound(w, r) 
        return 
    } 
    files := [] string{
        "./ui/html/pages/home.tmpl",
        "./ui/html/pages/base.tmpl",
         "./ui/html/partials/nav.tmpl",
    }
	ts,err :=template.ParseFiles(files...)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	   // We then use the Execute() method on the template set to write the 
    // template content as the response body. The last parameter to Execute() 
    // represents any dynamic data that we want to pass in, which for now 
	//we'll 
    // leave as nil. 

	err=ts.ExecuteTemplate(w,"base",nil)	
	if err!=nil {
		log.Print(err.Error())
		http.Error(w,"Internal Server Error",500)
	}
    w.Write([]byte("Hello from Snippetbox")) 
} 
 
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) { 
    id, err := strconv.Atoi(r.URL.Query().Get("id")) 
    if err != nil || id < 1 { 
        app.errorLog.Print(err.Error())
        http.NotFound(w, r) 
        return 
    } 
 
    fmt.Fprintf(w, "Display a specific snippet with ID %d...", id) 
} 
 
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) { 
    if r.Method != http.MethodPost { 
        w.Header().Set("Allow", http.MethodPost) 
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed) 
        return 
    } 
 
    w.Write([]byte("Create a new snippet...")) 
}
