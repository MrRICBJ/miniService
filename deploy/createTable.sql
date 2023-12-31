CREATE TABLE items (
         "#" INT,
         id INT,
         uid VARCHAR(255),
         domain VARCHAR(255),
         cn VARCHAR(255),
         department VARCHAR(255),
         title VARCHAR(255),
         who VARCHAR(255),
         logon_count INT,
         num_logons7 INT,
         num_share7 INT,
         num_file7 INT,
         num_ad7 INT,
         num_n7 INT,
         num_logons14 INT,
         num_share14 INT,
         num_file14 INT,
         num_ad14 INT,
         num_n14 INT,
         num_logons30 INT,
         num_share30 INT,
         num_file30 INT,
         num_ad30 INT,
         num_n30 INT,
         num_logons150 INT,
         num_share150 INT,
         num_file150 INT,
         num_ad150 INT,
         num_n150 INT,
         num_logons365 INT,
         num_share365 INT,
         num_file365 INT,
         num_ad365 INT,
         num_n365 INT,
         has_user_principal_name INT,
         has_mail INT,
         has_phone INT,
         flag_disabled INT,
         flag_lockout INT,
         flag_password_not_required INT,
         flag_password_cant_change INT,
         flag_dont_expire_password INT,
         owned_files INT,
         num_mailboxes INT,
         num_member_of_groups INT,
         num_member_of_indirect_groups INT,
         member_of_indirect_groups_ids VARCHAR(255),
         member_of_groups_ids VARCHAR(500),
         is_admin INT,
         is_service INT
);

COPY items ("#",id,uid,domain,cn,department,title,who,logon_count,num_logons7,num_share7,num_file7,num_ad7,num_n7,num_logons14,num_share14,num_file14,num_ad14,num_n14,num_logons30,num_share30,num_file30,num_ad30,num_n30,num_logons150,num_share150,num_file150,num_ad150,num_n150,num_logons365,num_share365,num_file365,num_ad365,num_n365,has_user_principal_name,has_mail,has_phone,flag_disabled,flag_lockout,flag_password_not_required,flag_password_cant_change,flag_dont_expire_password,owned_files,num_mailboxes,num_member_of_groups,num_member_of_indirect_groups,member_of_indirect_groups_ids,member_of_groups_ids,is_admin,is_service)
    FROM '/ueba.csv'
    DELIMITER ','
    CSV HEADER;