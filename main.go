package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Give me the user's email to grant access to: ")
	user, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\nEnvs: qa, prod, dev\n")
	fmt.Print("Tell me the envs to grant access to (with space between): ")
	envsList, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("\nTell me the APIs to grant access to (with space between): ")
	apisList, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	stdOut, err := grantAccess(user, envsList, apisList)
	if err != nil {
		fmt.Printf("%v", string(stdOut[:]))
		log.Fatal(err)
	}
	fmt.Println("Done, grand master\n")
}

func grantAccess(user, envsList, apisList string) ([]byte, error) {
	envs := strings.Split(envsList, " ")
	APIs := strings.Split(apisList, " ")
	for _, env := range envs {
		for _, api := range APIs {
			fmt.Printf("Granting access to %s | env: %s | API: %s\n", user, envs, APIs)
			cmd := exec.Command("roadie2", "api-add-owner", env, api, user)
			stdoutErr, errAPI := cmd.CombinedOutput()
			if errAPI != nil {
				return stdoutErr, errAPI
			}
			fmt.Printf("Access granted to %s in API: %s | env: %s\n", user, api, env)
		}
	}
	return nil, nil
}
