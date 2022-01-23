//go:build mage
// +build mage

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var bardviewexe = "bardview5.exe"
var localhostPostgresql = "postgresql://postgres:mysecretpassword@localhost/bardview5?sslmode=disable"
var dockerComposeLocal = "docker-compose-local.yml"
var dockerCompose = sh.RunCmd("docker-compose", "-f", dockerComposeLocal)

func init() {
	if exe := os.Getenv("BARDVIEW5EXE"); exe != "" {
		bardviewexe = exe
	}
}

func Clean() error {
	fmt.Println("Run: Clean")
	return os.RemoveAll("build")
}

func ModDownload() error {
	fmt.Println("Run: ModDownload")
	return sh.Run("go", "mod", "download")
}

func Build() error {
	fmt.Println("Run: Build")

	newpath := filepath.Join(".", "build")
	if err := os.MkdirAll(newpath, os.ModePerm); err != nil {
		return err
	}

	return sh.Run("go", "build", "-o", fmt.Sprintf("build/%s", bardviewexe))
}

// https://alex.dzyoba.com/blog/go-connect-loop/
// ConnectLoop tries to connect to the DB under given DSN using a give driver
// in a loop until connection succeeds. timeout specifies the timeout for the
// loop.
func ConnectLoop(timeout time.Duration) error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	timeoutExceeded := time.After(timeout)
	for {
		select {
		case <-timeoutExceeded:
			return fmt.Errorf("db connection failed after %s timeout", timeout)

		case <-ticker.C:
			_, err := sql.Open("postgres", localhostPostgresql)
			if err == nil {
				return nil
			}
			fmt.Println("%v", err)
			fmt.Println("failed to connect to db")
		}
	}
}

func DumpSchema() error {
	fmt.Println("Run: DumpSchema")
	if err := dockerCompose("exec", "-T", "db", "/bin/bash", "-c", "pg_dump -U postgres -s bardview5 > /sql_dump/snapshot.sql"); err != nil {
		return err
	}
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	return sh.RunV("docker", "run", "--rm", "-v", fmt.Sprintf("%s:/src", strings.ReplaceAll(pwd, "\\", "/")), "-w", "/src", "kjconroy/sqlc", "generate")
}

func DockerComposeDown() error {
	return dockerCompose("down")
}

func DockerComposeUp() error {
	return dockerCompose("up", "-d")
}

func Migrate() error {
	fmt.Println("Run: Migrate")
	defer DockerComposeDown()

	DockerComposeUp()

	return InternalMigrate()
}

func InternalMigrate() error {
	if err := ConnectLoop(time.Minute); err != nil {
		return err
	}

	if err := sh.Run("go", "generate", "./..."); err != nil {
		return err
	}

	if err := sh.RunWith(map[string]string{
		"BARDVIEW5_CONNECTION": localhostPostgresql,
	}, "go", "run", ".", "migrate"); err != nil {
		return err
	}

	mg.Deps(DumpSchema)

	if err := sh.Run("go", "generate", "./..."); err != nil {
		return err
	}

	return nil
}

func Test() error {
	if err := sh.Run("go", "test", "./..."); err != nil {
		return err
	}

	return nil
}

func Up() error {
	if err := DockerComposeUp(); err != nil {
		return err
	}
	mg.Deps(InternalMigrate)
	return nil
}

func Down() error {
	if err := DockerComposeDown(); err != nil {
		return err
	}
	return nil
}

func All() error {
	defer DockerComposeDown()
	DockerComposeUp()
	mg.Deps(Clean)
	mg.Deps(InternalMigrate)
	mg.Deps(Test)
	mg.Deps(Build)
	return nil
}
