package main

import (
	"log"
	"strings"
	"time"

	"github.com/pa3ng/dt-breakfast/internal/gmail"
	"github.com/pa3ng/dt-breakfast/internal/sheets"
)

func main() {
	// Get dates
	now := time.Now()
	day := now.Format("Mon")
	if day != "Sat" {
		log.Println("This program can only be executed on Saturdays")
		return
	}

	// Get volunteer
	log.Println("Program executed")
	volunteer, err := sheets.GetVolunteer(now.Format("1/2/2006"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("This week's volunteer:", volunteer)

	// Email volunteer reminder
	email := getGpMail(volunteer)
	log.Println("Emailling", volunteer, "at", email)
	gmail.SendReminderEmail(email)
}

func getGpMail(fullname string) string {
	fullname = strings.ToLower(fullname)
	fullnamesplit := strings.Split(fullname, " ")
	firstname := fullnamesplit[0]
	lastnamesplit := fullnamesplit[1:] // ex. ['Santo', 'Domingo']
	email := firstname + "." + strings.Join(lastnamesplit, "") + "@gpmail.org"
	return email
}
