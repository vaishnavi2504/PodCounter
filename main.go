package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"net/http"
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)


//Global variable
var config *rest.Config

func initialise(){
    var useLocalAuth = true
	var kubeconfig *string
	var err error

	if useLocalAuth {
		//This flag ensures that configis read from ~/.kube/config while running locally as a stan alone go app or as a docker container
		home := homedir.HomeDir()
		fmt.Println("Will read local kube configuration for authentication.  Home is [", home, "]")
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		flag.Parse()
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	} else {
		//Else reads config from cluster
		fmt.Println("Will read cluster kube configuration for authentication")
		config, err = rest.InClusterConfig()
	}

	if err != nil {
		fmt.Println("Failed to build config")
		panic(err.Error())
	}

}

func podCount() int{
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("failed to build clientset")
		panic(err)
	}
	pods,err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
		return -1
	}
	return len(pods.Items)
}

//Json object which is returned as response
type Resp struct {
	Count    int     `json:"podcount"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	var pc  = podCount()
	w.Header().Set("Content-Type", "application/json")
	resp := Resp {
				  Count: pc,
			 }

   json.NewEncoder(w).Encode(resp)
}

func main() {
	initialise()
	fmt.Println("Starting pod count server")
	http.HandleFunc("/", jsonHandler)
    http.ListenAndServe(":8081", nil)
}