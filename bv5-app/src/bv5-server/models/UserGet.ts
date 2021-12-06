/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { Created } from './Created';
import type { User } from './User';

export type UserGet = (User & {
userId: number;
created: Created;
version: number;
});