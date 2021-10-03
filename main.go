package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
	// clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api/v1"
)

func main() {

	kubecofig := getConfig()

	file, err := ioutil.ReadFile(kubecofig)

	if err != nil {
		fmt.Println("couldnot read kubecofig file", err)
		return
	}

	fmt.Println(file)

	config := &clientcmdapi.Config{}
	err = yaml.Unmarshal(file, &config)

	if err != nil {
		fmt.Println("couldnot parse kubecofig file", err)
		return
	}

	fmt.Println(config)
}

func getConfig() string {
	k := os.Getenv("KUBECONFIG")

	if k != "" {
		return k
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	return home + "/.kube/config"
}
