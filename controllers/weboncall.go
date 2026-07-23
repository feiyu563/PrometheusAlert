package controllers

import (
	"PrometheusAlert/models"
	"encoding/json"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
)

// OnCall page
func (c *MainController) OnCall() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	ocs, err := models.GetAllOnCall()
	if err != nil {
		logs.Error("Error getting on-call schedules:", err)
	}

	// For display, we want to parse the JSON users
	type DisplayOnCall struct {
		Id    int64
		Date  string
		Users []models.OnCallUser
	}
	displayOcs := []DisplayOnCall{}
	for _, oc := range ocs {
		var users []models.OnCallUser
		json.Unmarshal([]byte(oc.Users), &users)
		displayOcs = append(displayOcs, DisplayOnCall{
			Id:    oc.Id,
			Date:  oc.Date,
			Users: users,
		})
	}

	c.Data["IsOnCall"] = true
	c.Data["OnCallSchedules"] = displayOcs
	c.TplName = "oncall.html"
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

// OnCallAdd page
func (c *MainController) OnCallAdd() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	today := time.Now().Format("2006-01-02")
	c.Data["IsOnCall"] = true
	c.Data["StartDate"] = today
	c.Data["EndDate"] = today
	c.TplName = "oncall_add.html"
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

// OnCallEdit page
func (c *MainController) OnCallEdit() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	id, _ := c.GetInt64("id")
	oc, err := models.GetOnCallById(id)
	if err != nil {
		logs.Error("Error getting on-call schedule by id:", err)
		c.Redirect("/oncall", 302)
		return
	}

	var users []models.OnCallUser
	json.Unmarshal([]byte(oc.Users), &users)

	// Convert date from "2006年1月2日" to "2006-01-02" for the date input
	parsedDate, err := time.Parse("2006年1月2日", oc.Date)
	var dateForInput string
	if err == nil {
		dateForInput = parsedDate.Format("2006-01-02")
	} else {
		logs.Error("Error parsing date for editing:", oc.Date, err)
		dateForInput = time.Now().Format("2006-01-02")
	}

	c.Data["IsOnCall"] = true
	c.Data["OnCall"] = oc
	c.Data["Users"] = users
	c.Data["StartDate"] = dateForInput
	c.Data["EndDate"] = dateForInput
	c.TplName = "oncall_add.html"
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
}

// SaveOnCall handles add/edit
func (c *MainController) SaveOnCall() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	startDateStr := c.GetString("start_date")
	endDateStr := c.GetString("end_date")
	names := c.GetStrings("name[]")
	phones := c.GetStrings("phone[]")

	var users []models.OnCallUser
	for i := 0; i < len(names); i++ {
		if strings.TrimSpace(names[i]) != "" && strings.TrimSpace(phones[i]) != "" {
			users = append(users, models.OnCallUser{
				Name:  names[i],
				Phone: phones[i],
			})
		}
	}

	usersJson, err := json.Marshal(users)
	if err != nil {
		logs.Error("Error marshalling users to JSON:", err)
		c.Redirect("/oncall", 302)
		return
	}
	usersJsonStr := string(usersJson)

	// Date range logic
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		logs.Error("Error parsing start date:", err)
		c.Redirect("/oncall", 302)
		return
	}
	endDate, err := time.Parse(layout, endDateStr)
	if err != nil {
		logs.Error("Error parsing end date:", err)
		c.Redirect("/oncall", 302)
		return
	}

	if endDate.Before(startDate) {
		logs.Error("End date is before start date")
		c.Redirect("/oncall/add", 302)
		return
	}

	// Loop through the date range and upsert records
	for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
		dateStringForDB := d.Format("2006年1月2日")

		existingOnCall, err := models.GetOnCallByDateString(dateStringForDB)
		if err == nil {
			// Exists, update it (overwrite)
			existingOnCall.Users = usersJsonStr
			models.UpdateOnCall(existingOnCall)
		} else {
			// Does not exist, add new
			models.AddOnCall(models.OnCall{Date: dateStringForDB, Users: usersJsonStr})
		}
	}

	c.Redirect("/oncall", 302)
}

// OnCallDel deletes a schedule
func (c *MainController) OnCallDel() {
	if !CheckAccount(c.Ctx) {
		c.Redirect("/login", 302)
		return
	}
	id, _ := c.GetInt64("id")
	err := models.DelOnCall(id)
	if err != nil {
		logs.Error("Error deleting on-call schedule:", err)
	}
	c.Redirect("/oncall", 302)
}
