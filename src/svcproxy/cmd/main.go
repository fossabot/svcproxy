package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"svcproxy/config"
	"svcproxy/service"
)

// Version to be filled by ldflags
var Version = "dev"

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "/etc/svcproxy/services.yaml"
	}

	cfg, err := config.Parse(configPath)
	if err != nil {
		log.Fatalf("Error parsing configuration: %s", err)
	}

	svc, err := service.NewService()
	if err != nil {
		log.Fatalf("Error creating service: %s", err)
	}

	for service := range cfg.Services {
		backend, err := url.Parse(service.Backend.URL)
		if err != nil {
			log.Fatalf("Error parsing url: %s", err)
		}

		svc.AddProxy(&service.Proxy{
			Frontend: &service.Frontend{
				FQDN: frontend,
			},
			Backend: &service.Backend{
				URL: backend,
			},
		})
	}

	http.ListenAndServe(":8080", svc)
}
