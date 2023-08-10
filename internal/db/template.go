package db

const (
	createItems = `
		CREATE TABLE IF NOT EXISTS items (
			ID INT PRIMARY KEY,
			UID VARCHAR(255),
			Domain VARCHAR(255),
			CN VARCHAR(255),
			Department VARCHAR(255),
			Title VARCHAR(255),
			Who VARCHAR(255),
			LogonCount INT,
			NumLogons7 INT,
			NumShare7 INT,
			NumFile7 INT,
			NumAD7 INT,
			NumN7 INT,
			NumLogons14 INT,
			NumShare14 INT,
			NumFile14 INT,
			NumAD14 INT,
			NumN14 INT,
			NumLogons30 INT,
			NumShare30 INT,
			NumFile30 INT,
			NumAD30 INT,
			NumN30 INT,
			NumLogons150 INT,
			NumShare150 INT,
			NumFile150 INT,
			NumAD150 INT,
			NumN150 INT,
			NumLogons365 INT,
			NumShare365 INT,
			NumFile365 INT,
			NumAD365 INT,
			NumN365 INT,
			HasUserPrincipalName BOOLEAN,
			HasMail BOOLEAN,
			HasPhone BOOLEAN,
			FlagDisabled BOOLEAN,
			FlagLockout BOOLEAN,
			FlagPasswordNotRequired BOOLEAN,
			FlagPasswordCantChange BOOLEAN,
			FlagDontExpirePassword BOOLEAN,
			OwnedFiles INT,
			NumMailboxes INT,
			NumMemberOfGroups INT,
			NumMemberOfIndirectGroups INT,
			MemberOfIndirectGroupsIDs TEXT,
			MemberOfGroupsIDs TEXT,
			IsAdmin BOOLEAN,
			IsService BOOLEAN
		);
	`

	dropItems = `
		DROP TABLE IF EXISTS items;
	`

	SelectById = `SELECT * FROM items WHERE id = $1`
)
