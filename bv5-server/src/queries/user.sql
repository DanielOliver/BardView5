/* @name sqlUsersFindByUid */
SELECT * FROM "user" WHERE uid = :uid!;

/* @name sqlUsersUpsert
  @param users -> ((uid!, name!, email!, tags!, created_by)...)
*/
INSERT INTO "user" as u (uid, name, email, tags, created_by)
VALUES :users
ON CONFLICT (email) DO UPDATE set name = u.name
RETURNING id, uid;

/* @name sqlUsersGetPaginated */
SELECT * FROM "user" ORDER BY id LIMIT :limit! OFFSET :offset!;
