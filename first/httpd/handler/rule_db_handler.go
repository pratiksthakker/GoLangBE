package handler

import (
	"first/db"
	"first/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetRules() []models.Rules {
	var rule models.Rules
	rules := []models.Rules{}
	database := db.OpenDbConnection()
	database.AutoMigrate(&rule)
	database.Raw("Select * from rules").Scan(&rules)
	//database.Find(&products)
	defer database.Close()
	//fmt.Println(rules)
	return rules
}

func RunRulesOnCart(c *gin.Context) {

	rules := GetRules()
	fmt.Println("got rules", rules)
	var cart models.CartOrder
	c.BindJSON(&cart)

	fmt.Println("got cart:", cart)
	orderid := UpdateOrder(cart)

	cartItems := cart.CartItems
	skipRule := false
	var discountPerc float32
	var discountOnPrice float32
	var toatalPrice float32
	var cart_item_ids []int
	var appliedRule models.Rules
	isCouponPresent := false

	if len(cart.CouponCode) > 0 {
		isCouponPresent = true
	} else {
		isCouponPresent = false
	}

	for _, cartitem := range cartItems {
		cart_item_ids = append(cart_item_ids, cartitem.Id)
		toatalPrice += cartitem.Price * float32(cartitem.Quantity)
	}

	if isCouponPresent {
		for _, rule := range rules {
			if strings.Compare(rule.CouponCode, cart.CouponCode) == 0 {
				couponRules := GetRulesConfigPerRuleId(rule.ID)
				for _, coupon_rule_item := range couponRules {
					appliedRule = rule
					discountPerc = coupon_rule_item.DiscountValue
					discountOnPrice = toatalPrice
				}
			}
		}
	} else {

		for _, rule := range rules {
			skipRule = false
			discountOnPrice = 0.0
			ruledetails := GetRulesConfigPerRuleId(rule.ID)
			var rule_item_ids []int
			for _, ruledetail := range ruledetails {
				rule_item_ids = append(rule_item_ids, ruledetail.ItemID)
			}
			fmt.Println("rule items id", rule_item_ids)
			fmt.Println("cart items id", cart_item_ids)
			if !isToConsiderRule(rule_item_ids, cart_item_ids) {
				continue
			}
			for _, rule_item := range ruledetails {
				for _, item := range cartItems {
					fmt.Println("first", rule_item.ItemID)
					fmt.Println("first", item.Id)
					fmt.Println("first", rule_item.Qty)
					fmt.Println("first", item.Quantity)
					if rule_item.ItemID == item.Id {
						if rule_item.Qty <= item.Quantity {
							discountPerc = rule_item.DiscountValue
							if strings.Compare(rule_item.DiscountType, "ATLEAST") == 0 {
								discountOnPrice += item.Price * float32(item.Quantity)
							} else if strings.Compare(rule_item.DiscountType, "FIXED") == 0 {
								discountOnPrice += item.Price * float32(rule_item.Qty)
							} else if strings.Compare(rule_item.DiscountType, "ALL") == 0 {
								discountOnPrice += toatalPrice
							}
							fmt.Println("Running rule", rule.ID)
							fmt.Println("discount perc", discountPerc)
							fmt.Println("discount price", discountOnPrice)
						} else {
							fmt.Println("skipping", rule_item.ItemID)
							fmt.Println("skipping", item.Id)
							skipRule = true
							break
						}
					}
				}
				if skipRule {
					discountPerc = 0.0
					discountOnPrice = 0.0
					fmt.Println("Skipping rule", rule.ID)
					break
				}
			}
			fmt.Println("discount perc", discountPerc)
			fmt.Println("discount price", discountOnPrice)
			if discountOnPrice > 0.0 {
				appliedRule = rule
				break
			}
		}

	}

	fmt.Println("discount perc", discountPerc)
	fmt.Println("discount price", discountOnPrice)

	if discountOnPrice > 0.0 {
		toatalDiscount := discountOnPrice * (discountPerc / 100)
		fmt.Println("Total price", toatalPrice)
		fmt.Println("Total Discount", toatalDiscount)
		cart.DiscountPercent = discountPerc
		cart.TotalPrice = toatalPrice
		cart.DiscountedPrice = toatalPrice - toatalDiscount
		cart.DiscountDescription = appliedRule.Description
	}

	if isCouponPresent && discountOnPrice <= 0.0 {
		cart.DiscountDescription = "Invalid Coupon Code Passed Please try again later"
	}
	cart.OrderId = orderid
	c.JSON(http.StatusOK, gin.H{
		"newCart": cart,
	})
	// database := db.OpenDbConnection()
	// database.AutoMigrate(&rule)
	// database.Raw("Select * from rules").Scan(&rules)
	// //database.Find(&products)
	// defer database.Close()
	// fmt.Println(rules)
	// return rules
}

func isToConsiderRule(rule_item_ids, cart_item_ids []int) bool {

	set := make(map[int]int)
	for _, value := range cart_item_ids {
		set[value]++
	}

	for _, value := range rule_item_ids {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

func GetRulesConfigPerRuleId(rule_id int) []models.RulesConfig {

	var rule models.RulesConfig
	rules := []models.RulesConfig{}
	database := db.OpenDbConnection()
	database.AutoMigrate(&rule)
	database.Raw("Select * from rule_config").Scan(&rules)
	//database.Find(&products)
	defer database.Close()
	var filteredRules []models.RulesConfig
	fmt.Println("current rule ID", rule_id)
	for _, temprule := range rules {
		fmt.Println("current ruleID", temprule.RulesId)
		if temprule.RulesId == rule_id {
			fmt.Println("Adding", temprule)
			filteredRules = append(filteredRules, temprule)
			fmt.Println("filtered Rule inside loop", filteredRules)
		}
	}
	fmt.Println("filtered Rule", filteredRules)
	return filteredRules
}

// func GetAllRulesConfig() []models.RulesConfig {

// 	ruledetail := models.RulesConfig{}
// 	ruledetails := []models.RulesConfig{}
// 	database := db.OpenDbConnection()
// 	database.AutoMigrate(&ruledetail)
// 	database.Raw("Select * from rule_config").Scan(&ruledetails)
// 	//database.Find(&products)
// 	defer database.Close()
// 	fmt.Println(ruledetails)
// 	return ruledetails
// }
