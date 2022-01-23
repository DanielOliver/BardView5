/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { SystemTags } from './SystemTags';
import type { UserTags } from './UserTags';

export type Dnd5eWorld = {
    name: string;
    description: string;
    module?: string;
    active: boolean;
    commonAccess: string;
    userTags: UserTags;
    systemTags: SystemTags;
}