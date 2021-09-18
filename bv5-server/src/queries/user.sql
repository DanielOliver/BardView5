/* @name sqlUsersFindByUid */
SELECT DISTINCT u.*
FROM role_assignment ra
INNER JOIN role r on ra.role_id = r.id
INNER JOIN role_permission rp on r.id = rp.role_id
INNER JOIN "user" u on evaluate_access_user(rp.conditions, :session_id!::bigint, u.id)
WHERE ra.user_id = :session_id!::bigint
  AND rp.subject = 'user'
  AND rp.is_active = true
  AND ra.is_active = true
  AND r.is_active = true
  AND u.uid = :uid!
ORDER BY u.id;

/* @name sqlUsersUpsert
  @param users -> ((uid!, name!, email!, tags!, created_by)...)
*/
INSERT INTO "user" as u (uid, name, email, tags, created_by)
VALUES :users
ON CONFLICT (email) DO UPDATE set name = u.name
RETURNING id, uid;

/* @name sqlUsersGetPaginated */
SELECT DISTINCT u.*
FROM role_assignment ra
INNER JOIN role r on ra.role_id = r.id
INNER JOIN role_permission rp on r.id = rp.role_id
INNER JOIN "user" u on evaluate_access_user(rp.conditions, :session_id!::bigint, u.id)
WHERE ra.user_id = :session_id!::bigint
  AND rp.subject = 'user'
  AND rp.is_active = true
  AND ra.is_active = true
  AND r.is_active = true
ORDER BY u.id
LIMIT :limit! OFFSET :offset!;
