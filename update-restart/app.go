package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const Version = "0.1.0"

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (b *App) startup(ctx context.Context) {
	// Perform your setup here
	b.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (b *App) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (b *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (b *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func SeltUpdate() bool {
	v := semver.MustParse(Version)
	latest, err := selfupdate.UpdateSelf(v, "achhabra2/riftshare")
	if err != nil {
		log.Println("Binary update failed:", err)
		return false
	}
	if latest.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		log.Println("Current binary is the latest version", Version)
		return true
	} else {
		log.Println("Successfully updated to version", latest.Version)
		log.Println("Release note:\n", latest.ReleaseNotes)
		return true
	}
}

func SelfUpdateMac() bool {
	latest, found, _ := selfupdate.DetectLatest("achhabra2/riftshare")
	if found {
		homeDir, _ := os.UserHomeDir()
		downloadPath := filepath.Join(homeDir, "Downloads", "RiftShare.zip")
		err := exec.Command("curl", "-L", latest.AssetURL, "-o", downloadPath).Run()
		if err != nil {
			log.Println("curl error:", err)
			return false
		}
		var appPath string
		cmdPath, err := os.Executable()
		appPath = strings.TrimSuffix(cmdPath, "RiftShare.app/Contents/MacOS/RiftShare")
		if err != nil {
			appPath = "/Applications/"
		}
		err = exec.Command("ditto", "-xk", downloadPath, appPath).Run()
		if err != nil {
			log.Println("ditto error:", err)
			return false
		}
		err = exec.Command("rm", downloadPath).Run()
		if err != nil {
			log.Println("removing error:", err)
			return false
		}
		return true
	} else {
		return false
	}
}

type Updater struct {
	l net.Listener
}

func (u *Updater) CheckForUpdate() string {
	/*
			latest, found, err := selfupdate.DetectLatest("achhabra2/riftshare")
			if err != nil {
				log.Println("Error occurred while detecting version:", err)
				return false, ""
			}

			v := semver.MustParse(Version)
			if !found || latest.Version.LTE(v) {
				log.Println("Current version is the latest")
				return false, ""
			}
		return true, latest.Version.String()
	*/

	// do a restart
	// goagain.ForkExec(u.l)
	aa := "/Users/wwestgarth/work/wails-update-restart/update-restart/build/bin/update-restart.app"
	cmd := exec.Command("open", "-a", aa)
	if err := cmd.Run(); err != nil {
		log.Fatal("ahhh")
	}

	// kill ourselves
	// log.Fatal("we killed ourselves")

	// pretend we always have a latest
	return "0.2.0"
}
