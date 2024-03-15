package branch_suppliers

const (
	GetAllBranchSuppliersQuery = `
        SELECT
            b.branch_id,
            b.branch_name,
            b.mgr_id,
            b.mgr_start_date,
            bs.supplier_name,
            bs.supply_type
        FROM
            branch b
        LEFT JOIN
            branch_supplier bs ON b.branch_id = bs.branch_id;
    `

	GetBranchSupplierByIDQuery = `
        SELECT
            b.branch_id,
            b.branch_name,
            b.mgr_id,
            b.mgr_start_date,
            bs.supplier_name,
            bs.supply_type
        FROM
            branch b
        LEFT JOIN
            branch_supplier bs ON b.branch_id = bs.branch_id
        WHERE
            bs.branch_id = ?;
    `
)
