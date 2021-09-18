/** Types generated for queries found in "src/queries/user.sql" */
import { PreparedQuery } from '@pgtyped/query';

export type stringArray = (string)[];

/** 'SqlUsersFindByUid' parameters type */
export interface ISqlUsersFindByUidParams {
  session_id: string;
  uid: string;
}

/** 'SqlUsersFindByUid' return type */
export interface ISqlUsersFindByUidResult {
  id: string;
  uid: string;
  createdBy: string | null;
  createdAt: Date;
  effectiveDate: Date;
  endDate: Date | null;
  isActive: boolean;
  email: string;
  name: string;
  tags: stringArray;
}

/** 'SqlUsersFindByUid' query type */
export interface ISqlUsersFindByUidQuery {
  params: ISqlUsersFindByUidParams;
  result: ISqlUsersFindByUidResult;
}

const sqlUsersFindByUidIR: any = {"name":"sqlUsersFindByUid","params":[{"name":"session_id","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":224,"b":234,"line":6,"col":60},{"a":271,"b":281,"line":7,"col":20}]}},{"name":"uid","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":409,"b":412,"line":12,"col":15}]}}],"usedParamSet":{"session_id":true,"uid":true},"statement":{"body":"SELECT DISTINCT u.*\nFROM role_assignment ra\nINNER JOIN role r on ra.role_id = r.id\nINNER JOIN role_permission rp on r.id = rp.role_id\nINNER JOIN \"user\" u on evaluate_access_user(rp.conditions, :session_id!::bigint, u.id)\nWHERE ra.user_id = :session_id!::bigint\n  AND rp.subject = 'user'\n  AND rp.is_active = true\n  AND ra.is_active = true\n  AND r.is_active = true\n  AND u.uid = :uid!\nORDER BY u.id","loc":{"a":30,"b":426,"line":2,"col":0}}};

/**
 * Query generated from SQL:
 * ```
 * SELECT DISTINCT u.*
 * FROM role_assignment ra
 * INNER JOIN role r on ra.role_id = r.id
 * INNER JOIN role_permission rp on r.id = rp.role_id
 * INNER JOIN "user" u on evaluate_access_user(rp.conditions, :session_id!::bigint, u.id)
 * WHERE ra.user_id = :session_id!::bigint
 *   AND rp.subject = 'user'
 *   AND rp.is_active = true
 *   AND ra.is_active = true
 *   AND r.is_active = true
 *   AND u.uid = :uid!
 * ORDER BY u.id
 * ```
 */
export const sqlUsersFindByUid = new PreparedQuery<ISqlUsersFindByUidParams,ISqlUsersFindByUidResult>(sqlUsersFindByUidIR);


/** 'SqlUsersUpsert' parameters type */
export interface ISqlUsersUpsertParams {
  users: readonly ({
    uid: string,
    name: string,
    email: string,
    tags: stringArray,
    created_by: string | null | void
  })[];
}

/** 'SqlUsersUpsert' return type */
export interface ISqlUsersUpsertResult {
  id: string;
  uid: string;
}

/** 'SqlUsersUpsert' query type */
export interface ISqlUsersUpsertQuery {
  params: ISqlUsersUpsertParams;
  result: ISqlUsersUpsertResult;
}

const sqlUsersUpsertIR: any = {"name":"sqlUsersUpsert","params":[{"name":"users","codeRefs":{"defined":{"a":463,"b":467,"line":16,"col":9},"used":[{"a":590,"b":594,"line":19,"col":8}]},"transform":{"type":"pick_array_spread","keys":[{"name":"uid","required":true},{"name":"name","required":true},{"name":"email","required":true},{"name":"tags","required":true},{"name":"created_by","required":false}]},"required":false}],"usedParamSet":{"users":true},"statement":{"body":"INSERT INTO \"user\" as u (uid, name, email, tags, created_by)\nVALUES :users\nON CONFLICT (email) DO UPDATE set name = u.name\nRETURNING id, uid","loc":{"a":521,"b":660,"line":18,"col":0}}};

/**
 * Query generated from SQL:
 * ```
 * INSERT INTO "user" as u (uid, name, email, tags, created_by)
 * VALUES :users
 * ON CONFLICT (email) DO UPDATE set name = u.name
 * RETURNING id, uid
 * ```
 */
export const sqlUsersUpsert = new PreparedQuery<ISqlUsersUpsertParams,ISqlUsersUpsertResult>(sqlUsersUpsertIR);


/** 'SqlUsersGetPaginated' parameters type */
export interface ISqlUsersGetPaginatedParams {
  session_id: string;
  limit: string;
  offset: string;
}

/** 'SqlUsersGetPaginated' return type */
export interface ISqlUsersGetPaginatedResult {
  id: string;
  uid: string;
  createdBy: string | null;
  createdAt: Date;
  effectiveDate: Date;
  endDate: Date | null;
  isActive: boolean;
  email: string;
  name: string;
  tags: stringArray;
}

/** 'SqlUsersGetPaginated' query type */
export interface ISqlUsersGetPaginatedQuery {
  params: ISqlUsersGetPaginatedParams;
  result: ISqlUsersGetPaginatedResult;
}

const sqlUsersGetPaginatedIR: any = {"name":"sqlUsersGetPaginated","params":[{"name":"session_id","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":891,"b":901,"line":28,"col":60},{"a":938,"b":948,"line":29,"col":20}]}},{"name":"limit","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":1082,"b":1087,"line":35,"col":7}]}},{"name":"offset","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":1097,"b":1103,"line":35,"col":22}]}}],"usedParamSet":{"session_id":true,"limit":true,"offset":true},"statement":{"body":"SELECT DISTINCT u.*\nFROM role_assignment ra\nINNER JOIN role r on ra.role_id = r.id\nINNER JOIN role_permission rp on r.id = rp.role_id\nINNER JOIN \"user\" u on evaluate_access_user(rp.conditions, :session_id!::bigint, u.id)\nWHERE ra.user_id = :session_id!::bigint\n  AND rp.subject = 'user'\n  AND rp.is_active = true\n  AND ra.is_active = true\n  AND r.is_active = true\nORDER BY u.id\nLIMIT :limit! OFFSET :offset!","loc":{"a":697,"b":1103,"line":24,"col":0}}};

/**
 * Query generated from SQL:
 * ```
 * SELECT DISTINCT u.*
 * FROM role_assignment ra
 * INNER JOIN role r on ra.role_id = r.id
 * INNER JOIN role_permission rp on r.id = rp.role_id
 * INNER JOIN "user" u on evaluate_access_user(rp.conditions, :session_id!::bigint, u.id)
 * WHERE ra.user_id = :session_id!::bigint
 *   AND rp.subject = 'user'
 *   AND rp.is_active = true
 *   AND ra.is_active = true
 *   AND r.is_active = true
 * ORDER BY u.id
 * LIMIT :limit! OFFSET :offset!
 * ```
 */
export const sqlUsersGetPaginated = new PreparedQuery<ISqlUsersGetPaginatedParams,ISqlUsersGetPaginatedResult>(sqlUsersGetPaginatedIR);


