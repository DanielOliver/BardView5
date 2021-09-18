/** Types generated for queries found in "src/queries/role.sql" */
import { PreparedQuery } from '@pgtyped/query';

export type stringArray = (string)[];

/** 'SqlRolesInsertUserRole' parameters type */
export interface ISqlRolesInsertUserRoleParams {
  uid: string;
  name: string;
  tags: stringArray;
  created_by: string;
}

/** 'SqlRolesInsertUserRole' return type */
export interface ISqlRolesInsertUserRoleResult {
  id: string;
  uid: string;
}

/** 'SqlRolesInsertUserRole' query type */
export interface ISqlRolesInsertUserRoleQuery {
  params: ISqlRolesInsertUserRoleParams;
  result: ISqlRolesInsertUserRoleResult;
}

const sqlRolesInsertUserRoleIR: any = {"name":"sqlRolesInsertUserRole","params":[{"name":"uid","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":112,"b":115,"line":4,"col":9}]}},{"name":"name","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":129,"b":133,"line":4,"col":26}]}},{"name":"tags","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":143,"b":147,"line":4,"col":40}]}},{"name":"created_by","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":159,"b":169,"line":4,"col":56}]}}],"usedParamSet":{"uid":true,"name":true,"tags":true,"created_by":true},"statement":{"body":"INSERT INTO \"role\" as u (uid, name, tags, created_by, role_type_id)\nVALUES (:uid!::char(27), :name!::text, :tags!::text[], :created_by!::bigint, get_user_role_type_id())\nON CONFLICT (uid) DO UPDATE SET name = u.name\nRETURNING id, uid","loc":{"a":35,"b":267,"line":3,"col":0}}};

/**
 * Query generated from SQL:
 * ```
 * INSERT INTO "role" as u (uid, name, tags, created_by, role_type_id)
 * VALUES (:uid!::char(27), :name!::text, :tags!::text[], :created_by!::bigint, get_user_role_type_id())
 * ON CONFLICT (uid) DO UPDATE SET name = u.name
 * RETURNING id, uid
 * ```
 */
export const sqlRolesInsertUserRole = new PreparedQuery<ISqlRolesInsertUserRoleParams,ISqlRolesInsertUserRoleResult>(sqlRolesInsertUserRoleIR);


/** 'SqlRolesLinkUserToRole' parameters type */
export interface ISqlRolesLinkUserToRoleParams {
  uid: string;
  created_by: string;
  role_id: string;
  user_id: string;
  tags: stringArray;
}

/** 'SqlRolesLinkUserToRole' return type */
export interface ISqlRolesLinkUserToRoleResult {
  id: string;
  uid: string;
}

/** 'SqlRolesLinkUserToRole' query type */
export interface ISqlRolesLinkUserToRoleQuery {
  params: ISqlRolesLinkUserToRoleParams;
  result: ISqlRolesLinkUserToRoleResult;
}

const sqlRolesLinkUserToRoleIR: any = {"name":"sqlRolesLinkUserToRole","params":[{"name":"uid","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":392,"b":395,"line":11,"col":8}]}},{"name":"created_by","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":409,"b":419,"line":11,"col":25}]}},{"name":"role_id","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":431,"b":438,"line":11,"col":47}]}},{"name":"user_id","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":450,"b":457,"line":11,"col":66}]}},{"name":"tags","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":469,"b":473,"line":11,"col":85}]}}],"usedParamSet":{"uid":true,"created_by":true,"role_id":true,"user_id":true,"tags":true},"statement":{"body":"INSERT INTO \"role_assignment\" as r (uid, created_by, role_id, user_id, tags)\nSELECT :uid!::char(27), :created_by!::bigint, :role_id!::bigint, :user_id!::bigint, :tags!::text[]\nON CONFLICT (uid) DO UPDATE SET tags = r.tags\nRETURNING id, uid","loc":{"a":307,"b":545,"line":10,"col":0}}};

/**
 * Query generated from SQL:
 * ```
 * INSERT INTO "role_assignment" as r (uid, created_by, role_id, user_id, tags)
 * SELECT :uid!::char(27), :created_by!::bigint, :role_id!::bigint, :user_id!::bigint, :tags!::text[]
 * ON CONFLICT (uid) DO UPDATE SET tags = r.tags
 * RETURNING id, uid
 * ```
 */
export const sqlRolesLinkUserToRole = new PreparedQuery<ISqlRolesLinkUserToRoleParams,ISqlRolesLinkUserToRoleResult>(sqlRolesLinkUserToRoleIR);


/** 'SqlRolesLinkUserToGlobalUserRole' parameters type */
export interface ISqlRolesLinkUserToGlobalUserRoleParams {
  uid: string;
  created_by: string;
  user_id: string;
  tags: stringArray;
}

/** 'SqlRolesLinkUserToGlobalUserRole' return type */
export type ISqlRolesLinkUserToGlobalUserRoleResult = void;

/** 'SqlRolesLinkUserToGlobalUserRole' query type */
export interface ISqlRolesLinkUserToGlobalUserRoleQuery {
  params: ISqlRolesLinkUserToGlobalUserRoleParams;
  result: ISqlRolesLinkUserToGlobalUserRoleResult;
}

const sqlRolesLinkUserToGlobalUserRoleIR: any = {"name":"sqlRolesLinkUserToGlobalUserRole","params":[{"name":"uid","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":680,"b":683,"line":18,"col":8}]}},{"name":"created_by","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":697,"b":707,"line":18,"col":25}]}},{"name":"user_id","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":777,"b":784,"line":18,"col":105},{"a":888,"b":895,"line":20,"col":60}]}},{"name":"tags","required":true,"transform":{"type":"scalar"},"codeRefs":{"used":[{"a":796,"b":800,"line":18,"col":124}]}}],"usedParamSet":{"uid":true,"created_by":true,"user_id":true,"tags":true},"statement":{"body":"INSERT INTO \"role_assignment\" as r (uid, created_by, role_id, user_id, tags)\nSELECT :uid!::char(27), :created_by!::bigint, (SELECT Id FROM \"role\" WHERE name = 'User Role, Global'), :user_id!::bigint, :tags!::text[]\nWHERE NOT EXISTS(\n    SELECT 1 FROM role_assignment as r2 WHERE r2.user_id = :user_id! AND r2.role_id = (SELECT ro.Id FROM \"role\" as ro WHERE ro.name = 'User Role, Global')\n    )","loc":{"a":595,"b":987,"line":17,"col":0}}};

/**
 * Query generated from SQL:
 * ```
 * INSERT INTO "role_assignment" as r (uid, created_by, role_id, user_id, tags)
 * SELECT :uid!::char(27), :created_by!::bigint, (SELECT Id FROM "role" WHERE name = 'User Role, Global'), :user_id!::bigint, :tags!::text[]
 * WHERE NOT EXISTS(
 *     SELECT 1 FROM role_assignment as r2 WHERE r2.user_id = :user_id! AND r2.role_id = (SELECT ro.Id FROM "role" as ro WHERE ro.name = 'User Role, Global')
 *     )
 * ```
 */
export const sqlRolesLinkUserToGlobalUserRole = new PreparedQuery<ISqlRolesLinkUserToGlobalUserRoleParams,ISqlRolesLinkUserToGlobalUserRoleResult>(sqlRolesLinkUserToGlobalUserRoleIR);


