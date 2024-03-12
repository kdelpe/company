package database

const (
	GetAllClientsQuery = `
        SELECT 
            b.branch_id,
            b.branch_name,
            b.mgr_id,
            b.mgr_start_date,
            c.client_id,
            c.client_name
        FROM 
            branch b
        LEFT JOIN client c on b.branch_id = c.branch_id;
    `

	GetClientByIDQuery = `
        SELECT
            b.branch_id,
            b.branch_name,
            b.mgr_id,
            b.mgr_start_date,
            c.client_id,
            c.client_name
        FROM
            branch b
        LEFT JOIN 
            client c on b.branch_id = c.branch_id
        WHERE
            c.client_id = ?;
    `
)
