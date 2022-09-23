package handlers

import (
	"net/http"
	"pricing/database"
	"pricing/models"
	"pricing/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	db *database.Database
}

// New returns new handler with given database connection
func New(db *database.Database) *handler {
	return &handler{db}
}

// CreateRule endpoint add array of given rules into
// database and returns a response for each rule
func (h *handler) CreateRule(c *gin.Context) {
	var rules []models.RawRule

	if err := c.BindJSON(&rules); err != nil {
		c.JSON(http.StatusOK, models.Response{
			OK:      false,
			Message: "invalid input",
		})
		return
	}

	responses := make([]models.Response, len(rules))

	for i := range rules {
		r := models.Response{OK: true}
		id, err := h.db.AddRule(rules[i])
		if err != nil {
			r.OK = false
			r.Message = err.Error()
		} else {
			r.Message = strconv.Itoa(id)
		}
		responses[i] = r
	}

	c.JSON(http.StatusOK, responses)
}

// ChangePrice endpoint returns the highest price with matching rules
func (h *handler) ChangePrice(c *gin.Context) {
	var tickets []models.RawTicket

	if err := c.BindJSON(&tickets); err != nil {
		c.JSON(http.StatusOK, models.Response{
			OK:      false,
			Message: "invalid input",
		})
		return
	}

	rules, err := h.db.FindRules(tickets)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			OK:      false,
			Message: err.Error(),
		})
		return
	}

	for i := range tickets {
		for j := range rules[i] {
			rule := rules[i][j]
			price, m := utils.AmountSum(
				tickets[i].BasePrice,
				rule[1][2:],
				rule[1][0],
			)
			if tickets[i].Markup < m {
				ruleID, _ := strconv.Atoi(rule[0])
				tickets[i].RuleID = ruleID
				tickets[i].PayablePrice = price
				tickets[i].Markup = m
			}
		}
	}

	c.JSON(http.StatusOK, tickets)
}
