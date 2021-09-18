/* @name sqlRolesInsertUserRole
*/
INSERT INTO "role" as u (uid, name, tags, created_by, role_type_id)
VALUES (:uid!::char(27), :name!::text, :tags!::text[], :created_by!::bigint, get_user_role_type_id())
ON CONFLICT (uid) DO UPDATE SET name = u.name
RETURNING id, uid;

/* @name sqlRolesLinkUserToRole
 */
INSERT INTO "role_assignment" as r (uid, created_by, role_id, user_id, tags)
SELECT :uid!::char(27), :created_by!::bigint, :role_id!::bigint, :user_id!::bigint, :tags!::text[]
ON CONFLICT (uid) DO UPDATE SET tags = r.tags
RETURNING id, uid;

/* @name sqlRolesLinkUserToGlobalUserRole
 */
INSERT INTO "role_assignment" as r (uid, created_by, role_id, user_id, tags)
SELECT :uid!::char(27), :created_by!::bigint, (SELECT Id FROM "role" WHERE name = 'User Role, Global'), :user_id!::bigint, :tags!::text[]
WHERE NOT EXISTS(
    SELECT 1 FROM role_assignment as r2 WHERE r2.user_id = :user_id! AND r2.role_id = (SELECT ro.Id FROM "role" as ro WHERE ro.name = 'User Role, Global')
    );




