/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { Created } from './Created';
import type { User } from './User';

export type UserGet = (User & {
userId: string;
created: Created;
version: number;
uuid: string;
});