package main

import (
	"context"
	repositories "go_practice/repositories/users"
	"go_practice/tests"
	"go_practice/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/jaswdr/faker"
)

var Fake = faker.New()

func main(){
	tests.ResetApp()
	RunSeeders()
	srv := &http.Server{
		// Addr:    ":8080",
		Handler: tests.TestApp,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}


func RunSeeders(){
	CreateUsersSeeds()
}

func generateAnyUser() repositories.CreateUserDataType {
	person := Fake.Person()
	return repositories.CreateUserDataType{
		Name: person.Name(),
		Username: utils.ReverseStr(strings.ToLower(person.FirstName())),
		Email: person.Contact().Email,
		Password: Fake.Internet().Password(),
	}
}

func CreateUsersSeeds(){
	numberOfUsers := 100
    for range make([]struct{}, numberOfUsers) {
		_, result = repositories.CreateUser(generateAnyUser())
		if result.Error() != nil { 
			
		}
    }
}