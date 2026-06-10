package main 
 
import ( 
    "log" 
    "net/http" 
    "flag"
    "os"
) 
type application struct {
    errorLog *log.Logger
    infoLog *log.Logger
}
 
func main() { 
      infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
       errorLog := log.New(os.Stderr, "ERROR\t", 
log.Ldate|log.Ltime|log.Lshortfile) 
    addr := flag.String("addr",":3000", "HTTP network address")
    // flag.String,flag.Int,flag.Bool,flag.Duration,flag.Float64-> Automatic type conversion of command line input
    flag.Parse()
    // Initialize a new instance of our application struct, containing the 
    // dependencies.
    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
    }
    mux := http.NewServeMux() 
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
    mux.HandleFunc("/", home) 
    mux.HandleFunc("/snippet/view", snippetView) 
    mux.HandleFunc("/snippet/create", snippetCreate) 
    srv := &http.Server{
        Addr: *addr,
        ErrorLog: errorLog,
        Handler: mux,
        
    }
    infoLog.Printf("Starting server on %s",*addr) 
    err := srv.ListenAndServe() 
    errorLog.Fatal(err) 
}