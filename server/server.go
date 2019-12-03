package server
import(
	"github.com/gorilla/mux"
)
func setServer(port string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", )
}
