/** Types generated for queries found in "src/queries/user.sql" */
import { PreparedQuery } from '@pgtyped/query';

export type stringArray = (string)[];

/** 'SqlUsersFindByUid' parameters type */
export interface ISqlUsersFindByUidParams {
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

const sqlUsersFindByUidIR: any = {"name":"sqlUsersFindByUid","params":[{"name":"uid","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":64,"b":67,"line":2,"col":34}]}}],"usedParamSet":{"uid":true},"statement":{"body":"SELECT * FROM \"user\" WHERE uid = :uid!","loc":{"a":30,"b":67,"line":2,"col":0}}};

/**
 * Query generated from SQL:
 * ```
 * SELECT * FROM "user" WHERE uid = :uid!
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

const sqlUsersUpsertIR: any = {"name":"sqlUsersUpsert","params":[{"name":"users","codeRefs":{"defined":{"a":104,"b":108,"line":5,"col":9},"used":[{"a":231,"b":235,"line":8,"col":8}]},"transform":{"type":"pick_array_spread","keys":[{"name":"uid","required":true},{"name":"name","required":true},{"name":"email","required":true},{"name":"tags","required":true},{"name":"created_by","required":false}]},"required":false}],"usedParamSet":{"users":true},"statement":{"body":"INSERT INTO \"user\" as u (uid, name, email, tags, created_by)\nVALUES :users\nON CONFLICT (email) DO UPDATE set name = u.name\nRETURNING id, uid","loc":{"a":162,"b":301,"line":7,"col":0}}};

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

const sqlUsersGetPaginatedIR: any = {"name":"sqlUsersGetPaginated","params":[{"name":"limit","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":378,"b":383,"line":13,"col":40}]}},{"name":"offset","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":393,"b":399,"line":13,"col":55}]}}],"usedParamSet":{"limit":true,"offset":true},"statement":{"body":"SELECT * FROM \"user\" ORDER BY id LIMIT :limit! OFFSET :offset!","loc":{"a":338,"b":399,"line":13,"col":0}}};

/**
 * Query generated from SQL:
 * ```
 * SELECT * FROM "user" ORDER BY id LIMIT :limit! OFFSET :offset!
 * ```
 */
export const sqlUsersGetPaginated = new PreparedQuery<ISqlUsersGetPaginatedParams,ISqlUsersGetPaginatedResult>(sqlUsersGetPaginatedIR);


