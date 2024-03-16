package client

const (
	GetAllClientsQuery = `
        SELECT 
            b.branch_id,
            b.branch_name,
            b.mgr_id,
            DATE_FORMAT(b.mgr_start_date, '%Y-%m-%d'),
            c.client_id,
            c.client_name
        FROM 
            branch b
        LEFT JOIN 
                client c on b.branch_id = c.branch_id
        ORDER BY 
            client_id;
    `

	GetClientByIDQuery = `
        SELECT
            b.branch_id,
            b.branch_name,
            b.mgr_id,
            DATE_FORMAT(b.mgr_start_date, '%Y-%m-%d'),
            c.client_id,
            c.client_name
        FROM
            branch b
        LEFT JOIN 
            client c on b.branch_id = c.branch_id
        WHERE
            c.client_id = ?;
    `

	PostClientQuery = `
	        INSERT INTO
				client (client_name, branch_id)
	        VALUES (?, ?);
`

	PUTClientQuery = `
	UPDATE 
	    client 
	SET 
	    client_name = ?, client.branch_id = ?
	WHERE 
	    client_id = ?;
`

	DELETEClientQuery = `
	DELETE FROM 
	           client 
	WHERE 
	    client_id = ?;
`
)
