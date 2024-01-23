package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitMeasurements := make(map[string]int)
	unitMeasurements["quarter_of_a_dozen"] = 3
	unitMeasurements["half_of_a_dozen"] = 6
	unitMeasurements["dozen"] = 12
	unitMeasurements["small_gross"] = 120
	unitMeasurements["gross"] = 144
	unitMeasurements["great_gross"] = 1728

	return unitMeasurements
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)

	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitName, exists := units[unit]
	if !exists {
		return false
	}

	_, present := bill[item]
	if !present {
		bill[item] = unitName
	} else {
		bill[item] += unitName
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	currentQuantity, billPresent := bill[item]
	if !billPresent {
		return false
	}
	unitQuantity, unitPresent := units[unit]
	if !unitPresent {
		return false
	}

	newQuantity := currentQuantity - unitQuantity
	if newQuantity < 0 {
		return false
	} else if newQuantity == 0 {
		delete(bill, item)
	} else {
		bill[item] = newQuantity
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, present := bill[item]
	return quantity, present
}
