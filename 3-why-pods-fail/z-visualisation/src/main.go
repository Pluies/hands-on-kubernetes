package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	//"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	//"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

type NSName string
type SvcName string
type Status string

const (
	NotDeployed Status = "NOT_DEPLOYED"
	Failing            = "FAILING"
	Deployed           = "DEPLOYED"
)

type ServiceAndStatus struct {
	Service    v1.Service
	Status     Status
	StatusCode int
	Error      error
}

type ServicesInNS struct {
	Namespace NSName
	Statuses  map[SvcName]*ServiceAndStatus
}

func makeSns(services string) map[SvcName]*ServiceAndStatus {
	watchService := make(map[SvcName]*ServiceAndStatus)
	for _, svc := range strings.Split(services, ",") {
		watchService[SvcName(svc)] = &ServiceAndStatus{
			Status: NotDeployed,
		}
	}
	return watchService
}

type PageData struct {
	PageTitle    string
	ServicesInNS map[NSName]ServicesInNS
}

func watched(service v1.Service, services string) bool {
	for _, svc := range strings.Split(services, ",") {
		if svc == service.Name {
			return true
		}
	}
	return false
}

func checkStatus(waitgroup *sync.WaitGroup, sns *ServiceAndStatus) {
	var httpClient = &http.Client{
		Timeout: time.Second * 2,
	}
	svc := sns.Service
	url := "http://" + svc.Name + "." + svc.Namespace + ".svc.cluster.local/healthz"
	fmt.Println("Checking", url, "...")
	resp, err := httpClient.Get(url)
	if err != nil {
		sns.Error = err
	} else {
		sns.StatusCode = resp.StatusCode
		if resp.StatusCode < 300 {
			sns.Status = Deployed
		}
	}
	waitgroup.Done()
}

func main() {
	port := ":8080"
	var err error
	var config *rest.Config

	services := os.Getenv("SERVICES")
	// Failsafe: use a near-empty array
	if services == "" {
		services = "a-frame,beluga"
	}

	// Kubernetes configuration setup
	if os.Getenv("IN_CLUSTER") == "true" {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	}
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	tmpl := template.Must(template.ParseFiles("templates/results.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var waitgroup sync.WaitGroup
		p := PageData{
			PageTitle:    "Hands-on Kubernetes Status Board",
			ServicesInNS: make(map[NSName]ServicesInNS),
		}
		fmt.Println("Handling request for /")
		serviceList, err := clientset.CoreV1().Services("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		for _, svc := range serviceList.Items {
			if !watched(svc, services) {
				continue
			}
			ns := NSName(svc.Namespace)
			if _, exists := p.ServicesInNS[ns]; !exists {
				p.ServicesInNS[ns] = ServicesInNS{
					Namespace: ns,
					Statuses:  makeSns(services),
				}
			}
			sns := &ServiceAndStatus{
				Service: svc,
				Status:  Failing,
			}
			p.ServicesInNS[ns].Statuses[SvcName(svc.Name)] = sns
			waitgroup.Add(1)
			go checkStatus(&waitgroup, sns)
		}
		waitgroup.Wait()
		err = tmpl.Execute(w, p)
		if err != nil {
			panic(err)
		}
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handling request for /healthz")
		fmt.Fprintf(w, "{\"status\": \"ok\"}")
	})

	fmt.Println("Server listening on", port)
	http.ListenAndServe(port, nil)
}
