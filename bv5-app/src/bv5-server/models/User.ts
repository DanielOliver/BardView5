/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { Email } from './Email';
import type { SystemTags } from './SystemTags';
import type { UserTags } from './UserTags';

export type User = {
    name: string;
    active: boolean;
    commonAccess: string;
    email: Email;
    userTags: UserTags;
    systemTags: SystemTags;
}