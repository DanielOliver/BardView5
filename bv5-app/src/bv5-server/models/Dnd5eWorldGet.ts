/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { Created } from './Created';
import type { Dnd5eWorld } from './Dnd5eWorld';

export type Dnd5eWorldGet = (Dnd5eWorld & {
    dnd5eWorldId: string;
    created: Created;
    version: number;
});
