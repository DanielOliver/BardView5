/** Types generated for queries found in "src/queries/user.sql" */
import { PreparedQuery } from '@pgtyped/query';

export type stringArray = (string)[];

/** 'FindUserByUid' parameters type */
export interface IFindUserByUidParams {
  uid: string;
}

/** 'FindUserByUid' return type */
export interface IFindUserByUidResult {
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

/** 'FindUserByUid' query type */
export interface IFindUserByUidQuery {
  params: IFindUserByUidParams;
  result: IFindUserByUidResult;
}

const findUserByUidIR: any = {"name":"FindUserByUid","params":[{"name":"uid","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":61,"b":64,"line":2,"col":34}]}}],"usedParamSet":{"uid":true},"statement":{"body":"SELECT * FROM \"user\" WHERE uid = :uid!","loc":{"a":27,"b":64,"line":2,"col":0}}};

/**
 * Query generated from SQL:
 * ```
 * SELECT * FROM "user" WHERE uid = :uid!
 * ```
 */
export const findUserByUid = new PreparedQuery<IFindUserByUidParams,IFindUserByUidResult>(findUserByUidIR);


/** 'InsertNewUser' parameters type */
export interface IInsertNewUserParams {
  users: readonly ({
    uid: string,
    name: string,
    email: string,
    tags: stringArray
  })[];
}

/** 'InsertNewUser' return type */
export interface IInsertNewUserResult {
  id: string;
  uid: string;
}

/** 'InsertNewUser' query type */
export interface IInsertNewUserQuery {
  params: IInsertNewUserParams;
  result: IInsertNewUserResult;
}

const insertNewUserIR: any = {"name":"InsertNewUser","params":[{"name":"users","codeRefs":{"defined":{"a":103,"b":107,"line":5,"col":9},"used":[{"a":209,"b":213,"line":8,"col":8}]},"transform":{"type":"pick_array_spread","keys":[{"name":"uid","required":true},{"name":"name","required":true},{"name":"email","required":true},{"name":"tags","required":true}]},"required":false}],"usedParamSet":{"users":true},"statement":{"body":"INSERT INTO \"user\" as u (uid, name, email, tags)\r\nVALUES :users\r\nON CONFLICT (email) DO UPDATE set name = u.name\r\nRETURNING id, uid","loc":{"a":151,"b":281,"line":7,"col":0}}};

/**
 * Query generated from SQL:
 * ```
 * INSERT INTO "user" as u (uid, name, email, tags)
 * VALUES :users
 * ON CONFLICT (email) DO UPDATE set name = u.name
 * RETURNING id, uid
 * ```
 */
export const insertNewUser = new PreparedQuery<IInsertNewUserParams,IInsertNewUserResult>(insertNewUserIR);


