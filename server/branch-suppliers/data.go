package branch_suppliers

import "example/company/server/branch"

type BranchSuppliers struct {
	Branch       branch.Branch `json:"branch"`
	SupplierName *string       `json:"supplier_name"`
	SupplyType   *string       `json:"supply_type"`
}
