/* @name FindUserByUid */
SELECT * FROM "user" WHERE uid = :uid!;

/* @name InsertNewUser
  @param users -> ((uid!, name!, email!, tags!)...)
*/
INSERT INTO "user" as u (uid, name, email, tags)
VALUES :users
ON CONFLICT (email) DO UPDATE set name = u.name
RETURNING id, uid;
